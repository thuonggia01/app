package sessions

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
var client string

func NewCookie(c *gin.Context, name string, email string, avatar string) {
	client = "client"
	session, err := store.Get(c.Request, client)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	// Tạo Thời Gian mà cookie sử dụng và cusstom
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
		// MaxAge:   5,
		HttpOnly: true,
	}
	// Set name session value.
	session.Values["fullName"] = name
	session.Values["email"] = email
	session.Values["avatar"] = avatar
	log.Println("session:", session)
	// Save it before we write to the response/return from the handler.
	// func (s *Session) Save(r *http.Request, w http.ResponseWriter) error
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

}

func ClearSession(c *gin.Context) {
	session, err := store.Get(c.Request, client)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
	}
	delete(session.Values, "fullName")
	delete(session.Values, "email")
	delete(session.Values, "avatar")
	err = session.Save(c.Request, c.Writer)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSession() *sessions.CookieStore {
	return store
}
