package main

import(
  "encoding/json"
  "log"
  "net/http"
  "fmt"
  "github.com/gorilla/mux"
  "io/ioutil"
  "os"

)


type Patient struct{

  UID string `json: "id,omitempty"`
  PatName string `json: "patname,omitempty"`

}
//create array of patients that are type patient from the struct
var Patients []Patient

//retrieve all the people function
func GetPatients(w http.ResponseWriter, req *http.Request){

  //LoadData()
  
  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Patient/")
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Content-Type", "application/json")
  SaveData(Patients)
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
  if resp.StatusCode == 200 {
    defer resp.Body.Close()
  }else {
    fmt.Println("error")
  }

  json.NewEncoder(w).Encode(Patients)

}//end of GetPatients
//load the data frim the fhir server

func LoadData(){
  file, err :=ioutil.ReadFile("PatientInfo.json")
  if err != nil {
    panic(err)
  }
  json.Unmarshal([]byte(string(file)), &Patients)
}

func SaveData(Data []Patient){

  str, err := json.Marshal(Data)

  if err != nil {
    fmt.Println("error")
    return
  }

  fo, err := os.Create("PatientInfo.json")

  if err != nil {
    panic(err)
  }

  defer func(){
    if err := fo.Close(); err != nil {
      panic(err)
    }
  }()
  fo.Write(str)
}//save the Patients

func main(){
//all the http handles are going to be in this already
//also run the main menu for the client on a webpage
//this  server handles the HARNESS part and the data transfer
//for the clients
router := mux.NewRouter()

LoadData()
router.HandleFunc("/Patient", GetPatients).Methods("GET")

http.ListenAndServe(":9090", router)


}
