package api

import (
	"log"
	"net/http"

	js "github.com/AjanShrestha/hexagonal-microservices-go/serializer/json"
	ms "github.com/AjanShrestha/hexagonal-microservices-go/serializer/msgpack"
	"github.com/AjanShrestha/hexagonal-microservices-go/shortener"
	"github.com/go-chi/chi"
	"github.com/pkg/errors"
)

// RedirectHandler interface defines the methods for http layer
type RedirectHandler interface {
	Get(http.ResponseWriter, *http.Request)
	Post(http.ResponseWriter, *http.Request)
}

// handler struct for RedirectService
type handler struct {
	redirectService shortener.RedirectService
}

// NewHandler takes RedirectService and returns back RedirectHandler
func NewHandler(redirectService shortener.RedirectService) RedirectHandler {
	return &handler{redirectService: redirectService}
}

// setupResponse to setup response header for different serializers
func setupResponse(w http.ResponseWriter, contentType string, body []byte, statusCode int) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	_, err := w.Write(body)
	if err != nil {
		log.Println(err)
	}
}

// serializer maps contenttype to right serializer
func (h *handler) serializer(contentType string) shortener.RedirectSerializer {
	if contentType == "application/x-msgpack" {
		return &ms.Redirect{}
	}
	return &js.Redirect{}
}

// Get redirects to the url that is stored in the database
func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	redirect, err := h.redirectService.Find(code)
	if err != nil {
		if errors.Cause(err) == shortener.ErrRedirectNotFound {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	http.Redirect(w, r, redirect.URL, http.StatusMovedPermanently)
}

// Post stores the code to be used for redirect
func (h *handler) Post(w http.ResponseWriter, r *http.Request) {

}
