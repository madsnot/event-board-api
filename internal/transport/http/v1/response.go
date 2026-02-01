package v1

import (
	"encoding/json"
	"errors"
	"github.com/madsnot/event-board-api/internal/domain/models"
	"github.com/madsnot/event-board-api/internal/domain/usecase"
	"log"
	"net/http"
)

const (
	ContentTypeHeader                   = "Content-Type"
	ApplicationJSONType                 = "application/json"
	AccessControlAllowCredentialsHeader = "Access-Control-Allow-Credentials"
	AccessControlAllowOriginHeader      = "Access-Control-Allow-Origin"
	AccessControlAllowHeadersHeader     = "Access-Control-Allow-Headers"
	AccessControlAllowMethodsHeader     = "Access-Control-Allow-Methods"
	AccessControlAllowCredentialsValue  = "true"
	AccessControlAllowOriginValue       = "*"
	AccessControlAllowHeadersValue      = "Origin, X-Requested-With, Content-Type, Accept"
	AccessControlAllowMethodsValue      = "GET,POST,OPTIONS,DELETE,PUT"
)

func writeError(w http.ResponseWriter, err error) {
	var (
		wErr models.WrapError
		bErr usecase.BusinessError
	)

	w.Header().Set(ContentTypeHeader, ApplicationJSONType)

	switch {
	case errors.Is(err, wErr):
		log.Println(wErr)

		errors.As(err, &bErr)

		w.WriteHeader(CODES[bErr.Code()])
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}

	respBody, err := json.Marshal(err.Error())
	if err == nil {
		_, _ = w.Write(respBody)
	}
}

func writeOK(w http.ResponseWriter, body interface{}) {
	response, err := json.Marshal(body)
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set(AccessControlAllowCredentialsHeader, AccessControlAllowCredentialsValue)
	w.Header().Set(AccessControlAllowOriginHeader, AccessControlAllowOriginValue)
	w.Header().Set(AccessControlAllowHeadersHeader, AccessControlAllowHeadersValue)
	w.Header().Set(AccessControlAllowMethodsHeader, AccessControlAllowMethodsValue)
	w.Header().Set(ContentTypeHeader, ApplicationJSONType)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
}
