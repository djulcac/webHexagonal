package v1
import(
"../../../models"
"net/http"
"encoding/json"
"fmt"
"errors"
	"strconv"
	"github.com/gorilla/mux"

)

func GetRestaurantes(w http.ResponseWriter, r *http.Request){
	out,_:=json.Marshal(models.GetRestaurantes())
	//string(out)
	fmt.Fprintf( w,string(out) )

}

func GetRestaurant(w http.ResponseWriter, r *http.Request){
	if rest, err := getRestByRequest(r); err != nil{
		fmt.Println("error")
	}else{
		out,_:=json.Marshal(rest)
		fmt.Fprintf( w,string(out) )
	}
}

func CreateRestaurant(w http.ResponseWriter, r *http.Request){
	restaurant := &models.Restaurant{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&restaurant); err != nil{
		fmt.Println("200")
		return
	}
	if err := restaurant.Insert(); err != nil{
		fmt.Println("200")
		return
	}

	out,_:=json.Marshal(restaurant)
	//string(out)
	fmt.Fprintf( w,string(out) )
}

func UpdateRestaurantes(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	restId, _ :=  vars["id"]

	restaurant, err := getRestByRequest(r)
	if err != nil{
		fmt.Println("400")
		return
	}

	request := &models.Restaurant{}
	decoder := json.NewDecoder(r.Body)

	if err:= decoder.Decode(request); err != nil{
		fmt.Println("200")
		return
	}

	restaurant.ID = restId
	restaurant.Latitud = request.Latitud
	restaurant.Longitud = request.Longitud
	restaurant.Nombre = request.Nombre


	if err := restaurant.Update(); err != nil{
		fmt.Println("200")
		return
	}

	out,_:=json.Marshal(restaurant)
	//string(out)
	fmt.Fprintf( w,string(out) )

}

func DeleteRestaurantes(w http.ResponseWriter, r *http.Request){
	if restaurant, err := getRestByRequest(r); err != nil{
		fmt.Println("500")
	}else{
		restaurant.Delete()

		fmt.Fprintf( w,"holi" )

	}
}

func getRestByRequest(r *http.Request) (*models.Restaurant, error){
	vars := mux.Vars(r)
	restId, _ :=  strconv.Atoi( vars["id"] )
	restaurant := models.GetRestauranteById(restId)

	if restaurant.ID == "0"{
		return restaurant, errors.New("El usuario no existe en la base de datos.")
	}
	return restaurant, nil
}
