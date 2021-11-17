package server

import "github.com/gin-gonic/gin"

func Route() error {
	app := setup()
	return app.Run()
}

func setup() *gin.Engine {
	app := gin.New()
	app.Static("/assets", "./assets")
	//app.SetFuncMap(viewhelper.Funcs())
	app.LoadHTMLGlob("assets/templates/**/*.html")
	router(app)
	return app
}
