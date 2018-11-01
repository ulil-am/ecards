package http

import (
	"net/http"

	"github.com/astaxie/beego"
)

// pageNotFound ..
func pageNotFound(rw http.ResponseWriter, r *http.Request) {
	_, err := rw.Write([]byte(""))
	if err != nil {
		beego.Warning("NOT FOUND ERROR")
	}
}
