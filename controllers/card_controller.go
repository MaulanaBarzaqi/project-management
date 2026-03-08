package controllers

import (
	"time"

	"github.com/MaulanaBarzaqi/project-management/models"
	"github.com/MaulanaBarzaqi/project-management/services"
	"github.com/MaulanaBarzaqi/project-management/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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

func (c *CardController) UpdateCard(ctx *fiber.Ctx) error {
	publicID := ctx.Params("id")
	type UpdateCardRequest struct {
		ListPublicID	string		`json:"list_id"`
		Title			string 		`json:"title"`
		Description		string		`json:"description"`
		DueDate			*time.Time	`json:"due_date"`
		Position		int			`json:"position"`
	}
	var req UpdateCardRequest
	if err := ctx.BodyParser(&req); err != nil {
		return utils.BadRequest(ctx, "failed to parsing data", err.Error())
	}
	if _, err := uuid.Parse(publicID); err != nil {
		return utils.BadRequest(ctx, "id not valid", err.Error())
	}
	card := &models.Card{
		Title: req.Title,
		Description: req.Description,
		DueDate: req.DueDate,
		Position: req.Position,
		PublicID: uuid.MustParse(publicID),
	}
	if err := c.services.Update(card, req.ListPublicID); err != nil {
		return utils.InternalServerError(ctx, "failed to pubdate data", err.Error())
	}
	// // masalah created at
	// updatedCard, err := c.services.GetByPublicID(publicID)
	// if err != nil {
	// 	return utils.InternalServerError(ctx, "failed to fetch updated data", err.Error())
	// }
	return utils.Success(ctx, "success to update card", card)
}