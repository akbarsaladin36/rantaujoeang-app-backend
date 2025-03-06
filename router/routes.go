package router

import (
	"os"
	"rantaujoeang-app-backend/controllers"
	"rantaujoeang-app-backend/database"
	"rantaujoeang-app-backend/middleware"
	"rantaujoeang-app-backend/repositories"
	"rantaujoeang-app-backend/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ConnectRoutes() {
	router := gin.Default()

	router.Use(cors.Default())

	authRepositories := repositories.NewAuthRepository(database.DB)
	authServices := services.NewAuthService(authRepositories)
	authControllers := controllers.NewAuthController(authServices)

	userRepositories := repositories.NewUserRepository(database.DB)
	userServices := services.NewUserService(userRepositories)
	userControllers := controllers.NewUserController(userServices)

	profileRepositories := repositories.NewProfileRepository(database.DB)
	profileServices := services.NewProfileService(profileRepositories)
	profileControllers := controllers.NewProfileController(profileServices)

	btRepositories := repositories.NewBalanceTransactionRepository(database.DB)
	btServices := services.NewBalanceTransactionService(btRepositories)
	btControllers := controllers.NewBalanceTransactionController(btServices)

	dormRepositories := repositories.NewDormRepository(database.DB)
	dormServices := services.NewDormService(dormRepositories)
	dormControllers := controllers.NewDormController(dormServices)

	paymentRepositories := repositories.NewPaymentRepository(database.DB)
	paymentServices := services.NewPaymentService(paymentRepositories)
	paymentControllers := controllers.NewPaymentController(paymentServices)

	invoiceRepositories := repositories.NewInvoiceRepository(database.DB)
	invoiceServices := services.NewInvoiceService(invoiceRepositories)
	invoiceControllers := controllers.NewInvoiceController(invoiceServices)

	messageRepositories := repositories.NewMessageRepository(database.DB)
	messageServices := services.NewMessageService(messageRepositories)
	messageControllers := controllers.NewMessageController(messageServices)

	v1 := router.Group("/rantaujoeang-app/v1")

	v1.POST("/auth/register", authControllers.RegisterController)
	v1.POST("/auth/login", authControllers.LoginController)

	// Admin role
	v1Admin := router.Group("/rantaujoeang-app/v1/admin").Use(middleware.AuthMiddleware(), middleware.IsAdminAccess())

	v1Admin.GET("/users", userControllers.FindUsersController)
	v1Admin.GET("/users/detail-user/:username", userControllers.FindUserController)
	v1Admin.POST("/users", userControllers.CreateUserController)
	v1Admin.PATCH("/users/detail-user/:username", userControllers.UpdateUserController)
	v1Admin.DELETE("/users/detail-user/:username", userControllers.DeleteUserController)

	v1Admin.GET("/profile", profileControllers.FindProfileController)
	v1Admin.PATCH("/profile", profileControllers.UpdateProfileController)

	v1Admin.GET("/balance-transactions", btControllers.FindBalanceTransactionsController)
	v1Admin.GET("/balance-transactions/detail-balance-transaction/:balance_transaction_code", btControllers.FindBalanceTransactionController)
	v1Admin.POST("/balance-transactions", btControllers.CreateBalanceTransactionController)

	v1Admin.GET("/dorms", dormControllers.FindDormsController)
	v1Admin.GET("/dorms/detail-dorm/:dorm_code", dormControllers.FindDormController)
	v1Admin.POST("/dorms", dormControllers.CreateDormController)
	v1Admin.PATCH("/dorms/detail-dorm/:dorm_code", dormControllers.UpdateDormController)
	v1Admin.DELETE("/dorms/detail-dorm/:dorm_code", dormControllers.DeleteDormController)

	v1Admin.GET("/payments", paymentControllers.FindPaymentsController)
	v1Admin.GET("/payments/detail-payment/:payment_code", paymentControllers.FindPaymentController)
	v1Admin.PATCH("/payments/detail-payment/:payment_code", paymentControllers.UpdatePaymentController)

	v1Admin.GET("/invoices", invoiceControllers.FindInvoicesController)
	v1Admin.GET("/invoices/detail-invoice/:invoice_code", invoiceControllers.FindInvoiceController)
	v1Admin.POST("/invoices", invoiceControllers.CreateInvoiceController)
	v1Admin.PATCH("/invoices/detail-invoice/:invoice_code", invoiceControllers.UpdateInvoiceController)
	v1Admin.DELETE("/invoices/detail-invoice/:invoice_code", invoiceControllers.DeleteInvoiceController)

	v1Admin.GET("/messages", messageControllers.FindMessagesController)
	v1Admin.GET("/messages/detail-message/:message_code", messageControllers.FindMessageController)
	v1Admin.POST("/messages", messageControllers.CreateMessageController)
	v1Admin.POST("/messages/detail-message/reply-message/:message_code", messageControllers.ReplyMessageController)
	v1Admin.PATCH("/messages/detail-message/:message_code", messageControllers.UpdateMessageController)
	v1Admin.DELETE("/messages/detail-message/:message_code", messageControllers.DeleteMessageController)

	// User Role
	v1User := router.Group("/rantaujoeang-app/v1/user").Use(middleware.AuthMiddleware(), middleware.IsUserAccess())

	v1User.GET("/profile", profileControllers.FindProfileController)
	v1User.PATCH("/profile", profileControllers.UpdateProfileController)

	v1User.GET("/balance-transactions/my-balance-transactions", btControllers.FindBalanceTransactionsByUserIdController)
	v1User.GET("/balance-transactions/detail-balance-transaction/:balance_transaction_code", btControllers.FindBalanceTransactionController)
	v1User.POST("/balance-transactions/my-balance-transactions", btControllers.CreateBalanceTransactionController)

	v1User.GET("/dorms", dormControllers.FindDormsController)
	v1User.GET("/dorms/detail-dorm/:dorm_code", dormControllers.FindDormController)

	v1User.GET("/payments/my-payments", paymentControllers.FindPaymentsByUserIdController)
	v1User.GET("/payments/detail-payment/:payment_code", paymentControllers.FindPaymentController)
	v1User.POST("/payments/my-payments", paymentControllers.CreatePaymentController)

	v1User.GET("/invoices/my-invoices", invoiceControllers.FindInvoicesByUserIdController)
	v1User.GET("/invoices/detail-invoice/:invoice_code", invoiceControllers.FindInvoiceController)
	v1User.POST("/invoices/my-invoices", invoiceControllers.CreateInvoiceController)

	v1User.GET("/messages/my-messages", messageControllers.FindMessagesByUserIdController)
	v1User.GET("/messages/detail-message/:message_code", messageControllers.FindMessageController)
	v1User.POST("/messages/my-messages", messageControllers.CreateMessageController)
	v1User.POST("/messages/detail-message/reply-message/:message_code", messageControllers.ReplyMessageController)
	v1User.PATCH("/messages/detail-message/:message_code", messageControllers.UpdateMessageController)
	v1User.DELETE("/messages/detail-message/:message_code", messageControllers.DeleteMessageController)

	router.Run(os.Getenv("APP_PORT"))
}
