package main
import(
  "net/http"
  //"fmt"
  "log"
  //"encoding/json"
  "github.com/gorilla/mux"
  "./handlers/api/v1"
  //"./config"
  "github.com/gorilla/handlers"


)


func main()  {
  r := mux.NewRouter()
    r.Use(handlers.ProxyHeaders)

    r.HandleFunc("/platos", v1.GetPlatos).Methods("GET")
    r.HandleFunc("/platos/{id:[0-9]+}", v1.GetPlato).Methods("GET")
    r.HandleFunc("/platos", v1.CreatePlato).Methods("POST")
  	r.HandleFunc("/platos/{id:[0-9]+}", v1.UpdatePlatos).Methods("PUT")
  	r.HandleFunc("/platos/{id:[0-9]+}", v1.DeletePlatos).Methods("DELETE")

    r.HandleFunc("/restaurantes", v1.GetRestaurantes).Methods("GET")
    r.HandleFunc("/restaurantes/{id:[0-9]+}", v1.GetRestaurant).Methods("GET")
    r.HandleFunc("/restaurantes", v1.CreateRestaurant).Methods("POST")
  	r.HandleFunc("/restaurantes/{id:[0-9]+}", v1.UpdateRestaurantes).Methods("PUT")
  	r.HandleFunc("/restaurantes/{id:[0-9]+}", v1.DeleteRestaurantes).Methods("DELETE")

    r.HandleFunc("/restaurantes_platos", v1.GetRestaurantes_platos).Methods("GET")
    r.HandleFunc("/restaurantes_platos/{id:[0-9]+}", v1.GetRestaurant_platos).Methods("GET")
    r.HandleFunc("/restaurantes_platos", v1.CreateRestaurant_platos).Methods("POST")
  	r.HandleFunc("/restaurantes_platos/{id:[0-9]+}", v1.UpdateRestaurantes_platos).Methods("PUT")
  	r.HandleFunc("/restaurantes_platos/{id:[0-9]+}/{id2:[0-9]+}", v1.DeleteRestaurantes_platos).Methods("DELETE")

    //
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe("[::]:8000", nil))
  //mux := mux.NewRouter()
	//log.Println("El servidor esta a la escucha en el puerto ", 3000 )
	//server := &http.Server{
	//	Addr: config.UrlServer(),
	//	Handler: mux,
	//}
	//log.Fatal(server.ListenAndServe())
/*
  url:="https://c5u2jk30nd.execute-api.us-west-2.amazonaws.com/nuevo/restaurantes/1"
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
  log.Fatal("NewRequest: ", err)
  return
  }


  client := &http.Client{}


  resp, err := client.Do(req)
  if err != nil {
  log.Fatal("Do: ", err)
  return
  }

  defer resp.Body.Close()
  plato:=Restaurantes{}
  if err := json.NewDecoder(resp.Body).Decode(&plato); err != nil {
  log.Println(err)
  }
  fmt.Println(plato)

*/
}
