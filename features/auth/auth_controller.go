package auth

import (
	"Find-Backend/core/common"
	"Find-Backend/features/user"

	"github.com/gofiber/fiber/v2"
)

func AuthController(group fiber.Router, service Service) {
	group.Post("/register", func (ctx *fiber.Ctx) error {
		var user user.User
		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.NewFailedResponse("Gagal mendaftarkan pengguna", "gagal memproses data yang diberikan"))
		}

		if err := service.Register(user.Fullname, user.Username, user.Password); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mendaftarkan pengguna", err.Error()))
		}

		return ctx.Status(fiber.StatusCreated).JSON(common.NewSuccessResponse("Berhasil mendaftarkan pengguna"))
	})

	group.Post("/login", func (ctx *fiber.Ctx) error  {
		var user user.User
		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(common.NewFailedResponse("Gagal mengautentikasi pengguna pengguna", "gagal memproses data yang diberikan"))
		}

		status, token, err := service.Login(user.Username, user.Password)
		if err != nil {
			return ctx.Status(status).JSON(common.NewFailedResponse("Gagal mengautentikasi pengguna", err.Error()))
		}
		
		return ctx.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Berhasil mengautentikasi pengguna", map[string]interface{}{"token": token}))
	})
	
	
	group.Get("/logout", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Berhasil melakukan logout"))
	})
	
	group.Get("/profile", func (ctx *fiber.Ctx) error  {
		id := ctx.Locals("id")
		if id == nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mendapatkan profil pengguna", "Id tidak ditemukan"))			
		}	

		user, err := service.Profile(id)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(common.NewFailedResponse("Gagal mengautentikasi pengguna", err.Error()))			
		}
	
		return ctx.Status(fiber.StatusOK).JSON(common.NewSuccessResponse("Berhasil mendapatkan profil pengguna", map[string]interface{}{"profile": user}))
	})
}