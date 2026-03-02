package controllers

import (
	"time"

	"github.com/MaulanaBarzaqi/project-management/models"
	"github.com/MaulanaBarzaqi/project-management/services"
	"github.com/MaulanaBarzaqi/project-management/utils"
	"github.com/gofiber/fiber/v2"
)

type CardController struct {
	services services.CardService
}

func NewCardController(s services.CardService) *CardController {
	return &CardController{services: s}
}

func (c *CardController) CreateCard(ctx *fiber.Ctx) error {
	type CreateCardRequest struct {
		ListPublicID 	string 		`json:"list_id"`
		Title			string		`json:"title"`
		Description		string		`json:"description"`
		DueDate			time.Time	`json:"due_date"`
		Position		int			`json:"position"`
	}
	var req CreateCardRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "failed to get data", err.Error())
	}
	card := &models.Card{
		Title: req.Title,
		Description: req.Description,
		DueDate: &req.DueDate,
		Position: req.Position,
	}
	if err := c.services.Create(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "failed to create card", err.Error())
	}
	return utils.Success(ctx, "success to created card", card)
}