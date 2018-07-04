import json

from flask import render_template,request
from wtforms import Form, BooleanField, StringField, PasswordField, validators,IntegerField,FloatField

#import domain.api.api as _api
import domain.data.data as _data
def vista():
	f = MiForm(request.form)
	print("1"*5)
	if request.method =='POST' and f.validate():
		print("2"*5)
		dic = {}
		dic['titulo'] = f.titulo.data
		dic['contenido'] = f.contenido.data
		print(dic)
		r = _data.create(dic)
		print(r)
		info={}
		info['message'] = 'Correcto'
		return render_template('platos/nuevo.html',info=info)
	print("3"*5)
	return render_template('platos/nuevo.html',form=f)

class MiForm(Form):
	titulo = StringField('Título', [validators.DataRequired()])
	contenido = StringField('Contenido', [validators.DataRequired()])
