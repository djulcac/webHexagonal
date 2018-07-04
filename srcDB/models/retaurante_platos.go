package models
import(
  "strconv"
)
type Restaurant_platos struct{
  ID          string  `json:"idr"`
	IDP     string `json:"idp"`

  }
const restau_platosSchema string=`CREATE TABLE restaurantes_platos(
        id_res INT(6) UNSIGNED,
        idp INT(6) NOT NULL,
        UNIQUE(id_res,idp)
        )`
type Restaurantes_platos []Restaurant_platos

func GetRestaurantes_platos()  Restaurantes_platos{
  sql := "SELECT * FROM restaurantes_platos"
    restaurantes_platos := Restaurantes_platos{}
    rows, _ := Query(sql)

    for rows.Next(){
      restaurant := Restaurant_platos{}
      rows.Scan(&restaurant.ID, &restaurant.IDP)
      restaurantes_platos = append(restaurantes_platos, restaurant)
    }

    return restaurantes_platos
}
func GetRestaurante_platosById(id int) *Restaurant_platos{
  sql := "SELECT * FROM restaurantes_platos WHERE id_res=?"
  return GetRestaurant_platos(sql, id)
}
func GetRestaurante_platosById2(id int, id2 int) *Restaurant_platos{
  sql := "SELECT * FROM restaurantes_platos WHERE id_res=? AND idp=?"


  user := &Restaurant_platos{}
  rows, err := Query(sql, id, id2)
  if err != nil{
    return user
  }

  for rows.Next(){
    rows.Scan(&user.ID, &user.IDP)
  }
  return user
}
func GetRestaurant_platos(sql string, conditional interface{}) *Restaurant_platos{
  user := &Restaurant_platos{}
  rows, err := Query(sql, conditional)
  if err != nil{
    return user
  }

  for rows.Next(){
    rows.Scan(&user.ID, &user.IDP)
  }
  return user
}
func (this *Restaurant_platos) Save() error {
  if this.ID == "0"{
    return this.Insert()
  }else{
    return this.Update()
  }
}
func (this *Restaurant_platos) Insert() error {
  sql := "INSERT restaurantes_platos SET id_res=?, idp=?"
  id, err := InsertData(sql, this.ID, this.IDP)
  this.ID = strconv.FormatInt(id, 10)
  return err
}

func (this *Restaurant_platos) Update() error {
  sql := "UPDATE restaurantes_platos SET idp WHERE id_res=?"
  _, err := Exec(sql, this.IDP, this.ID)
  return err
}

func (this *Restaurant_platos) Delete() error {
  sql := "DELETE FROM restaurantes_platos WHERE id_res=? AND idp=?"
  _, err := Exec(sql, this.ID, this.IDP)
  return err
}
