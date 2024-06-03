package main 
import ("log"
  "github.com/arwysyah/go_tart/cmd/api")
// import "../go_start/cmd/api/api";
func main (){
  server:= api.NewAPIServer (":8080",nil)
  if err := server.Run(); err != nil{
    log.Fatal(err)
  }
}
