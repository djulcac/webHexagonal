package models
import(
  "strconv"
)
type Restaurant struct{
  ID          string  `json:"id"`
	Latitud     string `json:"latitud"`
	Longitud    string `json:"longitud"`
	Nombre      string `json:"nombre"`
}
const restauSchema string=`CREATE TABLE restaurantes(
        id_res INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        latitud VARCHAR(30) NOT NULL,
        longitud VARCHAR(60) NOT NULL,
        nombre VARCHAR(40) NOT NULL)`
type Restaurantes []Restaurant

func GetRestaurantes()  Restaurantes{
  sql := "SELECT * FROM restaurantes"
    restaurantes := Restaurantes{}
    rows, _ := Query(sql)

    for rows.Next(){
      restaurant := Restaurant{}
      rows.Scan(&restaurant.ID, &restaurant.Latitud, &restaurant.Longitud, &restaurant.Nombre)
      restaurantes = append(restaurantes, restaurant)
    }

    return restaurantes
}
func GetRestauranteById(id int) *Restaurant{
  sql := "SELECT * FROM restaurantes WHERE id_res=?"
  return GetRestaurant(sql, id)
}
func GetRestaurant(sql string, conditional interface{}) *Restaurant{
  user := &Restaurant{}
  rows, err := Query(sql, conditional)
  if err != nil{
    return user
  }

  for rows.Next(){
    rows.Scan(&user.ID, &user.Latitud, &user.Longitud, &user.Nombre)
  }
  return user
}
func (this *Restaurant) Save() error {
  if this.ID == "0"{
    return this.Insert()
  }else{
    return this.Update()
  }
}
func (this *Restaurant) Insert() error {
  sql := "INSERT restaurantes SET latitud=?, longitud=?, nombre=?"
  id, err := InsertData(sql, this.Latitud, this.Longitud, this.Nombre)
  this.ID = strconv.FormatInt(id, 10)
  return err
}

func (this *Restaurant) Update() error {
  sql := "UPDATE restaurantes SET latitud=?, longitud=?, nombre=? WHERE id_res=?"
  _, err := Exec(sql, this.Latitud, this.Longitud, this.Nombre , this.ID)
  return err
}

func (this *Restaurant) Delete() error {
  sql := "DELETE FROM restaurantes WHERE id_res=?"
  _, err := Exec(sql, this.ID)
  return err
}
