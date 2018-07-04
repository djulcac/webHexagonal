import json

from flask import render_template,request
from wtforms import Form, BooleanField, StringField, PasswordField, validators,IntegerField,FloatField

#import domain.api.api as _api
import domain.data.data as _data
def vista():
	f = MiForm(request.form)
	if request.method =='POST' and f.validate():
		dic = {}
		dic['id_plato'] = f.id_plato.data
		info = {}
		r = _data.read(dic['id_plato'])
		if r==None:
			info['message']='Error 1'
		js = json.loads(r)
		if js['status']==200:
			info['data'] = js['data']
			info['message'] = 'Correcto'
		else:
			info['message'] = 'Error 2'
		#print(info)
		return render_template('plato/get/id.html',info=info)
	return render_template('plato/get/id.html',form=f)

class MiForm(Form):
    id_plato = IntegerField('Id_plato',[validators.DataRequired()])

