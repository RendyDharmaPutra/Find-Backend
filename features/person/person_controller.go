package person

import (
	"Find-Backend/core/common"

	"github.com/gofiber/fiber/v2"
)

func PersonController(group fiber.Router, service Service) {
	group.Get("/", func (ctx *fiber.Ctx) error {
		persons, err := service.GetAllPersons()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mendapatkan data orang", err.Error()))
		}

		return ctx.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Berhasil mendapatkan data orang", persons))
	}) 

	group.Post("/", func (ctx *fiber.Ctx) error {
		var person Person
		if err := ctx.BodyParser(&person); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.NewFailedResponse("Gagal menambahkan orang", "gagal memproses data yang diberikan"))
		}

		if err := service.CreatePerson(&person); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal menambahkan orang", err.Error()))
		}

		return ctx.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse("Berhasil menambahkan orang"))
	})
}