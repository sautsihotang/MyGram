package router

import (
	"MyGram/controllers"
	"MyGram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
		userRouter.PUT("/:userId", middlewares.Authentication(), controllers.UpdateUser)
		userRouter.DELETE("/:userId", middlewares.Authentication(), controllers.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetPhotos)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comment")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", controllers.GetComments)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middlewares.Authentication())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaId", middlewares.SocialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	return r
}
