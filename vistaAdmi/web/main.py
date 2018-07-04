import time
import random
import json

from flask import Flask, Response, render_template,send_file,request,flash,redirect,url_for
from flask import session,g

from c import index,platoS
from c.platos import id,nuevo,delete

from c.plato import set
from c.plato.get import all,id
#from c.insumos.get import all
import c

app = Flask(__name__)
app.secret_key='adsfawefasdvcrgaebf'

# pagina no encontrada
@app.errorhandler(404)
def page_not_found(e):
	return redirect(url_for('index'))

@app.route("/")
@app.route("/index")
def index():
	#return 'ok'
	return c.index.vista()

@app.route("/insumo/get/all",methods=['GET','POST'])
def insumo_get_all():
	return c.insumo.get.all.vista()

#
@app.route("/platos",methods=['GET','POST'])
def platos():
	return c.platoS.vista()

@app.route("/platos/<int:id>",methods=['GET','POST'])
def platos_id(id):
	return c.platos.id.vista()
	return "s"+str(id)

@app.route("/platos/nuevo",methods=['GET','POST'])
def platos_nuevo():
	return c.platos.nuevo.vista()

@app.route("/platos/delete",methods=['GET','POST'])
def platos_delete():
	return c.platos.delete.vista()


# anterior
@app.route("/plato/set",methods=['GET','POST'])
def plato_set():
	return c.plato.set.vista()
@app.route("/plato/get/all",methods=['GET','POST'])
def plato_get_all():
	return c.plato.get.all.vista()
@app.route("/plato/get/id",methods=['GET','POST'])
def plato_get_id():
	return c.plato.get.id.vista()



if __name__ == "__main__":
	app.run(debug=True,port=5000)
#app.run(debug=False,port=80,host='0.0.0.0')


