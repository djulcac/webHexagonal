import json

from flask import render_template,request
from wtforms import Form, BooleanField, StringField, PasswordField, validators,IntegerField,FloatField

#import domain.api.api as _api
import domain.data.data as _data
def vista():
	info = {}
	key = request.args.get('id', default = '0', type = str)
	if key=='0':
		info['message']= 'Error'
	else:
		_data.delete(key)
		info['message']= 'Correcto'
	return render_template('platos/delete.html',info=info)
class MiForm(Form):
	titulo = StringField('Título', [validators.DataRequired()])
	contenido = StringField('Contenido', [validators.DataRequired()])
