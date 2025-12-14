package handler

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"user-api/internal/logger"
	"user-api/internal/models"
	"user-api/internal/service"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service:  s,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}
	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.service.CreateUser(c.Context(), req.Name, req.Dob)
	if err != nil {
		logger.Log.Error("failed to create user",
			zap.Error(err),
			zap.String("name", req.Name),
		)

		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	logger.Log.Info("user created",
		zap.Int32("id", user.ID),
		zap.String("name", user.Name),
	)
	return c.Status(201).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	user, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		logger.Log.Error("failed to get user",
			zap.Error(err),
			zap.Int("user_id", id),
		)
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	logger.Log.Info("user fetched",
		zap.Int("user_id", id),
	)
	return c.JSON(user)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid limit"})
	}

	offset, err := strconv.Atoi(c.Query("offset", "0"))
	if err != nil || offset < 0 {
		return c.Status(400).JSON(fiber.Map{"error": "invalid offset"})
	}

	users, err := h.service.ListUsers(
		c.Context(),
		int32(limit),
		int32(offset),
	)
	if err != nil {
		logger.Log.Error("failed to list users",
			zap.Error(err),
		)
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to fetch users",
		})
	}

	logger.Log.Info("users listed",
		zap.Int("count", len(users)),
	)
	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.UpdateUser(c.Context(), int32(id), req.Name, req.Dob)
	logger.Log.Error("failed to update user",
		zap.Error(err),
		zap.Int("user_id", id),
	)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "update failed"})
	}

	logger.Log.Info("user updated",
		zap.Int("user_id", id),
	)
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid user id"})
	}

	if err := h.service.DeleteUser(c.Context(), int32(id)); err != nil {
		logger.Log.Error("failed to delete user",
			zap.Error(err),
			zap.Int("user_id", id),
		)
		return c.Status(500).JSON(fiber.Map{"error": "delete failed"})
	}

	logger.Log.Info("user deleted",
		zap.Int("user_id", id),
	)
	return c.SendStatus(204)
}
