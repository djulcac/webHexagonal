import requests
url = '10.4.20.81'

def create(strjson):
	'''
	Input
		id_plato
		nombre
		descripcion
		precio
		foto
	Output:
		status :
			200 : correcto
		message :
			1 : ok
		data :
	'''
	return '{"status": 200,"valor":1}'

def read(strjson):
	'''
	Input
		id_plato(int)
	Output:
		estado :
			200 : correcto
		valor :
			1 : ok
		data :
			id_plato
			nombre
			precio
			foto
			categoria
	'''
	return "{'status':200,'valor':1,'data':{'id_plato':1,'nombre':'tequeños','precio':12.3,'foto':'asdfasd','categoria':3 }}"
def all():
	'''
	Input
	Output:
		estado :
			200 : correcto
		valor :
			1 : ok
		data :
			[]
				id_plato
				nombre
				descripcion
				precio
				foto
	'''
	return '''
	{
		"status":200,
		"message":1,
		"data":[
			{"id_plato":1,"nombre":"tequeños","precio":12.3,"foto":"asdfasd","descripcion":"asdf" },
			{"id_plato":2,"nombre":"sopa","precio":12.3,"foto":"asdfasd","descripcion":"asdf" }
		]
	}
	'''
def update():
	pass
def delete():
	pass

