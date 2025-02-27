package user

import (
	"Find-Backend/core/common"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

func UserController(group fiber.Router, service Service) {
	group.Get("/", func(ctx *fiber.Ctx) error {
		users, err := service.GetAllUsers()
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mendapatkan data pengguna", err.Error()))
		}

		return ctx.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Berhasil mendapatkan data pengguna", users))
	})

	group.Get("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		uuidRegex := regexp.MustCompile(`^[a-f0-9]{8}-[a-f0-9]{4}-[1-5][a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}$`)
    	if !uuidRegex.MatchString(id) {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.NewFailedResponse("Gagal mendapatkan data pengguna", "format ID tidak valid"))
    	}

		user, err := service.GetUser(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mendapatkan data pengguna",  err.Error()))
		}

		return ctx.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Berhasil mendapatkan data pengguna", user))
	})

	group.Post("/", func (ctx *fiber.Ctx) error {
		var request struct {
			Fullname string `json:"fullname"`
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := ctx.BodyParser(&request); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.NewFailedResponse("Gagal mendaftarkan pengguna",  "hagal memproses data yang diberikan"))
		}

		if err := service.CreateUser(request.Fullname, request.Username, request.Password); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mendaftarkan pengguna", err.Error()))
		}

		return ctx.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse("Berhasil mendaftar pengguna"))
	})
}