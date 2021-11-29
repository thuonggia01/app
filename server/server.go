package server

import "github.com/gin-gonic/gin"

func Route() error {
	app := setup()
	return app.Run(":3000")
}

func setup() *gin.Engine {
	app := gin.Default()
	app.Static("/assets", "./assets")
	//app.SetFuncMap(viewhelper.Funcs())
	app.LoadHTMLGlob("templates/**/*.html")
	router(app)
	return app
}
