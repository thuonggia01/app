package server

import (
	"github.com/gin-gonic/gin"
)

var (
// loginLogout authen.AuthenInterface
)

var (
// ForgotPassword fwPass.ForgotPwInterface
// resetPassword  rsPass.ResetPwInterface
)

func router(app *gin.Engine) {
	newHandler()
	// app.GET("/login", loginLogout.TemplateLogin)
	// app.POST("/login", loginLogout.LoginAuthen)
	// app.GET("/logout", loginLogout.Logout)

	// app.GET("/resetpass", resetPassword.TemplateResetPass)
	// app.POST("/resetpass", resetPassword.TemplateResetPass)

	// app.POST("/forgotpass", ForgotPassword.ForgotPassword)
}

func newHandler() {
	// loginLogout = authen.NewHandlerLogin()
	// ForgotPassword = fwPass.NewHandlerForgotPassword()
	// resetPassword = rsPass.NewHandlerRsPass()
}
