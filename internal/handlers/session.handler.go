package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type SessionHandler struct {
	sessionService *services.SessionService
	logger         *zap.Logger
}

func NewSessionHandler(sessionService *services.SessionService, logger *zap.Logger) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
		logger:         logger,
	}
}

// AddSession handles adding a new session
func (h *SessionHandler) AddSession(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		h.logger.Warn("Yetkisiz erişim girişimi", zap.String("sebep", "userID boş"))
		response.Error(ctx, http.StatusUnauthorized, "Giriş yapmadınız")
		return
	}

	var req models.AddSessionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("İstek verisi bağlanamadı", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Geçersiz istek")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			response.Error(ctx, http.StatusBadRequest, "Doğrulama hatası", response.WithValidationErrors(validationErrors))
			return
		}
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.UserId = userID

	session, err := h.sessionService.AddSession(req)
	if err != nil {
		h.logger.Error("Oturum eklenemedi", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Oturum eklenemedi, lütfen tekrar deneyin")
		return
	}

	response.Success(ctx, "Oturum başarıyla eklendi.", response.WithPayload(session))
}

// // UpdateSession handles updating an existing session
// func (h *SessionHandler) UpdateSession(ctx *gin.Context) {
// 	userID := ctx.GetString("userId")
// 	if userID == "" {
// 		h.logger.Warn("Yetkisiz erişim girişimi", zap.String("sebep", "userID boş"))
// 		response.Error(ctx, http.StatusUnauthorized, "Giriş yapmadınız")
// 		return
// 	}

// 	var req models.UpdateSessionRequest
// 	if err := ctx.ShouldBindJSON(&req); err != nil {
// 		h.logger.Error("İstek verisi bağlanamadı", zap.Error(err))
// 		response.Error(ctx, http.StatusBadRequest, "Geçersiz istek")
// 		return
// 	}

// 	validate := validator.New()
// 	if err := validate.Struct(&req); err != nil {
// 		var validationErrors validator.ValidationErrors
// 		if errors.As(err, &validationErrors) {
// 			response.Error(ctx, http.StatusBadRequest, "Doğrulama hatası", response.WithValidationErrors(validationErrors))
// 			return
// 		}
// 		response.Error(ctx, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	req.UserId = userID

// 	session, err := h.sessionService.UpdateSession(req)
// 	if err != nil {
// 		h.logger.Error("Oturum güncellenemedi", zap.Error(err))
// 		response.Error(ctx, http.StatusInternalServerError, "Oturum güncellenemedi, lütfen tekrar deneyin")
// 		return
// 	}

// 	response.Success(ctx, "Oturum başarıyla güncellendi.", response.WithPayload(session))
// }

// DeleteSession handles deleting a session
func (h *SessionHandler) DeleteSession(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		h.logger.Warn("Yetkisiz erişim girişimi", zap.String("sebep", "userID boş"))
		response.Error(ctx, http.StatusUnauthorized, "Giriş yapmadınız")
		return
	}

	id := ctx.Param("id")
	if id == "" {
		response.Error(ctx, http.StatusBadRequest, "Oturum ID'si gereklidir")
		return
	}

	err := h.sessionService.DeleteSession(id, userID)
	if err != nil {
		h.logger.Error("Oturum silinemedi", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Oturum silinemedi, lütfen tekrar deneyin")
		return
	}

	response.Success(ctx, "Oturum başarıyla silindi.")
}

// GetSessionsOfDay handles retrieving all sessions for a specific day
func (h *SessionHandler) GetSessionsOfDay(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		h.logger.Warn("Yetkisiz erişim girişimi", zap.String("sebep", "userID boş"))
		response.Error(ctx, http.StatusUnauthorized, "Giriş yapmadınız")
		return
	}

	dayStr := ctx.Param("day")
	if dayStr == "" {
		response.Error(ctx, http.StatusBadRequest, "Gün bilgisi gereklidir")
		return
	}

	day, err := time.Parse(time.RFC3339, dayStr)
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "Geçersiz tarih formatı")
		return
	}

	sessions, err := h.sessionService.GetSessionsOfDay(userID, day)
	if err != nil {
		h.logger.Error("Oturumlar alınamadı", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Oturumlar alınamadı, lütfen tekrar deneyin")
		return
	}

	response.Success(ctx, "Oturumlar başarıyla alındı.", response.WithPayload(sessions))
}

// ToggleCompletion handles toggling the completion status of a session
func (h *SessionHandler) ToggleCompletion(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		h.logger.Warn("Yetkisiz erişim girişimi", zap.String("sebep", "userID boş"))
		response.Error(ctx, http.StatusUnauthorized, "Giriş yapmadınız")
		return
	}

	id := ctx.Param("id")
	if id == "" {
		response.Error(ctx, http.StatusBadRequest, "Oturum ID'si gereklidir")
		return
	}

	var req struct {
		IsCompleted bool `json:"isCompleted"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("İstek verisi bağlanamadı", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "Geçersiz istek")
		return
	}

	err := h.sessionService.ToggleCompletion(id, userID, req.IsCompleted)
	if err != nil {
		h.logger.Error("Tamamlanma durumu değiştirilemedi", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "Tamamlanma durumu değiştirilemedi, lütfen tekrar deneyin")
		return
	}

	response.Success(ctx, "Tamamlanma durumu başarıyla değiştirildi.")
}
