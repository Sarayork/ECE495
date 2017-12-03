package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"fmt"
)

type Patient struct{
	ID string `json:"id"`
	Name []Name `json:"name"`
	Telecom []Telecom `json:"telecom"`
	Gender string `json:"gender"`
	BirthDate string `json:"birthDate"`
}

type Name struct {
	Family string `json:"family, omitempty"`
	Given []string `json:"given"`
	Prefix []string `json:"prefix"`
}

type Telecom struct{
	System string `json:"system"`
	Value string `json:"value"`
	Use string `json:"use"`
}

func LookUpID(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	ID := params["id"]
	//for _, item := range people{
	//	if item.ID == params["id"]{
	//		json.NewEncoder(w).Encode(item)
	//		return
	//	}
	//}
	//json.NewEncoder(w).Encode(&Person{})
	url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Patient/" + ID)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	if err != nil{
		log.Fatal("Do: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil{
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()

	var person Patient

	if err := json.NewDecoder(resp.Body).Decode(&person); err != nil {
		log.Println(err)
	}

	fmt.Println("Family = ", person.Name[0])
	json.NewEncoder(w).Encode(person)

	//if resp.StatusCode == 200 {
	//
	//	responseData, err := ioutil.ReadAll(resp.Body)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	//responseString := string(responseData)
	//	//_, _ = json.Marshal(responseString)
	//	//return
	//	json.NewEncoder(w).Encode(responseString)
	//}
}//end look up

func main(){
	router := mux.NewRouter()

	//make some already people
	//people = append(people, Person{ID:"1", Firstname: "Sara"})
	//people = append(people, Person{ID:"2", Firstname: "Baby"})
	//people = append(people, Person{ID:"3", Firstname: "Cookie"})
	//people = append(people, Person{ID:"4", Firstname: "Stuart"})

	//router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	//router.HandleFunc("/people/sub/{id}", CreatePersonEndpoint).Methods("POST")
	//router.HandleFunc("/people/sub/{id}/{firstname}", CreatePersonEndpoint).Methods("POST")
	//  router.HandleFunc("/people/lku/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", LookUpID).Methods("GET")
	//router.HandleFunc("/people/NAME/{firstname}", LookUpName).Methods("GET")
	//router.HandleFunc("/people/Add/{id}/{contact}", AddContact).Methods("POST")
	//router.HandleFunc("/people/Rmc/{id}/{contact}", RemoveContact).Methods("POST")

	http.ListenAndServe(":7000", router)
}
