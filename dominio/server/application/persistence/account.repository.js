'use strict';

module.exports = {
    createAccount: createAccount,
    updateAccount: updateAccount,
    getAccount: getAccount,
    getAccounts: getAccounts,
    deleteAccount: deleteAccount
};

var knex = require('./db').knex;
var joinjs = require('join-js').default;
var resultMaps = require('./resultmaps');
var Account = require('../../domain').Account;

/**
 * Crea un nuevo producto y lo inserta en la base de datos.
 * @param {Object} accountData minus the id
 * @return {Promise} A promise that returns the inserted account (including the id)
 */
function createAccount(accountData) {

    var account = new Account(accountData);

    return knex('productos')
        .returning('id')
        .insert(account)
        .then(function(ids) {
            account.id = ids[0];
            return account;
        });
}

/**
 * Actualiza un producto existente.
 * @param {Object} accountData including the id
 * @return {Promise} A promise that returns the updated account (including the id)
 */
function updateAccount(accountData) {

    var account = new Account(accountData);

    return knex('productos')
        .where('id', account.id)
        .update(account)
        .then(function() {
            return account;
        });
}

/**
 * Obtiene un producto existente.
 * @param {integer} id
 * @return {Promise} A promise that returns the desired account.
 */
function getAccount(id) {
    return knex
        .select('id', 'nombre', 'precio', 'tipo')
        .from('productos')
        .where('id', id)

        .then(function(resultSet) {
            return joinjs.mapOne(resultSet, resultMaps, 'accountMap');
        });
}

/**
 * Obtiene todos los productos.
 * @return {Promise} A promise that returns an array of all accounts.
 */
function getAccounts() {
    return knex
        .select('id', 'nombre', 'precio', 'tipo')
        .from('productos')

        .then(function(resultSet) {
            return joinjs.map(resultSet, resultMaps, 'accountMap');
        });
}

/**
 * Elimina un producto.
 * @param {integer} id
 * @return {Promise} A promise that gets fulfilled when the account is deleted.
 */
function deleteAccount(id) {
    return knex('productos')
        .where('id', id)
        .delete();
}
