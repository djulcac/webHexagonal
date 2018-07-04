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

func GetPlatos(w http.ResponseWriter, r *http.Request){
	out,_:=json.Marshal(models.GetPlatos())
	//string(out)
	fmt.Fprintf( w,string(out) )

}

func GetPlato(w http.ResponseWriter, r *http.Request){
	if user, err := getUserByRequest(r); err != nil{
		fmt.Println("error")
	}else{
		out,_:=json.Marshal(user)
		fmt.Fprintf( w,string(out) )
	}
}

func CreatePlato(w http.ResponseWriter, r *http.Request){
	plato := &models.Plato{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&plato); err != nil{
		fmt.Println("200")
		return
	}
	if err := plato.Insert(); err != nil{
		fmt.Println("200")
		return
	}

	out,_:=json.Marshal(plato)
	//string(out)
	fmt.Fprintf( w,string(out) )
}

func UpdatePlatos(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	userId, _ :=  vars["id"] 

	plato, err := getUserByRequest(r)
	if err != nil{
		fmt.Println("400")
		return
	}

	request := &models.Plato{}
	decoder := json.NewDecoder(r.Body)

	if err:= decoder.Decode(request); err != nil{
		fmt.Println("200")
		return
	}

	plato.ID = userId
	plato.Nombre = request.Nombre
	plato.Descripcion = request.Descripcion
	plato.Precio = request.Precio


	if err := plato.Update(); err != nil{
		fmt.Println("200")
		return
	}

	out,_:=json.Marshal(plato)
	//string(out)
	fmt.Fprintf( w,string(out) )

}

func DeletePlatos(w http.ResponseWriter, r *http.Request){
	if plato, err := getUserByRequest(r); err != nil{
		fmt.Println("500")
	}else{
		plato.Delete()

		fmt.Fprintf( w,"holi" )

	}
}

func getUserByRequest(r *http.Request) (*models.Plato, error){
	vars := mux.Vars(r)
	userId, _ :=  strconv.Atoi( vars["id"] )
	plato := models.GetPlatoById(userId)

	if plato.ID == "0"{
		return plato, errors.New("El usuario no existe en la base de datos.")
	}
	return plato, nil
}
