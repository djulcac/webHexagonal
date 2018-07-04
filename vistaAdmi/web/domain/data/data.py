import requests
import json
import os
url = 'http://localhost:8080/'
url = 'http://10.20.0.233:3000/'

def create(dic):
	print("hhhhh")
	titulo = dic['titulo']
	contenido = dic['contenido']
	#titulo = 'hhh'
	c = '''curl -d '{"id":2,"title":"'''+titulo+'''","content":"'''+contenido+'''","cocinero":{"id":1,"name":"Iman Tumorang","created_at":"2017-05-18T13:50:19Z","updated_at":"2017-05-18T13:50:19Z"},"updated_at":"2018-05-18T13:50:19Z","created_at":"2018-05-18T13:50:19Z"}' -H "Content-Type: application/json" -X POST '''+url+'''plato'''
	a=os.system(c)
	print('ggggggg')
	print(a)
	return '{"status": 200,"valor":1}'

def create1(strjson):
	#r = requests.post('http://httpbin.org/post', data = {'key':'value'})
	print("hhhhh")
	c = '''curl -d '{"id":2,"title":"Arroz con cffhele","content":"chele con arroz","cocinero":{"id":1,"name":"Iman Tumorang","created_at":"2017-05-18T13:50:19Z","updated_at":"2017-05-18T13:50:19Z"},"updated_at":"2018-05-18T13:50:19Z","created_at":"2018-05-18T13:50:19Z"}' -H "Content-Type: application/json" -X POST '''+url+''':3000/plato'''
	os.system(c)
	dic = {}
	dic['title'] = "otro"
	dic['content'] = "cc"
	dic['cocinero'] = {}
	dic['cocinero']['id'] = 1
	dic['update_at'] = '2018-05-18T13:50:19Z'
	dic['create_at'] = '2018-05-18T13:50:19Z'
	a = '''{"id":2,"title":"Addrroz con chele","content":"chele con arroz","cocinero":{"id":1,"name":"Iman Tumorang","created_at":"2017-05-18T13:50:19Z","updated_at":"2017-05-18T13:50:19Z"},"updated_at":"2018-05-18T13:50:19Z","created_at":"2018-05-18T13:50:19Z"}'''
	
	#r = requests.post(url+'plato', data = dic)
	r = requests.post(url+'plato', data = json.loads(json.dumps(a)))
	print("####")
	print(r)
	print("####gg")
	print(r.text)
	print("kkkk")
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
	return '''
	{
		"status":200,
		"message":1,
		"data":
			{"id_plato":1,"nombre":"tequeños","precio":12.3,"foto":"asdfasd","descripcion":"asdf" }
	}
	'''
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
				---
				id
				nombre
				precio
				tipo
	'''
	#r = requests.get(url+'productos')
	print("gg")
	r = requests.get(url+'plato')
	print("gg2")
	print(r.text)
	js = json.loads(r.text)
	#js['status'] = 200
	#js['messaje'] = 1
	d = {}
	d['data'] = js
	d['status'] = 200
	d['messaje'] = 1
	print(d)
	return json.dumps(d)
	print(js)
	return '''
	{
		"status":200,
		"message":1,
		"data":[
			{"id_plato":1,"nombre":"tequeños","precio":12.3,"foto":"asdfasd","descripcion":"asdf" },
			{"id_plato":2,"nombre":"sopa","precio":12.3,"foto":"asdfasd","descripcion":"asdf" },
			{"id_plato":3,"nombre":"sopa","precio":12.3,"foto":"asdfasd","descripcion":"asdf" }
		]
	}
	'''
def update():
	pass
def delete(id):
	c = '''curl -H "Content-Type: application/json" -X DELETE '''+url+'''plato/'''+str(id)
	a=os.system(c)
	print('ggggggg')
	print(a)
	return '{"status": 200,"valor":1}'

'''

curl -d '{"id":3,"title":"Arroz con pato","content":"Arroz,pato,etc","cocinero":{"id":3,"name":"Don Pedrito","created_at":"2018-05-18T13:50:19Z","updated_at":"2018-05-18T13:50:19Z"},"updated_at":"2018-05-18T13:50:19Z","created_at":"2018-06-18T13:50:19Z"}' -H "Content-Type: application/json" -X POST 10.20.0.233:3000/plato 




'''
