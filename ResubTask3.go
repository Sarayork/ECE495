package main

import(
  "encoding/json"
  //"log"
  "net/http"
  "github.com/gorilla/mux"
)

type Person struct{
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Address *Address `json:"address,omitempty"`
  Contact string `json: "contact,omitempty"`

}

type Address struct{
  City string `json:city,omitempty"`
  State string `json:"state,omitempty"`
}

var people []Person

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
  json.NewEncoder(w).Encode(people)
}

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for _, item := range people{
    if item.ID == params["id"]{
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&Person{})//return empty person if no one

}//end GetPersonEndpoint

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  var person Person
  _ = json.NewDecoder(req.Body).Decode(&person)
  person.ID = params["id"]
  person.Firstname = params["firstname"]
  person.Lastname = params["lastname"]
  person.Contact = params["contact"]
  people = append(people, person)
  json.NewEncoder(w).Encode(people)
}//end create

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for index, item := range people{
    if item.ID == params["id"]{
      people = append(people[:index], people[index+1:]...)
      break
    }
  }
  json.NewEncoder(w).Encode(people)
}//end delete person

func LookUpName(w http.ResponseWriter, req *http.Request){
  params := mux.Vars(req)
  for _, item := range people{
    if item.Firstname == params["firstname"]{
      json.NewEncoder(w).Encode(item)
      return
    }/* else if item.ID == params["id"]{
      json.NewEncoder(w).Encode(item)
      return
    }*/
  }
  json.NewEncoder(w).Encode(&Person{})
}//end find by common name
 func LookUpID(w http.ResponseWriter, req *http.Request){
   params := mux.Vars(req)
   for _, item := range people{
     if item.ID == params["id"]{
       json.NewEncoder(w).Encode(item)
       return
     }
   }
   json.NewEncoder(w).Encode(&Person{})
 }//end look up

 func AddContact(w http.ResponseWriter, req *http.Request){
   params := mux.Vars(req)
   var person Person
   _ = json.NewDecoder(req.Body).Decode(&person)
   if person.ID == params["id"]{
     person.Contact = params["contact"]
   }


  json.NewEncoder(w).Encode(people)

 }//end add contact
 func RemoveContact(w http.ResponseWriter, req *http.Request){
   params := mux.Vars(req)
   for index, item := range people{
     if item.Contact == params["contact"]{
       people = append(people[:index], people[index+1:]...)
       break
     }
   }
   json.NewEncoder(w).Encode(people)
 }
func main(){
  router := mux.NewRouter()

  //make some already people
  people = append(people, Person{ID:"1", Firstname: "Sara"})
  people = append(people, Person{ID:"2", Firstname: "Baby"})
  people = append(people, Person{ID:"3", Firstname: "Cookie"})
  people = append(people, Person{ID:"4", Firstname: "Stuart"})

  router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
  router.HandleFunc("/people/sub/{id}", CreatePersonEndpoint).Methods("POST")
  router.HandleFunc("/people/sub/{id}/{firstname}", CreatePersonEndpoint).Methods("POST")
//  router.HandleFunc("/people/lku/{id}", GetPersonEndpoint).Methods("GET")
  router.HandleFunc("/people/UID/{id}", LookUpID).Methods("GET")
  router.HandleFunc("/people/NAME/{firstname}", LookUpName).Methods("GET")
  router.HandleFunc("/people/Add/{id}/{contact}", AddContact).Methods("POST")
  router.HandleFunc("/people/Rmc/{id}/{contact}", RemoveContact).Methods("POST")

  http.ListenAndServe(":7000", router)
}
