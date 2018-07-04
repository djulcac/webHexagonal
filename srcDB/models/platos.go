package models
import(
  "strconv"
)
type Plato struct{
  ID          string  `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Precio      float64 `json:"precio"`
}
const platoSchema string = `CREATE TABLE platos(
        id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
        nombre VARCHAR(30) NOT NULL,
        descripcion VARCHAR(60) NOT NULL,
        precio VARCHAR(40) NOT NULL)`
type Platos []Plato
func GetPlatos()  Platos{
  sql := "SELECT * FROM platos"
    platos := Platos{}
    rows, _ := Query(sql)

    for rows.Next(){
      plato := Plato{}
      rows.Scan(&plato.ID, &plato.Nombre, &plato.Descripcion, &plato.Precio)
      platos = append(platos, plato)
    }

    return platos
}
func GetPlatoById(id int) *Plato{
  sql := "SELECT * FROM platos WHERE id=?"
  return GetPlato(sql, id)
}
func GetPlato(sql string, conditional interface{}) *Plato{
  user := &Plato{}
  rows, err := Query(sql, conditional)
  if err != nil{
    return user
  }

  for rows.Next(){
    rows.Scan(&user.ID, &user.Nombre, &user.Descripcion, &user.Precio)
  }
  return user
}
func (this *Plato) Save() error {
  if this.ID == "0"{
    return this.Insert()
  }else{
    return this.Update()
  }
}
func (this *Plato) Insert() error {
  sql := "INSERT platos SET nombre=?, descripcion=?, precio=?"
  id, err := InsertData(sql, this.Nombre, this.Descripcion, this.Precio)
  this.ID = strconv.FormatInt(id, 10)
  return err
}

func (this *Plato) Update() error {
  sql := "UPDATE platos SET nombre=?, descripcion=?, precio=? WHERE id=?"
  _, err := Exec(sql, this.Nombre, this.Descripcion, this.Precio , this.ID)
  return err
}

func (this *Plato) Delete() error {
  sql := "DELETE FROM platos WHERE id=?"
  _, err := Exec(sql, this.ID)
  return err
}
