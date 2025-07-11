package raw

import (
	"filter-service/request"
	"filter-service/response"
	"net/http"
)

type RawHandler struct {
	*RawService
}

type RawHandlerDeps struct {
	*RawService
}

func NewRawHandler(router *http.ServeMux, deps RawHandlerDeps) {
	rawHandler := RawHandler{
		RawService: deps.RawService,
	}
	router.HandleFunc("POST /raw", rawHandler.Put())

}

func (handler *RawHandler) Put() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RawRequest](&w, r)
		if err != nil {
			return
		}
		err = handler.RawService.SaveToBD(body.Data, body.Valid)
		if err != nil {
			response.Json(w, err, http.StatusInternalServerError)
		}
		response.Json(w, "success", http.StatusOK)
	}
}
