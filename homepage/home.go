package homepage

import (
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

const message = "Hello Niranjan"

type Handlers struct {
	logger *log.Logger
	db  *sqlx.DB
}

func (h *Handlers)HomeHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("processed Request")
	h.db.ExecContext(r.Context(),"")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte (message))
}

func NewHandlers(logger *log.Logger)*Handlers{
	return &Handlers{
		logger:logger,
	}

}
func (h *Handlers) Logger (next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter , r *http.Request){
		startTime:= time.Now()
		defer h.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

func (h *Handlers)SetUpRoutes(mux *http.ServeMux){
	mux.HandleFunc("/" , h.Logger(h.HomeHandler))

}