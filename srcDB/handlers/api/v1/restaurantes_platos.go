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

func GetRestaurantes_platos(w http.ResponseWriter, r *http.Request){
	out,_:=json.Marshal(models.GetRestaurantes_platos())
	//string(out)
	fmt.Fprintf( w,string(out) )

}

func GetRestaurant_platos(w http.ResponseWriter, r *http.Request){
	if rest, err := getRest_platosByRequest(r); err != nil{
		fmt.Println("error")
	}else{
		out,_:=json.Marshal(rest)
		fmt.Fprintf( w,string(out) )
	}
}

func CreateRestaurant_platos(w http.ResponseWriter, r *http.Request){
	restaurant := &models.Restaurant_platos{}
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

func UpdateRestaurantes_platos(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	restId, _ :=  vars["id"]

	restaurant, err := getRest_platosByRequest(r)
	if err != nil{
		fmt.Println("400")
		return
	}

	request := &models.Restaurant_platos{}
	decoder := json.NewDecoder(r.Body)

	if err:= decoder.Decode(request); err != nil{
		fmt.Println("200")
		return
	}

	restaurant.ID = restId
	restaurant.IDP = request.IDP



	if err := restaurant.Update(); err != nil{
		fmt.Println("200")
		return
	}

	out,_:=json.Marshal(restaurant)
	//string(out)
	fmt.Fprintf( w,string(out) )

}

func DeleteRestaurantes_platos(w http.ResponseWriter, r *http.Request){
	if restaurant, err := getRest_platosByRequest2(r); err != nil{
		fmt.Println("500")
	}else{
		restaurant.Delete()

		fmt.Fprintf( w,"holi" )

	}
}

func getRest_platosByRequest(r *http.Request) (*models.Restaurant_platos, error){
	vars := mux.Vars(r)
	restId, _ :=  strconv.Atoi( vars["id"] )
	restaurant := models.GetRestaurante_platosById(restId)

	if restaurant.ID == "0"{
		return restaurant, errors.New("El usuario no existe en la base de datos.")
	}
	return restaurant, nil
}

func getRest_platosByRequest2(r *http.Request) (*models.Restaurant_platos, error){
	vars := mux.Vars(r)
	restId, _ :=  strconv.Atoi( vars["id"] )
  restId2, _ :=  strconv.Atoi( vars["id2"] )

  restaurant := models.GetRestaurante_platosById2(restId,restId2)

	if restaurant.ID == "0"{
		return restaurant, errors.New("El usuario no existe en la base de datos.")
	}
	return restaurant, nil
}
