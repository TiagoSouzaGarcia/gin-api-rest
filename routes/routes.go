package routes

import (
	"github.com/TiagoSouzaGarcia/api-go-gin/controllers"
	docs "github.com/TiagoSouzaGarcia/api-go-gin/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/alunos", controllers.ExibirTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.CriarNovoAluno)
	r.GET("/alunos/:id", controllers.BuscarAlunoPorId)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoPorCpf)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/index", controllers.ExibePaginaIndex)

	r.NoRoute(controllers.RotaNaoEncontrada)

	r.Run()
}
