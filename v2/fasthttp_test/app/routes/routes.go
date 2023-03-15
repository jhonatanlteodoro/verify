package routes

import (
	"github.com/jhonatanlteodoro/fasthttp_test/app/handlers"
	routing "github.com/qiangxue/fasthttp-routing"
)

func userRoutes(baseRouter *routing.Router) {
	router := baseRouter.Group("/users")

	router.Get("/", handlers.HanlderRepo.GetAllUsers)
	router.Get("/<userID>", handlers.HanlderRepo.GetUserById)
	router.Post("/", handlers.HanlderRepo.CreateUser)
	router.Patch("/<userID>", handlers.HanlderRepo.UpdateUser)
	router.Delete("/<userID>", handlers.HanlderRepo.DeleteUserById)

}

func RegistryRoutes(router *routing.Router) {

	userRoutes(router)
}
