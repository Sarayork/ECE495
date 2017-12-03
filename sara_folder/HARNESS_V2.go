package main

import(
  "encoding/json"
  "fmt"
  //"io/ioutil"
  "log"
  "math/rand"
  "net"
  "net/http"
  //"os"
  "time"
  "strconv"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"

)
//////hub that stores more than one client in a map
type Hub struct{
  clients map[*Client]bool
  broadcast chan []byte
  addClient chan *Client
  removeClient chan *Client
}

//client type for the hub for messaging
type Client struct{
  ws *websocket.Conn
  send chan []byte
}

//new Hub, for the messaging part of the program
var hub = Hub{
  broadcast: make(chan []byte),
  addClient: make(chan *Client),
  removeClient: make(chan *Client),
  clients: make(map[*Client]bool),
}
/////user struct to store information
type User struct {
  uid string `json: "uid, omitempty"`
  currentIP string `json:"currentIP, omitempty"`
  timestamp string `json: "timestamp"`
  name string `json:"name, omitempty"`
  password string `json:"password, omitempty"`
  email string `json:"email"`
  address string `json:"addr"`
  phoneNumber string `json:"phoneNumber"`
  roles string `json:"roles"`

}
/////notification struct to store notification information
type Notification struct{
  src string `json:"src"`
  destUid string `json:"destUid"`
  role string `json:"role"`
  dateTime string `json:"dateTime"`

}

//array of people of type User
var People []User
//
var upgrader = websocket.Upgrader{}

/////hub function to set up mapping the clients
func (hub *Hub) start(){
  for{
    select{
    case conn := <-hub.addClient:
      //add new clients
      hub.clients[conn] = true
    case conn := <-hub.removeClient:
      //remove client
      if _, ok := hub.clients[conn]; ok{
        delete(hub.clients, conn)
        close(conn.send)
      }//end if
    case message := <-hub.broadcast:
      //braodcast message to all
      for conn := range hub.clients{
        select{
        case conn.send <- message:
        default:
          close(conn.send)
          delete(hub.clients, conn)
        }//end inside select
      }//end for loop

    }//end select
  }//end for loop
}/////////////////end hub

/**
 * @api {post} /user Subscribe a new user
 * @apiName PostUser
 * @apiGroup User
 * @apiVersion 0.1.0
 *
 * @apiDescription Generate unique Users-ID and add new user account in the HArNESS data base
 *
 * @apiParam {string} name User's name
 * @apiParam {string} password User's password
 * @apiParam {string} email User's email
 * @apiParam {string} address User's address
 * @apiParam {string} phoneNumber User's phoneNumber
 * @apiParam {string} roles User's roles
 * @apiParam {string} timestamp User's timestamp
 * @apiParam {string} currentIp User's currentIp
 *
 * @apiSuccess {string} uid Users-ID
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 201 (Created)
 *     {
 *       "uid": "12461"
 *     }
 *
 * @apiUse UserNameAlreadyExist
 */
func subscribe(w http.ResponseWriter, req *http.Request) {

  fmt.Println("method: ", req.Method)//get the request Method
  //add in html parseing for the user's subscription information
  //add json data from FHIR server
  var usr User
  decoder :=json.NewDecoder(req.Body)
  err := decoder.Decode(&usr)
  if err != nil{
    panic(err)
  }

//random user ID generator
  s1 := rand.NewSource(time.Now().UnixNano())
  r1 := rand.New(s1)
  key := r1.Intn(1000)//create random ID for user
  //convert the generated key to be saved in the User struct
  usr.uid = strconv.Itoa(key)
  //save Users IP
  usr.currentIP = GetIP()
  //append usr to the People array
  People = append(People, usr)

  json.NewEncoder(w).Encode(usr)

}

/**
 * @api {put} /beatHeart/:uid Beat Heart
 * @apiName BeatHeart
 * @apiGroup User
 * @apiVersion 0.1.0
 *
 * @apiDescription Update the current IP address of the user
 *
 * @apiParam {string} uid User's id
 * @apiParam {string} currentIp User's current IP address
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 204 No content
 *
 * @apiUse UidNotFoundError
 */
func beatHeart(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
  _ = json.NewDecoder(req.Body).Decode(&People)


  for _, item := range People{
  //loop trough users to find the user ID the client is searching for
    if item.uid == params["uid"]{
      //if found we send the user a "heartbeat"
      //from the ohter user they were searching for
      item.uid = GetIP()
      json.NewEncoder(w).Encode(item)
    } else {
      fmt.Println("User Not Found")
    }
  }


}

func GetMedication(w http.ResponseWriter, req *http.Request){

  params := mux.Vars(req)
  var patient User
  uid := params["uid"]

  url := fmt.Sprintf("http.//fhirtest.uhn.ca/baseDstu3/Patient/" + uid)

  req, err := http.NewRequest("GET", url, nil)
  req.Header.Set("Content-Type", "application/json")

  client := &http.Client{}
  if err != nil{
    log.Fatal("Do: ", err)
    return
  }
  resp, err := client.Do(req)
  if err != nil{
    log.Fatal("DO:", err)
    return
  }
  defer resp.Body.Close()
  json.NewEncoder(w).Encode(patient)
}
/**
 * @api {get} /user/:uid Look up ip
 * @apiName GetUserIpByUid
 * @apiGroup User
 * @apiVersion 0.1.0
 *
 * @apiDescription Get <code>currentIp</code> of the user with given <code>uid</code>. If not found,
 * return a UidNotFoundError
 *
 * @apiParam {string} uid User's id
 *
 * @apiSuccess {string} currentIp User's current IP address
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "currentIp": "127.15.1.145"
 *     }
 *
 * @apiUse UidNotFoundError
 */
func lookUpIp(w http.ResponseWriter, req *http.Request) {

  params := mux.Vars(req)
  _ = json.NewDecoder(req.Body).Decode(&People)

  for _, item := range People{
    if item.uid == params["uid"]{
      item.uid = GetIP()

      json.NewEncoder(w).Encode(item)
    }
  }
}
/**
 * @api {post} /addInvitation Add invitation
 * @apiName PostInvitation
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 *
 * @apiDescription Add new invitation. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid destination user id
 * @apiParam {string} srcRole Sender role
 * @apiParam {string} destRole destination role
 *
 * @apiSuccess {string} src Sender user id
 * @apiSuccess {string} dest Target user id
 * @apiSuccess {string} status Invitation status
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 201 Created
 *
 * @apiUse UidNotFoundError
 */
func addInvitation(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {get} /invitations/:uid Look up invitations
 * @apiName GetReceivedInvitations
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 *
 * @apiDescription Get all pending invitation. If the <code>uid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid Destination user's id
 *
 * @apiSuccess {Relation} invitations List of the received invitation
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "invitations": [
 *        	{
 *       	 	"src": "556171",
 *       	 	"role": {
 *       	 		"srcRole": "doctor",
 *       	 		"destRole": "patient"
 *       	 	}
 *          },
 *        	{
 *       	 	"src": "17117",
 *       	 	"role": {
 *       	 		"srcRole": "pharmacist",
 *       	 		"destRole": "patient"
 *       	 	}
 *          }
 *       }
 *     }
 *
 * @apiUse UidNotFoundError
 */
func lookUpInvitations(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {delete} /invitation/:srcUid/:destUid/:srcRole/:destRole Remove invitation
 * @apiName DeleteInvitation
 * @apiGroup Invitation
 * @apiVersion 0.1.0
 *
 * @apiDescription Remove relation. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid Destination user id
 * @apiParam {string} srcRole Sender user role
 * @apiParam {string} destRole Destination user role
 *
 * @apiUse UidNotFoundError
 */
func removeInvitation(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {post} /notification Add notification
 * @apiName PostNotification
 * @apiGroup Notification
 * @apiVersion 0.1.0
 *
 * @apiDescription Add new message notification. If the <code>srcUid</code> or <code>destUid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} srcUid Sender user id
 * @apiParam {string} destUid Target user id
 *
 * @apiSuccess {string} src Sender user id
 * @apiSuccess {string} dest Target user id
 * @apiSuccess {string} dateTime Time of sending notification
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 Created
 *
 * @apiUse UidNotFoundError
 */
func addMsgNotif(w http.ResponseWriter, req *http.Request) {
}

/**
 * @api {get} /notifications/:uid Look up notifications
 * @apiName GetReceivedNotifications
 * @apiGroup Notification
 * @apiVersion 0.1.0
 *
 * @apiDescription Get all received message notification to user. If the <code>uid</code>
 * do not exist, return UidNotFoundError
 *
 * @apiParam {string} uid Target user's id
 *
 * @apiSuccess {[]Notif} notifications List of the received notifications
 *
 * @apiSuccessExample Success-Response:
 *     HTTP/1.1 200 OK
 *     {
 *       "notifications": [
 *        	{
 *       	 	"src": "556171",
 *         	 	"dateTime": "2017-11-19 17:15:45"
 *          },
 *        	{
 *       	 	"src": "17117",
 *         	 	"dateTime": "2017-11-20 8:34:22"
 *          },
 *        	{
 *       	 	"src": "364346",
 *         	 	"dateTime": "2017-11-20 11:55:53"
 *          }
 *       ]
 *     }
 *
 * @apiUse UidNotFoundError
 */
func lookUpMsgNotif(w http.ResponseWriter, req *http.Request) {
}
/*
* GET
* used to get the current ip of the user that subscribed
* On Success : returns currentIP of user
*/
func GetIP() string{
  addrs, err := net.InterfaceAddrs()
  if err != nil{
    panic(err)
  }

  return addrs[1].String()
}
func (c *Client) write(){
  defer func(){
    c.ws.Close()
  }()//end defer

  for{
    select{
    case message, ok := <-c.send:
      if !ok {
        c.ws.WriteMessage(websocket.CloseMessage, []byte{})
        return
      }//end if
      c.ws.WriteMessage(websocket.TextMessage, message)
    }//end select
  }//end for

}////////////end write

func (c *Client) read(){

  defer func(){
    hub.removeClient <- c
    c.ws.Close()
  }()//end defer

  for {
    _, message, err := c.ws.ReadMessage()
    if err != nil {
      hub.removeClient <- c
      c.ws.Close()
      break
    }//end if
    hub.broadcast <- message
  }//end for

}//////////////////////////end read
func Messager(res http.ResponseWriter, req *http.Request){

  http.ServeFile(res, req, "index.html")
  conn, err := upgrader.Upgrade(res, req, nil)
  if err != nil {

    return
  }

  client := &Client {
    ws: conn,
    send: make(chan []byte),
  }

  hub.addClient <-client

  go client.write()
  go client.read()

}
/*Main funciton for the handlers
*/
func main(){

  go hub.start()

  http.HandleFunc("/user", subscribe)
  http.HandleFunc("/message", Messager)
  http.HandleFunc("/beatHeart/uid", beatHeart)
  http.HandleFunc("/lookUpIp/uid", lookUpIp)
  http.HandleFunc("/medication/uid", GetMedication)

  log.Fatal(http.ListenAndServe(":54321", nil))
}
