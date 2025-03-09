package person

import (
	"Find-Backend/core/common"

	"github.com/gofiber/fiber/v2"
)

func PersonController(group fiber.Router, service Service) {
	group.Get("/", func (ctx *fiber.Ctx) error {
		id := ctx.Locals("id")
		if id == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(common.NewFailedResponse("Gagal mendapatkan data orang", "Id tidak ditemukan"))			
		}

		persons, err := service.GetAllPersons(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mendapatkan data orang", err.Error()))
		}

		return ctx.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Berhasil mendapatkan data orang", map[string]interface{}{"persons": persons}))
	}) 

	group.Post("/", func (ctx *fiber.Ctx) error {
		id := ctx.Locals("id")
		if id == nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(common.NewFailedResponse("Gagal menambahkan orang", "Id tidak ditemukan"))			
		}

		var person Person
		if err := ctx.BodyParser(&person); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.NewFailedResponse("Gagal menambahkan orang", "gagal memproses data yang diberikan"))
		}

		if err := service.CreatePerson(&person, id); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal menambahkan orang", err.Error()))
		}

		return ctx.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse("Berhasil menambahkan orang"))
	})
}