package router

import (
	"ewallet/database"
	"ewallet/handler"
	"ewallet/middleware"
	"ewallet/repository"
	"ewallet/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	db := database.NewDB()
	up := repository.NewUserRepo(db)
	tp := repository.NewTransactionRepo(db)
	wp := repository.NewWalletRepo(db)

	uu := usecase.NewUserUsecase(up, wp)
	tu := usecase.NewTransactionUsecase(tp, wp, up)

	uh := handler.NewUserHandler(uu)
	th := handler.NewTransactionHandler(tu)

	r := gin.Default()

	r.Static("/docs", "dist/")

	r.POST("/register", uh.Register)
	r.POST("/login", uh.Login)

	user := r.Group("/user")
	user.Use(middleware.JwtAuthMiddleware())
	{
		user.GET("/", uh.GetDetail)

	}

	transaction := r.Group("/transaction")
	transaction.Use(middleware.JwtAuthMiddleware())
	{
		transaction.GET("/", th.GetAll)

	}

	topup := r.Group("/topup")
	topup.Use(middleware.JwtAuthMiddleware())
	{
		topup.POST("/", th.TopUp)

	}

	transfer := r.Group("/transfer")
	transfer.Use(middleware.JwtAuthMiddleware())
	{
		transfer.POST("/", th.Transfer)

	}

	log.Println("Running HTTP server at 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal()
	}

}
