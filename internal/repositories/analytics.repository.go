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
		{{Key: "$setOnInsert", Value: bson.M{
			"_id":            bson.NewObjectID(),
			"user_id":        analytics.UserId,
			"exam_type":      analytics.ExamType,
			"type":           "exam",
			"exam_count":     0,
			"max_result":     analytics.Result,
			"average_result": 0,
			"result_series":  bson.A{},
		}}},

		{{Key: "$set", Value: bson.M{
			"exam_count": bson.M{"$add": bson.A{"$exam_count", 1}},
		}}},

		{{Key: "$set", Value: bson.M{
			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{analytics.Result, "$max_result"}},
					"then": analytics.Result,
					"else": "$max_result",
				},
			},
		}}},

		{{Key: "$set", Value: bson.M{
			"average_result": bson.M{
				"$divide": bson.A{
					bson.M{
						"$add": bson.A{
							bson.M{"$multiply": bson.A{"$average_result", bson.M{"$subtract": bson.A{"$exam_count", 1}}}},
							analytics.Result,
						},
					},
					"$exam_count",
				},
			},
		}}},

		{{Key: "$push", Value: bson.M{
			"result_series": bson.M{
				"date":   analytics.Date,
				"name":   analytics.Name,
				"result": analytics.Result,
			},
		}}},
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
		{{Key: "$setOnInsert", Value: bson.M{
			"_id":            bson.NewObjectID(),
			"user_id":        analytics.UserId,
			"exam_type":      analytics.ExamType,
			"type":           "lesson",
			"lesson":         analytics.Lesson,
			"exam_count":     0,
			"max_result":     analytics.Result,
			"average_result": 0,
			"average_time":   0,
			"result_series":  bson.A{},
		}}},

		{{Key: "$set", Value: bson.M{
			"exam_count": bson.M{"$add": bson.A{"$exam_count", 1}},
		}}},

		{{Key: "$set", Value: bson.M{
			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{analytics.Result, "$max_result"}},
					"then": analytics.Result,
					"else": "$max_result",
				},
			},
		}}},

		{{Key: "$set", Value: bson.M{
			"average_result": bson.M{
				"$divide": bson.A{
					bson.M{
						"$add": bson.A{
							bson.M{"$multiply": bson.A{"$average_result", bson.M{"$subtract": bson.A{"$exam_count", 1}}}},
							analytics.Result,
						},
					},
					"$exam_count",
				},
			},
		}}},

		{{Key: "$set", Value: bson.M{
			"average_time": bson.M{
				"$divide": bson.A{
					bson.M{
						"$add": bson.A{
							bson.M{"$multiply": bson.A{"$average_time", bson.M{"$subtract": bson.A{"$exam_count", 1}}}},
							analytics.Time,
						},
					},
					"$exam_count",
				},
			},
		}}},

		{{Key: "$push", Value: bson.M{
			"result_series": bson.M{
				"date":   analytics.Date,
				"name":   analytics.Name,
				"result": analytics.Result,
				"time":   analytics.Time,
			},
		}}},
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

		// 1) Verilen exam'ı result_series'ten çıkar
		{{Key: "$set", Value: bson.M{
			"result_series": bson.M{
				"$filter": bson.M{
					"input": "$result_series",
					"as":    "item",
					"cond":  bson.M{"$ne": bson.A{"$$item.name", analytics.Name}},
				},
			},
		}}},

		// 2) Exam count'ı yeniden hesapla
		{{Key: "$set", Value: bson.M{
			"exam_count": bson.M{"$size": "$result_series"},
		}}},

		// 3) max_result yeniden hesapla
		{{Key: "$set", Value: bson.M{
			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{bson.M{"$size": "$result_series"}, 0}},
					"then": bson.M{"$max": "$result_series.result"},
					"else": 0,
				},
			},
		}}},

		// 4) average_result yeniden hesapla
		{{Key: "$set", Value: bson.M{
			"average_result": bson.M{
				"$cond": bson.M{
					"if": bson.M{"$gt": bson.A{bson.M{"$size": "$result_series"}, 0}},
					"then": bson.M{
						"$divide": bson.A{
							bson.M{"$sum": "$result_series.result"},
							bson.M{"$size": "$result_series"},
						},
					},
					"else": 0,
				},
			},
		}}},
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

		// 1) Verilen exam'ı result_series'ten çıkar
		{{Key: "$set", Value: bson.M{
			"result_series": bson.M{
				"$filter": bson.M{
					"input": "$result_series",
					"as":    "item",
					"cond":  bson.M{"$ne": bson.A{"$$item.name", analytics.Name}},
				},
			},
		}}},

		// 2) Exam count'ı yeniden hesapla
		{{Key: "$set", Value: bson.M{
			"exam_count": bson.M{"$size": "$result_series"},
		}}},

		// 3) max_result yeniden hesapla
		{{Key: "$set", Value: bson.M{
			"max_result": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$gt": bson.A{bson.M{"$size": "$result_series"}, 0}},
					"then": bson.M{"$max": "$result_series.result"},
					"else": 0,
				},
			},
		}}},

		// 4) average_result yeniden hesapla
		{{Key: "$set", Value: bson.M{
			"average_result": bson.M{
				"$cond": bson.M{
					"if": bson.M{"$gt": bson.A{bson.M{"$size": "$result_series"}, 0}},
					"then": bson.M{
						"$divide": bson.A{
							bson.M{"$sum": "$result_series.result"},
							bson.M{"$size": "$result_series"},
						},
					},
					"else": 0,
				},
			},
		}}},

		{{Key: "$set", Value: bson.M{
			"average_time": bson.M{
				"$cond": bson.M{
					"if": bson.M{"$gt": bson.A{bson.M{"$size": "$result_series"}, 0}},
					"then": bson.M{
						"$divide": bson.A{
							bson.M{"$sum": "$result_series.time"},
							bson.M{"$size": "$result_series"},
						},
					},
					"else": 0,
				},
			},
		}}},
	}

	_, err := r.collection.UpdateOne(context.TODO(), filter, update)
	return err
}
