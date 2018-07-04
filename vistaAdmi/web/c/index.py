import json

from flask import render_template,request
from wtforms import Form, BooleanField, StringField, PasswordField, validators,IntegerField,FloatField

#import domain.api.api as _api
import domain.data.data as _data
def vista():	
	return render_template('index.html',)

class MiForm(Form):
    nombre = StringField('Nombre', [validators.Length(min=3, max=20)])
    descripcion = StringField('Descripción', [validators.Length(min=3, max=20)])
    precio = FloatField('Precio',[validators.DataRequired()])
    #foto = StringField('Nombre del plato', [validators.Length(min=3, max=20)])

