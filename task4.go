package main

import(
  "encoding/json"
  "log"
  "net/http"
  "fmt"
  "io/ioutil"
  "os"

)

type Patient struct{
  ID string  `json: "id,omitempty"`
  Firstname string `json: "firstname,omitempty"`
  Lastname string `json: "lastname,omitempty"`
  Encounter string `json: "encounter,omitempty"`
  Medication string `json: "medication,omitempty"`
  Condition string  `json: "condition,omitempty"`

}
//
var Patients []Patient

func main(){

  //var endconnection bool = false
  var input int

  for{
    LoadData()
    fmt.Println("Choose a task")
    fmt.Println("1: Get patient information ")
    fmt.Println("2: Get patient encounter")
    fmt.Println("3: Get patient medications")
    fmt.Println("4: Get conditions")
    fmt.Println("5: Exit service")
    fmt.Scanln(&input)

    if input == 1 {
      GetInfo()
    }else if input == 2{
      //GetEncounter()
    }else if input == 3 {
      //GetMeds()
    }else if input == 4{
      //GetCondition()
    }else {

      return
    }
  }//end while loop
}

func LoadData(){
  file, err := ioutil.ReadFile("PatientInfo.json")
  if err != nil{
    panic(err)
  }
  json.Unmarshal([]byte(string(file)), &Patients)
}//LoadData

func UpdateLocalRecords(Data Patient){
  Patients = append(Patients, Data)
  SaveData(Patients)
}

func SaveData(Data []Patient){
  str, err := json.Marshal(Data)
  if err != nil{
    fmt.Println("error")
    return
  }

  fo, err := os.Create("PatientInfor.json")
  if err != nil{
    panic(err)
  }

  defer func(){
    if err := fo.Close(); err !=nil {
      panic(err)
    }
  }()
  fo.Write(str)
}//SaveData

func GetInfo(){

  //need user ID of patient, client sends request, if patient exists they will be displayed and
  //saved

  var person Patient
  var ID string

  fmt.Println("Enter patient ID")
  fmt.Scanln(&ID)

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

  if resp.StatusCode == 200 {
    defer resp.Body.Close()

    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    responseString := string(responseData)
    fmt.Println(responseString)
    person.ID = ID
    person.Firstname = responseString
    UpdateLocalRecords(person)
  }else {
    fmt.Println("Patient not found")
  }//end else, if
}//end Patient

func GetEncounter(){
//need patient ID
//I dont understand this encounter one

}

func GetMedication(){
  //enter patient ID, then medication ID
  var patientID, conditionID string

  fmt.Println("Enter patient ID first")
  fmt.Scanln(&patientID)

  fmt.Println("Enter medication ID")
  fmt.Scanln(&conditionID)

  person := FindPatient(patientID)

  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Condition/" + conditionID)
  req,err := http.NewRequest("GET", url, nil)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}

  resp := client.Do(req)

  defer resp.Body.Close()

  if resp.StatusCode == 200 {


    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    responseString := string(responseData)
    person.Condition = responseString
    fmt.Println(responseString)
    UpdateLocalRecords(person)

  }else {
    fmt.Println("Patient not found")
  }


}

func GetCondition(){
  var patientID, medID string

  fmt.Println("Enter patient ID first")
  fmt.Scanln(&patientID)

  fmt.Println("Enter medication ID")
  fmt.Scanln(&medID)

  person := FindPatient(patientID)

  url := fmt.Sprintf("http://fhirtest.uhn.ca/baseDstu3/Medication/" + medID)
  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}

  resp := client.Do(req)

  defer resp.Body.Close()

  if resp.StatusCode == 200 {


    responseData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Fatal(err)
    }
    responseString := string(responseData)
    person.Medication = responseString
    fmt.Println(responseString)
    UpdateLocalRecords(person)

  }else {
    fmt.Println("Patient not found")
  }

}

func FindPatient(PatientID string) (bool, Patient){
  for _, Record := range Patients{
    if Record.ID == PatientID{
      return true, Record
    }
  }
}
