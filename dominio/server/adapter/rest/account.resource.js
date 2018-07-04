'use strict';

module.exports = {
    addRoutes: addRoutes
};

var urldb = "https://c5u2jk30nd.execute-api.us-west-2.amazonaws.com/nuevo/";
/**
 * Adición de las rutas a la API.
 */
function addRoutes(api) {
    api.post('/productos', createAccount);
    api.put('/productos/:id', updateAccount);
    api.get('/productos/:id', getAccount);
    api.get('/productos', getAccounts);
    api.get('/r', getAccounts);
    api.delete('/productos/:id', deleteAccount);
    
    api.post('/restaurantes', postRestaurantes);
    //api.put('/restaurantes/:id', putRestaurante);
    //api.get('/restaurantes/:id', getRestaurante);
    api.get('/restaurantes', getRestaurantes);
    //api.delete('/restaurantes/:id', deleteRestaurante);
}

var infrastructure = require('../../infrastructure');
var log = infrastructure.logger;
var errors = infrastructure.errors;

var accountService = require('../../application').accountService;

/**
 * Crea un nuevo producto y lo inserta en la base de datos.
 * @param {Object} req - req.body contiene el productoData menos el id
 * @param {Object} res - res.body contiene el producto insertado (including the id)
 */
function createAccount(req, res) {
	
    var accountData = req.body;

    accountService.createAccount(accountData)
        .then(function(account) {
            res.send(account);
            //res.send("ss");
            //console.log('hh');
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

/**
 * Actualiza un producto existente.
 * @param {Object} req - req.body contiene productoData incluyendo el id
 * @param {Object} res - res.body contiene el producto actualizado (incluyendo el id)
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
 * @param {Object} req - req.params.id contiene el id del producto a obtener
 * @param {Object} res - res.body contiene el producto solicitado
 */
function getAccount(req, res) {

    var id = req.params.id;
    console.log('gg2');

    accountService.getAccount(id)
        .then(function(account) {
            //res.send(account);
            res.send('gg');
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
 * @param {Object} req - no usado
 * @param {Object} res - res.body contiene un array de todos los productos
 */
function getAccounts(req, res) {
    accountService.getAccounts()
        .then(function(accounts) {
            //res.send(accounts);
            var Request = require("request");

Request.get("https://c5u2jk30nd.execute-api.us-west-2.amazonaws.com/nuevo/platos", (error, response, body) => {
    if(error) {
    	console.log('bb');
        return console.dir(error);
    }
    res.send(JSON.parse(body));
    console.log('xd');
    console.log(JSON.parse(body));
    console.dir(JSON.parse(body));
});
            
            
            
            
            
            //.send('ff');
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

/**
 * Elimina un producto.
 * @param {Object} req - req.params.id contiene el id del producto a eliminar
 * @param {Object} res - res.body no contiene nada
 */
function deleteAccount(req, res) {

    var id = req.params.id;

    accountService.deleteAccount(id)
        .then(function() {
            res.status(204).send({'message':'eliminado correctamente'});  // No Content
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

///////////////////////////////////////////////
// restaurantes
function getRestaurantes(req, res) {
    accountService.getAccounts()
        .then(function(accounts) {
            //res.send(accounts);
            var Request = require("request");
			Request.get(urldb+"platos", (error, response, body) => {
				if(error) {
					console.log('bb');
					return console.dir(error);
					//res.send(JSON.parse(body));
				}
				res.send(JSON.parse(body));
			});
            //.send('ff');
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

function postRestaurantes(req, res) {
    accountService.getAccounts()
        .then(function(accounts) {
            //res.send(accounts);
            var Request = require("request");
			Request.get(urldb+"platos", (error, response, body) => {
				if(error) {
					console.log('bb');
					return console.dir(error);
					//res.send(JSON.parse(body));
				}
				res.send(JSON.parse(body));
			});
            //.send('ff');
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}

// no va
function createAccount(req, res) {
	
    var accountData = req.body;

    accountService.createAccount(accountData)
        .then(function(account) {
            res.send(account);
            //res.send("ss");
            //console.log('hh');
        })
        .catch(function(error) {
            log.error(error);
            res.status(500).send({'message': error.toString()});
        });
}
