import json

from flask import render_template,request
from wtforms import Form, BooleanField, StringField, PasswordField, validators,IntegerField,FloatField

#import domain.api.api as _api
import domain.data.data as _data
def vista():
	f = MiForm(request.form)
	if request.method=='POST' and f.validate():
		dic = {}
		dic['nombre'] = f.nombre.data
		dic['descripcion'] = f.descripcion.data
		dic['precion'] = f.precio.data
		r = _data.create(json.dumps(dic))
		#print(r)
		info = {}
		if r == None:
			info['status'] = 'Error 1'
		js = json.loads(r)
		if js['status'] == 200:
			info['status'] = 'Correcto'
		else:
			info['status'] = 'Error 2'
		return render_template('plato/set.html',info=info)		
	return render_template('plato/set.html',form=f)

class MiForm(Form):
    nombre = StringField('Nombre', [validators.Length(min=3, max=20)])
    descripcion = StringField('Descripción', [validators.Length(min=3, max=20)])
    precio = FloatField('Precio',[validators.DataRequired()])
    #foto = StringField('Nombre del plato', [validators.Length(min=3, max=20)])

