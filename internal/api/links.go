package api

import (
	"errors"

	"github.com/gofiber/fiber/v3"
	"github.com/lila-network/go-urlsh/internal/fault"
	"github.com/lila-network/go-urlsh/internal/links"
)

type CreateLinkRequest struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Description string `json:"description"`
}

type UpdateLinkRequest struct {
	URL         *string `json:"url"`
	Description *string `json:"description"`
}

type LinkHandler struct {
	repository links.Repository
}

func NewLinkHandler(repository links.Repository) *LinkHandler {
	return &LinkHandler{repository: repository}
}

func (l *LinkHandler) GetAll(c fiber.Ctx) error {
	ctx := c.Context()

	dbLinks, err := l.repository.GetAll(ctx)
	if err != nil {
		if errors.Is(err, fault.ErrNotFound) {
			return WriteSuccess(c, fiber.StatusOK, []links.Link{})
		}

		return WriteError(c, fiber.StatusInternalServerError, "Internal Server Error")

	}

	return WriteSuccess(c, fiber.StatusOK, dbLinks)
}

func (l *LinkHandler) Get(c fiber.Ctx) error {
	ctx := c.Context()

	id := c.Params("id")
	if id == "" {
		return WriteError(c, fiber.StatusBadRequest, "Request is invalid")
	}

	link, err := l.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, fault.ErrNotFound) {
			return WriteError(c, fiber.StatusNotFound, "Link not found")
		}
		return WriteError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	return WriteSuccess(c, fiber.StatusOK, link)
}

func (l *LinkHandler) Post(c fiber.Ctx) error {
	ctx := c.Context()
	var linkRequest CreateLinkRequest

	err := c.Bind().Body(&linkRequest)
	if err != nil {
		return WriteError(c, fiber.StatusBadRequest, "Request is invalid")
	}

	link := links.Link{}
	link.ID = linkRequest.ID
	link.URL = linkRequest.URL
	link.Description = linkRequest.Description

	dbLink, err := l.repository.Create(ctx, link)
	if err != nil {
		if errors.Is(err, fault.ErrAlreadyExists) {
			return WriteError(c, fiber.StatusConflict, "A duplicate link already exists")
		}

		return WriteError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	return WriteSuccess(c, fiber.StatusCreated, dbLink)
}

func (l *LinkHandler) Patch(c fiber.Ctx) error {
	ctx := c.Context()

	id := c.Params("id")
	if id == "" {
		return WriteError(c, fiber.StatusBadRequest, "Request is invalid")
	}

	var requestBody UpdateLinkRequest

	err := c.Bind().Body(&requestBody)
	if err != nil {
		return WriteError(c, fiber.StatusBadRequest, "Request is invalid")
	}

	dbLink, err := l.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, fault.ErrNotFound) {
			return WriteError(c, fiber.StatusNotFound, "Link not found")
		}

		return WriteError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	if requestBody.Description != nil {
		dbLink.Description = *requestBody.Description
	}

	if requestBody.URL != nil {
		dbLink.URL = *requestBody.URL
	}

	if !dbLink.IsValid() {
		return WriteError(c, fiber.StatusBadRequest, "Request is invalid")
	}

	updatedDBLink, err := l.repository.Update(ctx, dbLink)
	if err != nil {
		if errors.Is(err, fault.ErrNotFound) {
			return WriteError(c, fiber.StatusNotFound, "Link not found")
		}

		return WriteError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	return WriteSuccess(c, fiber.StatusOK, updatedDBLink)
}

func (l *LinkHandler) Delete(c fiber.Ctx) error {
	ctx := c.Context()
	id := c.Params("id")

	if id == "" {
		return WriteError(c, fiber.StatusBadRequest, "Request is invalid")
	}

	err := l.repository.DeleteByID(ctx, id)
	if err != nil {
		if errors.Is(err, fault.ErrNotFound) {
			return WriteError(c, fiber.StatusNotFound, "Link not found")
		}

		return WriteError(c, fiber.StatusInternalServerError, "Internal Server Error")
	}

	return WriteNoContent(c)
}
