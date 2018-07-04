import json

from flask import render_template,request
from wtforms import Form, BooleanField, StringField, PasswordField, validators,IntegerField,FloatField

#import domain.api.api as _api
import domain.data.data as _data
def vista():
	info = {}
	#r = _api.all()
	r = _data.all()
	if r==None:
		info['message']='Error 1'
	js = json.loads(r)
	if js['status']==200:
		info['data'] = js['data']
		info['message'] = 'Correcto'
	else:
		info['message'] = 'Error 2'
	print(info)
	return render_template('platos.html',info=info)

