'use strict';

module.exports = {
    addRoutes: addRoutes
};

/**
 * Adición de las rutas a la API.
 */
function addRoutes(api) {
    api.post('/productos', createAccount);
    api.put('/productos/:id', updateAccount);
    api.get('/productos/:id', getAccount);
    api.get('/restaurantes', getAccounts);
    api.delete('/productos/:id', deleteAccount);
}

var infrastructure = require('../../infrastructure');
var log = infrastructure.logger;
var errors = infrastructure.errors;

var accountService = require('../../application').accountService;

/**
 * Crea un nuevo producto y lo inserta en la base de datos.
 * @param {Object} req - req.body contains accountData minus the id
 * @param {Object} res - res.body contains the inserted account (including the id)
 */
function createAccount(req, res) {

    var accountData = req.body;

    accountService.createAccount(accountData)
        .then(function(account) {
            res.send(account);
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

/**
 * Actualiza un producto existente.
 * @param {Object} req - req.body contains accountData including the id
 * @param {Object} res - res.body contains the updated account (including the id)
 */
function updateAccount(req, res) {

    var accountData = req.body;

    accountService.updateAccount(accountData)
        .then(function(account) {
            res.send(account);
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

/**
 * Obtiene un producto existente.
 * @param {Object} req - req.params.id contains id of the account to get
 * @param {Object} res - res.body contains the requested account
 */
function getAccount(req, res) {

    var id = req.params.id;

    accountService.getAccount(id)
        .then(function(account) {
            res.send(account);
        })
        .catch(errors.NotFoundError, function() {
            res.status(404).send({'message': 'Account ' + id + ' does not exist'});
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

/**
 * Obtiene todos los productos.
 * @param {Object} req - no used
 * @param {Object} res - res.body contains an array of all accounts
 */
function getAccounts(req, res) {
    accountService.getAccounts()
        .then(function(accounts) {
            res.send(accounts);
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

/**
 * Elimina un producto.
 * @param {Object} req - req.params.id contains id of the account to delete
 * @param {Object} res - res.body contains no content
 */
function deleteAccount(req, res) {

    var id = req.params.id;

    accountService.deleteAccount(id)
        .then(function() {
            res.status(204).send();  // No Content
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}
