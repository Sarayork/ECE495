package sara_folder
import(
  "encoding/gob"
  "fmt"
  "net"
  "bufio"
	"io"
	"log"
	"os"
	"strings"
)
type Client struct {
  conn  net.Conn
  nickname  string
  ch  chan string
  email string
  address string
  phone string
  role  string
}

func main(){
  //lkisten for port
  ln, err := net.Listen("tcp", ":9999")
  if err != nil{
    fmt.Println(err)
    os.Exit(1)
    return
  }//error handling

  //add channels for users
  msgchan := make(chan string)
  addchan := make(chan Client)
  rmchan :=make(chan Client)
  for{
    c, err := ln.Accept() //accept connectionlistener
    if err != nil{
      fmt.Println(err)
      continue
    }
    go handleConnention(conn, msgchan, addchan, rmchan)

    go handleConnection(c) //concurrency for connection
  }//end for
}//end server main

func handleConnection(c net.Conn, msgchan chan<-string, addchan chan<-Client, rmchan chan<-Client){
  //reader for a new connection
  bufc := bufio.NewReader(c)
  defer c.Close()

  client := Client{
    conn: c,
    nickname: promptName(c, bufc),
    ch: make(chan string),
    phone: promptPhone(c, bufc),
    email: promptEmail(c, bufc),

  }

  //error handling for empty Name
  if strings.TrimSpace(client.nickname) == ""{
    io.WriteString(c, "invalid username\n")
    return
  }

  //register user
  addchan <- client
  //defer for a disconnecting client
  defer func(){
    msgchan <- fmt.Sprintf("User %s left\n", client.nickname)
    log.Printf("%s connection closed\n", c.RemoteAddr())
    rmchan <- client
  }()
  //gretting to new client
  io.WriteString(c, fmt.Sprintf("Welcome, %s\n", client.nickname))
  msgchan <- fmt.Sprintf("%s: New user", client.nickname)
//save information into a file
  file, _ := os.OpenFile("clients.txt", os.O_WRONLY|os.O_APPEND, 0666)
  defer file.Close()
  io.WriteString(file, client.nickname+'\n')

  go client.ReadLinesInto(msgchan)
  client.WriteLinesFrom(client.ch)

}//end handleConnection

func (c Client) ReadLinesInto(ch chan<-string){
  bufc := bufio.NewReader(c.conn)
  for{
    line, err:= bufc.ReadString('\n')
    if err != nil{
      break
    }
    ch <-fmt.Sprintf("%s: %s", c.nickname, line)
  }
}//end ReadLinesInto
func (c Client) WriteLinesFrom(c.conn, msg){
  for msg := range ch{
    _,err := io.WriteString(c.conn, msg)
    if err != nil{
      return
    }
  }
}//end WriteLinesFrom
