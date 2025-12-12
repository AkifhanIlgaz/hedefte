package repositories

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type AnalyticsRepository interface {
	UpsertExamAnalytics(analytics models.UpsertExamAnalytics) error
	UpsertLessonAnalytics(analytics models.UpsertLessonAnalytics) error
}

type analyticsRepository struct {
	collection *mongo.Collection
}

func NewAnalyticsRepository(db *mongo.Database) AnalyticsRepository {
	collection := db.Collection(constants.AnalyticsCollection)

	return analyticsRepository{
		collection: collection,
	}
}

func (r analyticsRepository) UpsertExamAnalytics(analytics models.UpsertExamAnalytics) error {
	filter := bson.M{
		"user_id":   analytics.UserId,
		"type":      "exam",
		"exam_type": analytics.ExamType,
	}

	update := mongo.Pipeline{
		// 1️⃣ Default alanları garanti altına al
		{{Key: "$set", Value: bson.M{
			"user_id":   analytics.UserId,
			"exam_type": analytics.ExamType,
			"type":      "exam",

			"exam_count":     bson.M{"$ifNull": bson.A{"$exam_count", 0}},
			"max_result":     bson.M{"$ifNull": bson.A{"$max_result", analytics.Result}},
			"average_result": bson.M{"$ifNull": bson.A{"$average_result", 0}},
			"result_series":  bson.M{"$ifNull": bson.A{"$result_series", bson.A{}}},
		}}},

		// 2️⃣ new_exam_count hesapla
		{{Key: "$set", Value: bson.M{
			"new_exam_count": bson.M{"$add": bson.A{"$exam_count", 1}},
		}}},

		// 3️⃣ Metrics güncelle
		{{Key: "$set", Value: bson.M{
			"exam_count": "$new_exam_count",

			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{analytics.Result, "$max_result"}},
					"then": analytics.Result,
					"else": "$max_result",
				},
			},

			"average_result": bson.M{
				"$divide": bson.A{
					bson.M{
						"$add": bson.A{
							bson.M{
								"$multiply": bson.A{
									"$average_result",
									bson.M{"$subtract": bson.A{"$new_exam_count", 1}},
								},
							},
							analytics.Result,
						},
					},
					"$new_exam_count",
				},
			},
		}}},

		// 4️⃣ result_series append (pipeline uyumlu)
		{{Key: "$set", Value: bson.M{
			"result_series": bson.M{
				"$concatArrays": bson.A{
					"$result_series",
					bson.A{
						bson.M{
							"date":   analytics.Date,
							"name":   analytics.Name,
							"result": analytics.Result,
						},
					},
				},
			},
		}}},

		// 5️⃣ temp field temizle
		{{Key: "$unset", Value: "new_exam_count"}},
	}
	opts := options.UpdateOne().SetUpsert(true)

	_, err := r.collection.UpdateOne(context.TODO(), filter, update, opts)
	return err

}

func (r analyticsRepository) UpsertLessonAnalytics(analytics models.UpsertLessonAnalytics) error {
	filter := bson.M{
		"user_id":   analytics.UserId,
		"type":      "lesson",
		"exam_type": analytics.ExamType,
		"lesson":    analytics.Lesson,
	}

	update := mongo.Pipeline{
		{{Key: "$set", Value: bson.M{
			"user_id":   analytics.UserId,
			"exam_type": analytics.ExamType,
			"type":      "lesson",
			"lesson":    analytics.Lesson,

			"exam_count":     bson.M{"$ifNull": bson.A{"$exam_count", 0}},
			"max_result":     bson.M{"$ifNull": bson.A{"$max_result", analytics.Result}},
			"average_result": bson.M{"$ifNull": bson.A{"$average_result", 0}},
			"average_time":   bson.M{"$ifNull": bson.A{"$average_time", 0}},
			"result_series":  bson.M{"$ifNull": bson.A{"$result_series", bson.A{}}},
		}}},

		{{Key: "$set", Value: bson.M{
			"new_exam_count": bson.M{"$add": bson.A{"$exam_count", 1}},
		}}},
		{{Key: "$set", Value: bson.M{
			"exam_count": "$new_exam_count",

			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{analytics.Result, "$max_result"}},
					"then": analytics.Result,
					"else": "$max_result",
				},
			},

			"average_result": bson.M{
				"$divide": bson.A{
					bson.M{
						"$add": bson.A{
							bson.M{
								"$multiply": bson.A{
									"$average_result",
									bson.M{"$subtract": bson.A{"$new_exam_count", 1}},
								},
							},
							analytics.Result,
						},
					},
					"$new_exam_count",
				},
			},

			"average_time": bson.M{
				"$divide": bson.A{
					bson.M{
						"$add": bson.A{
							bson.M{
								"$multiply": bson.A{
									"$average_time",
									bson.M{"$subtract": bson.A{"$new_exam_count", 1}},
								},
							},
							analytics.Time,
						},
					},
					"$new_exam_count",
				},
			},
		}}},

		{{Key: "$set", Value: bson.M{
			"result_series": bson.M{
				"$concatArrays": bson.A{
					"$result_series",
					bson.A{
						bson.M{
							"date":   analytics.Date,
							"name":   analytics.Name,
							"result": analytics.Result,
							"time":   analytics.Time,
						},
					},
				},
			},
		}}},

		// 5️⃣ temp field temizle
		{{Key: "$unset", Value: "new_exam_count"}},
	}
	opts := options.UpdateOne().SetUpsert(true)

	_, err := r.collection.UpdateOne(context.TODO(), filter, update, opts)
	return err
}

func (r analyticsRepository) DeleteExamAnalytics(analytics models.DeleteExamAnalytics) error {
	filter := bson.M{
		"user_id":   analytics.UserId,
		"exam_type": analytics.ExamType,
		"type":      "exam",
		"name":      analytics.Name,
	}

	update := mongo.Pipeline{

		// 1️⃣ İlgili exam'ı result_series'ten çıkar
		{{Key: "$set", Value: bson.M{
			"result_series": bson.M{
				"$filter": bson.M{
					"input": "$result_series",
					"as":    "item",
					"cond":  bson.M{"$ne": bson.A{"$$item.name", analytics.Name}},
				},
			},
		}}},

		// 2️⃣ result değerlerini flat array'e çıkar
		{{Key: "$set", Value: bson.M{
			"result_values": bson.M{
				"$map": bson.M{
					"input": "$result_series",
					"as":    "item",
					"in":    "$$item.result",
				},
			},
		}}},

		// 3️⃣ exam_count, max_result, average_result hesapla
		{{Key: "$set", Value: bson.M{
			"exam_count": bson.M{"$size": "$result_values"},

			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{bson.M{"$size": "$result_values"}, 0}},
					"then": bson.M{"$max": "$result_values"},
					"else": 0,
				},
			},

			"average_result": bson.M{
				"$cond": bson.M{
					"if": bson.M{"$gt": bson.A{bson.M{"$size": "$result_values"}, 0}},
					"then": bson.M{
						"$divide": bson.A{
							bson.M{"$sum": "$result_values"},
							bson.M{"$size": "$result_values"},
						},
					},
					"else": 0,
				},
			},
		}}},

		// 4️⃣ temp field temizle
		{{Key: "$unset", Value: "result_values"}},
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func (r analyticsRepository) DeleteLessonAnalytics(analytics models.DeleteLessonAnalytics) error {
	filter := bson.M{
		"user_id":   analytics.UserId,
		"exam_type": analytics.ExamType,
		"type":      "lesson",
		"name":      analytics.Name,
		"lesson":    analytics.Lesson,
	}

	update := mongo.Pipeline{

		// 1️⃣ Exam'i result_series'ten çıkar
		{{Key: "$set", Value: bson.M{
			"result_series": bson.M{
				"$filter": bson.M{
					"input": "$result_series",
					"as":    "item",
					"cond":  bson.M{"$ne": bson.A{"$$item.name", analytics.Name}},
				},
			},
		}}},

		// 2️⃣ result ve time değerlerini ayrı array’lere çıkar
		{{Key: "$set", Value: bson.M{
			"result_values": bson.M{
				"$map": bson.M{
					"input": "$result_series",
					"as":    "item",
					"in":    "$$item.result",
				},
			},
			"time_values": bson.M{
				"$map": bson.M{
					"input": "$result_series",
					"as":    "item",
					"in":    "$$item.time",
				},
			},
		}}},

		// 3️⃣ exam_count, max, average hesapla
		{{Key: "$set", Value: bson.M{
			"exam_count": bson.M{"$size": "$result_values"},

			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{bson.M{"$size": "$result_values"}, 0}},
					"then": bson.M{"$max": "$result_values"},
					"else": 0,
				},
			},

			"average_result": bson.M{
				"$cond": bson.M{
					"if": bson.M{"$gt": bson.A{bson.M{"$size": "$result_values"}, 0}},
					"then": bson.M{
						"$divide": bson.A{
							bson.M{"$sum": "$result_values"},
							bson.M{"$size": "$result_values"},
						},
					},
					"else": 0,
				},
			},

			"average_time": bson.M{
				"$cond": bson.M{
					"if": bson.M{"$gt": bson.A{bson.M{"$size": "$time_values"}, 0}},
					"then": bson.M{
						"$divide": bson.A{
							bson.M{"$sum": "$time_values"},
							bson.M{"$size": "$time_values"},
						},
					},
					"else": 0,
				},
			},
		}}},

		// 4️⃣ Temp field'leri temizle
		{{Key: "$unset", Value: bson.A{
			"result_values",
			"time_values",
		}}},
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
