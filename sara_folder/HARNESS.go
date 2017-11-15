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
  ID string `json:"id,omitempty"`
  Firstname string `json:"firstname,omitempty"`
  Lastname string `json:"lastname,omitempty"`
  Medication string `json:"medication,omitempty"`
  Encounter string `json: "encounter,omitempty"`
  Condition string `json: "condition, omitempty"`
  Contact string `json: "contact,omitempty"`
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

//load the html and ask for ID
  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("askForID.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("ID: ", r.Form["ID"])
    person.ID = r.FormValue("ID")
  }
  //encode the person and save them into the local file
  json.NewEncoder(w).Encode(person)
  UpdateRecord(person)

//tried to retrieve FHIR data for the patients ID
  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Patient" + person.ID)
  r, err := http.NewRequest("GET", url, nil)
  r.Header.Set("Content-Type", "application/json")

  client := &http.Client{}

  resp := client.Do(r)
  defer resp.Body.Close()


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
  //

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
//loads the html file to display to the web browser
  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("Medication.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("Medication: ", r.Form["Medication"])
    medicationID = r.FormValue("Medication")
  }
//
//tried to retrieve FHIR information with user input
  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Medication" + medicationID)
  r, err := http.NewRequest("GET", url, nil)
  r.Header.Set("Content-Type", "application/json")

  client := &http.Client{}

  resp := client.Do(r)
  defer resp.Body.Close()

}
//////////////////////////////////////////////////
func LoadData(){
  file := ioutil.ReadFile("PatientInfo.json")
  json.Unmarshall([]byte(string(file)), &Patiens)
}
//////////////////////////////////////////////////////
func SaveData(Data []Patient){
  st := json.Marshal(Data)

  fo := os.Create("PatientInfo.json")
  defer func(){
    if err := fo.Close(); err != nil{
      panic(err)
    }
  }()
  fo.Write(str)
}
//////////////////////////////////////////////////////////
func UpdateRecord(Data Patient){
  Patients = append(Patients, Data)
  SaveData(Patients)
}
////////////////////////////////////////////////////////////
func GetCondition(w http.ResponseWriter, r *http.Request){
  var person Patient
  var condition string
//loads the html file to display to the web browser
  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("Condition.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("Condition: ", r.Form["Condition"])
    condition = r.FormValue("Condition")
  }
}
///////////////////////////////////////////////////////
func GetEncounter(w http.ResponseWriter, r *http.Request){
  var person Patient
  var encounter string
//loads the html file to display to the web browser
  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("Encounter.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("Encounter: ", r.Form["Encounter"])
    encounter = r.FormValue("Encounter")
  }
}
//////////////////////////////////////////////////////////////
func AddContact(w http.ResponseWriter, r *http.Request){
  var person Patient
  var person Patient
  var contact string
//loads the html file to display to the web browser
  fmt.Println("method: ", r.Method)
  if r.Method == "GET"{
    t, _ := template.ParseFiles("AddContact.html")
    t.Execute(w, nil)
  } else {
    r.ParseForm()
    fmt.Println("Contact: ", r.Form["Contact"])
    contact = r.FormValue("Contact")
  }

  person.Contact = contact
  json.NewEncoder(w).Encode(person)
  UpdateRecord(person)
}
func main(){
//all the http handles are going to be in this already
//also run the main menu for the client on a webpage
//this  server handles the HARNESS part and the data transfer
//for the clients
router := mux.NewRouter()
LoadData()

router.HandleFunc("/Patient", GetPatients).Methods("GET")
router.HandleFunc("/log", UserLogin).Methods("GET")
router.HandleFunc("/lkuID", GetPatientByID).Methods("GET")
router.HandleFunc("/lku", GetPatientByName).Methods("GET")
router.HandleFunc("/med", GetMedication).Methods("GET")
router.HandleFunc("/cond", GetCondition).Methods("GET")
router.HandleFunc("/enc", GetEncounter).Methods("GET")
router.HandleFunc("/rmc", RemoveContact).Methods("DELETE")
router.HangleFunc("/add", AddContact).Methods("POST")

http.ListenAndServe(":9090", router)


}
