package api

import (
	"log"
	"net/http"
	"todo/models"

	"todo/app/database"
	"todo/constants"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	todoRepository database.TodoRepository
}

func NewTodoHandler(todoRepository database.TodoRepository) *TodoHandler {
	return &TodoHandler{
		todoRepository: todoRepository,
	}
}

func (h *TodoHandler) CreateTodo(ctx *fiber.Ctx) error {
	request := new(models.CreateTodoRequest) //

	if err := ctx.BodyParser(&request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, constants.BadRequestMessage, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	} //แปลงค่า
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, nil, constants.BadRequestMessage).SendResponse(ctx, http.StatusBadRequest)
	} //ค่า validat ห้ามเป็นค่าว่าง ทดสอบความถูกต้อง

	if err := h.todoRepository.CreateTodo(ctx.Context(), request); err != nil {
		return models.Response(constants.StatusCodeSystemError, nil, constants.StatusCodeSystemErrorMessage).SendResponse(ctx, http.StatusInternalServerError)
	} //เป็นฟังชันก์ที่เราเขียนเอง   //การเรียกใช้
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, nil).SendResponseSuccess(ctx, http.StatusOK)
}

func (h *TodoHandler) ReadTodo(ctx *fiber.Ctx) error {
	request := new(models.ReadTodoRequest)
	if err := ctx.BodyParser(&request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, constants.BadRequestMessage, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, nil, constants.BadRequestMessage).SendResponse(ctx, http.StatusBadRequest)
	}
	data, err := h.todoRepository.ReadTodo(ctx.Context(), request)
	if err != nil {
		log.Println(err.Error())
		return models.Response(constants.StatusCodeBadRequest, data, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, data).SendResponseSuccess(ctx, http.StatusOK)
}

func (h *TodoHandler) UpdateTodo(ctx *fiber.Ctx) error {
	request := new(models.UpdateTodoRequest)

	if err := ctx.BodyParser(&request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, constants.BadRequestMessage, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	} //แปลงค่า
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, nil, constants.BadRequestMessage).SendResponse(ctx, http.StatusBadRequest)
	} //ค่า validat ห้ามเป็นค่าว่าง ทดสอบความถูกต้อง

	if err := h.todoRepository.UpdateTodo(ctx.Context(), request); err != nil {
		return models.Response(constants.StatusCodeSystemError, nil, constants.StatusCodeSystemErrorMessage).SendResponse(ctx, http.StatusInternalServerError)
	} //เป็นฟังชันก์ที่เราเขียนเอง   //การเรียกใช้
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, nil).SendResponseSuccess(ctx, http.StatusOK)
}

func (h *TodoHandler) DeleteTodo(ctx *fiber.Ctx) error {
	request := new(models.DeleteTodo)

	if err := ctx.BodyParser(&request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, constants.BadRequestMessage, err.Error()).SendResponse(ctx, http.StatusBadRequest)
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return models.Response(constants.StatusCodeBadRequest, nil, constants.BadRequestMessage).SendResponse(ctx, http.StatusBadRequest)
	}

	if err := h.todoRepository.DeleteTodo(ctx.Context(), request); err != nil {
		return models.Response(constants.StatusCodeSystemError, nil, constants.StatusCodeSystemErrorMessage).SendResponse(ctx, http.StatusInternalServerError)
	}
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, nil).SendResponseSuccess(ctx, http.StatusOK)
}

func (h *TodoHandler) ReadTodoAll(ctx *fiber.Ctx) error {
	data, err := h.todoRepository.ReadTodoAll(ctx.Context())
	if err != nil {
		return models.Response(constants.StatusCodeSystemError, nil, constants.StatusCodeSystemErrorMessage).SendResponse(ctx, http.StatusInternalServerError)
	}
	return models.ResponseSuccess(constants.StatusCodeSuccess, constants.SuccessMessage, data).SendResponseSuccess(ctx, http.StatusOK)
}
