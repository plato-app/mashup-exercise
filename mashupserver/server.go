package mashupserver

import (
	"log"
	"net/http"

	"git.playchat.net/playchat/mashup-exercise/simpleserver"
)

// MashupServer is a simple server that combines 2 data sources together.
type MashupServer struct {
	*simpleserver.Server
}

// New creates a new mashup server that combines Source A and Source B.
func New() *MashupServer {
	srv, err := simpleserver.ServeAddr(":12000", []*simpleserver.PathHandler{
		{Path: "/result", Handler: handle}})
	if err != nil {
		log.Fatal(err)
	}
	return &MashupServer{srv}
}

// Does the actual work to combine the two data sources:  This is hard-coded, just to show as an example.
func handle(w http.ResponseWriter, r *http.Request) {
	// TODO: Fix me!
	w.Write([]byte("{\n   \"final_result\":511\n}\n"))
}
