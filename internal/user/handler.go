package user

import (
	"net/http"
	"todolist/internal/handlers"

	"github.com/julienschmidt/httprouter"
)

const (
	listsURL    = "/list"
	listItemURL = "/list/:item"
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(listsURL, h.GetList)
	router.POST(listItemURL, h.AddItem)
	router.DELETE(listItemURL, h.DeleteItem)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is Todolist"))
}
func (h *handler) AddItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is Add item"))
}
func (h *handler) DeleteItem(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Write([]byte("this is delete item"))
}
