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
}