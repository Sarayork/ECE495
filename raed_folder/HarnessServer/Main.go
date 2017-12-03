package main

import(
	//"encoding/json"
	"fmt"
	//"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session, "users", "fhirid")

	mux := goji.NewMux()
	mux.HandleFunc(pat.Post("/user"), subscribe(session))
	mux.HandleFunc(pat.Get("/user/getUid/:name"), lookUpUid(session))
	mux.HandleFunc(pat.Get("/user/getIp/:uid"), lookUpIp(session))
	mux.HandleFunc(pat.Put("/beatHeart/:uid"), beatHeart(session))

	mux.HandleFunc(pat.Post("/invitation/:srcUid/:destUid/:srcRole/:destRole"), addInvitation(session))
	mux.HandleFunc(pat.Get("/invitations/:uid"), lookUpInvitations(session))
	mux.HandleFunc(pat.Delete("/invitation/:srcUid/:destUid/:srcRole/:destRole"), removeInvitation(session))

	mux.HandleFunc(pat.Post("/notification"), addMsgNotif(session))
	mux.HandleFunc(pat.Get("/notifications/:uid"), lookUpMsgNotif(session))

	http.ListenAndServe("localhost:12345", mux)
}

func ensureIndex(s *mgo.Session, collection string, pk string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(DBNAME).C(collection)

	index := mgo.Index{
		Key:        []string{pk},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

/**
	Invoked function when return error in json format
 */
func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{error: %q}", message)
}

/**
	Invoked function when return success in json format
 */
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}