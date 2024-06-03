 package api
import ("database/sql"
"log"
  "net/http"
 "github.com/arwysyah/go_tart/service/user" 
  "github.com/gorilla/mux");

type APIServer struct {
  addr string
  db *sql.DB
}


//Create API Server;
func NewAPIServer (addr string , db *sql.DB) *APIServer {
  return &APIServer {
    addr : addr,
    db:db,
  }
}
func (s *APIServer ) Run() error {
 router := mux.NewRouter();
  subrouter := router.PathPrefix("/api/v1").Subrouter()
  userService := user.NewHandler();
  userService.RegisterRoutes(subrouter);
  log.Println("Listening to the port $s");
  return http.ListenAndServe(s.addr, router)
}
