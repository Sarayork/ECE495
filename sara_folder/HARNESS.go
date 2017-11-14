package main

import(
  "encoding/json"
  "log"
  "net/http"
  "fmt"
  "github.com/gorilla/mux"
  //"io/ioutil"
  //"os"
  "html/template"

)


type Patient struct{

  ID string `json: "id,omitempty"`
  Name string `json: "patname,omitempty"`
  Telecom string  `json: "telecom"`
  Gender string `json: "gender"`
  BirthDate string `json:"birthDate"`

	Family string `json:"family, omitempty"`
	Given string `json:"given"`
	Prefix string `json:"prefix"`

	System string `json:"system"`
	Value string `json:"value"`
	Use string `json:"use"`
}
//create array of patients that are type patient from the struct
var Patients []Patient

//retrieve all the people function//////////////////////////////////////////
func GetPatients(w http.ResponseWriter, req *http.Request){

  //LoadData()

  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Patient")
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Content-Type", "application/json")
  //SaveData(Patients)
  client := &http.Client{}

  //error control
  if err != nil{
    log.Fatal("Do: ", err)
    return
  }

  resp, err := client.Do(req)
//error
  if err != nil{
    log.Fatal("Do: ", err)
    return
  }

  //if ok defer the close of the Body
  defer resp.Body.Close()


  json.NewEncoder(w).Encode(Patients)

}//end of GetPatients
//##############################################################
//login for the user
func UserLogin(w http.ResponseWriter, req *http.Request){
  //params := mux.Vars(req)
  //var person Patient
  fmt.Println("method: ", req.Method)//get request Method
  if req.Method == "GET" {
    t, _ := template.ParseFiles("login.html")
    t.Execute(w, nil)
  } else {
    req.ParseForm()
    fmt.Println("User ID: ", req.Form["User ID"])
    fmt.Println("Password: ", req.Form["Password"])
  }


  _ = json.NewDecoder(req.Body).Decode(&Patients)

}
//###############################################
func GetPatientByID(w http.ResponseWriter, r *http.Request){

  //var ID string
  var person Patient

  _ = json.NewDecoder(r.Body).Decode(&person)

  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("askForID.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("ID: ", r.Form["ID"])
    person.ID = r.FormValue("ID")
  }

  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Patient" + person.ID)
  r, err := http.NewRequest("GET", url, nil)
  r.Header.Set("Content-Type", "application/json")

  client := &http.Client{}

  resp := client.Do(r)
  defer resp.Body.Close()

  json.NewEncoder(w).Encode(Patients)

}//end of look up by ID
/////////////////////////////////////////////////////////////////////////
func GetPatientByName(w http.ResponseWriter, r *http.Request){
  //var Name string
  var person Patient

  _ = json.NewDecoder(r.Body).Decode(&person)

  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("askForName.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("Patient: ", r.Form["Patient"])
    person.Name = r.FormValue("Patient")
  }

  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Patient" + person.Name)
  r, err := http.NewRequest("GET", url, nil)
  r.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  resp := client.Do(r)
  defer resp.Body.Close()

  json.NewEncoder(w).Encode(Patients)

}
////////////////////////////////////////////////////////////////////////////
func GetMedication(w http.ResponseWriter, r *http.Request){

  var person Patient
  var medicationID string

  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("Medication.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("Medication: ", r.Form["Medication"])
    medicationID = r.FormValue("Medication")
  }

  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Medication" + medicationID)
  r, err := http.NewRequest("GET", url, nil)
  r.Header.Set("Content-Type", "application/json")

  client := &http.Client{}

  resp := client.Do(r)
  defer resp.Body.Close()


}
func main(){
//all the http handles are going to be in this already
//also run the main menu for the client on a webpage
//this  server handles the HARNESS part and the data transfer
//for the clients
router := mux.NewRouter()

router.HandleFunc("/Patient", GetPatients).Methods("GET")
router.HandleFunc("/log", UserLogin).Methods("GET")
router.HandleFunc("/lkuID", GetPatientByID).Methods("GET")
router.HandleFunc("/lku", GetPatientByName).Methods("GET")
router.HandleFunc("/med", GetMedication).Methods("GET")

http.ListenAndServe(":9090", router)


}
