package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oganes5796/employment-test/internal/domain"
)

// @Summary Создание задачи
// @Description Создаёт новую задачу
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body models.Task true "Детали задачи"
// @Success 201 {object} models.Task
// @Failure 400 {object} fiber.Map "Некорректный запрос"
// @Failure 500 {object} fiber.Map "Ошибка сервера"
// @Router /tasks [post]

func (i *Implementation) CreateTodo(c *fiber.Ctx) error {
	var input domain.Todo
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	id, err := i.todoService.Create(c.Context(), &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": id})
}

// @Summary Получение списка задач
// @Description Возвращает список всех задач
// @Tags tasks
// @Produce json
// @Success 200 {array} models.Task
// @Failure 500 {object} fiber.Map "Ошибка сервера"
// @Router /tasks [get]

func (i *Implementation) GetTodo(c *fiber.Ctx) error {
	todo, err := i.todoService.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(todo)
}

// @Summary Обновление задачи
// @Description Обновляет данные о задаче
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID задачи"
// @Param task body models.Task true "Обновлённые данные задачи"
// @Success 200
// @Failure 400 {object} fiber.Map "Некорректный ID"
// @Failure 500 {object} fiber.Map "Ошибка сервера"
// @Router /tasks/{id} [put]

func (i *Implementation) UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var input domain.Todo
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := i.todoService.Update(c.Context(), id, &input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// @Summary Удаление задачи
// @Description Удаляет задачу по ID
// @Tags tasks
// @Param id path int true "ID задачи"
// @Success 204
// @Failure 400 {object} fiber.Map "Некорректный ID"
// @Failure 500 {object} fiber.Map "Ошибка сервера"
// @Router /tasks/{id} [delete]

func (i *Implementation) DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	err := i.todoService.Delete(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
