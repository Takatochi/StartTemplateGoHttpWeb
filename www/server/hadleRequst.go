package server

import (
	"fmt"
	"net/http"
	"text/template"
)

type Server struct {
	_handle    string
	handlename string
}

func Init(_handle string, handlename string) *Server {
	return &Server{
		_handle:    _handle,
		handlename: handlename,
	}

}
func (s *Server) Handle() {
	http.Handle(s._handle, http.StripPrefix(s._handle, http.FileServer(http.Dir("."+s._handle))))

}

func (s *Server) Request() {
	http.HandleFunc(s.handlename, s.index)

}
func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	//старт темлейтов для метода index (головна сторінка)
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

	if err != nil {
		fmt.Println(w, err.Error())
	}
	t.ExecuteTemplate(w, "index", nil)
}
