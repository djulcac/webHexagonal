'use strict';

/**
 * Un producto.
 *   {int}     id
 *   {String}  nombre - nombre del producto
     {String}  Precio - precio del producto
     {String}  Tipo - tipo del producto
 *
 * Ejemplo:
 *   {
 *       id: [number],
 *       nombre: 'Arroz con pollo',
	 precio: 12.5,
	 tipo: 'Segundo'
 *   }
 */

var _ = require('lodash');

var Account = function(accountData) {
    if (accountData) {
        _.extend(this, accountData);
    }
};

module.exports = Account;
