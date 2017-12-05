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

const DBNAME = "harness"          // for database name
const USERSCOLL = "users"         // for users collection
const RELATIONSCOLL = "relations" // for relations collection

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session, USERSCOLL, []string{"fhirid"})
	//ensureIndex(session, RELATIONSCOLL, []string{"srcuid", "destuid", "srcrole", "destrole"})

	mux := goji.NewMux()
	mux.HandleFunc(pat.Post("/user"), subscribe(session))
	mux.HandleFunc(pat.Get("/user/getUid/:name"), lookUpUid(session))
	mux.HandleFunc(pat.Get("/user/getIp/:uid"), lookUpIp(session))
	mux.HandleFunc(pat.Put("/beatHeart/:uid"), beatHeart(session))

	mux.HandleFunc(pat.Post("/relation/:srcUid/:destUid/:srcRole/:destRole"), addRelation(session))
	mux.HandleFunc(pat.Get("/invitations/:uid"), lookUpInvitations(session))
	mux.HandleFunc(pat.Put("/relation/:srcUid/:destUid/:srcRole/:destRole"), acceptRelation(session))
	mux.HandleFunc(pat.Delete("/relation/:srcUid/:destUid/:srcRole/:destRole"), removeRelation(session))

	mux.HandleFunc(pat.Post("/notification"), addMsgNotif(session))
	mux.HandleFunc(pat.Get("/notifications/:uid"), lookUpMsgNotif(session))

	http.ListenAndServe("localhost:12345", mux)
}

/**
	Make fields pk as indexes in the given collection
 */
func ensureIndex(s *mgo.Session, collection string, pk []string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(DBNAME).C(collection)

	var keys []string
	for _, k := range pk {
		keys = append(keys, "$text:" + k)
	}
	//for _, p := range pk {
		index := mgo.Index{
			Key:        keys,
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}

		err := c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}
	//}
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
	Invoked function when return success in json format body
 */
func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	ResponseWithNoContent(w, code)
	w.Write(json)
}

/**
	Invoked function when return success with no content
 */
func ResponseWithNoContent(w http.ResponseWriter, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
}