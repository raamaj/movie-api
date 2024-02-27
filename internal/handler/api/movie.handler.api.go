package api

import (
	"github.com/gofiber/fiber/v2"
	"movie-api/internal/domain"
	"movie-api/internal/service"
	"strconv"
)

type movieHandler struct {
	Service service.MovieService
}

func NewMovieHandler(service service.MovieService) movieHandler {
	return movieHandler{
		Service: service,
	}
}

// movieListHandler function
// @Summary Get movie results
// @Description Retrieve the results for each movie
// @Tags Movie
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /Movies [get]
func (m *movieHandler) movieListHandler(c *fiber.Ctx) error {
	results, err := m.Service.List()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Get List Movie Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}

// movieViewHandler function
// @Summary Get movie data
// @Description Retrieve the movie data
// @Tags Movie
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /Movies/:id [get]
func (m *movieHandler) movieViewHandler(c *fiber.Ctx) error {

	idStr := c.Params("id")

	id, _ := strconv.Atoi(idStr)

	results, err := m.Service.View(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Get Data Movie Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}

// movieInsertHandler function
// @Summary createMovie to the application
// @Description Create Movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param body body domain.MovieRequest true "Movie information"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /Movies [post]
func (m *movieHandler) movieInsertHandler(c *fiber.Ctx) error {

	var body domain.MovieRequest

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.ValidateInsert(); err != nil {
		var errorString []string
		for _, e := range err {
			errorString = append(errorString, e.Error())
		}
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    errorString,
			StatusCode: fiber.StatusBadRequest,
		})
	}

	results, err := m.Service.Insert(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Insert Movie Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}

// movieUpdateHandler function
// @Summary updateMovie to the application
// @Description Update Movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Param body body domain.MovieRequest true "Movie information"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /Movies/:id [patch]
func (m *movieHandler) movieUpdateHandler(c *fiber.Ctx) error {

	var body domain.MovieRequest

	idStr := c.Params("id")

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	id, _ := strconv.Atoi(idStr)

	body.ID = uint(id)

	if err := body.ValidateUpdate(); err != nil {
		var errorString []string
		for _, e := range err {
			errorString = append(errorString, e.Error())
		}
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    errorString,
			StatusCode: fiber.StatusBadRequest,
		})
	}

	results, err := m.Service.Update(&body)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Update Movie Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}

// movieDeleteHandler function
// @Summary deleteMovie to the application
// @Description Delete Movie
// @Tags Movie
// @Accept json
// @Produce json
// @Param id path string true "Movie ID"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /Movies/:id [delete]
func (m *movieHandler) movieDeleteHandler(c *fiber.Ctx) error {

	idStr := c.Params("id")

	id, _ := strconv.Atoi(idStr)

	results, err := m.Service.Delete(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			StatusCode: fiber.StatusBadRequest,
			Message:    err.Error(),
			Error:      true,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "Delete Movie Success",
		Data:       results,
		StatusCode: fiber.StatusOK,
	})
}
