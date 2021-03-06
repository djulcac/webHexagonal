'use strict';

var dbConfig = {
    client: 'postgresql',
    debug: true,
    connection: {
        host: 'localhost',
        //user: 'bruno',
        //password: 'contrasena',
        user: 'postgres',
        password: '123456',
        database: 'demo',
        charset: 'utf8'
    }
};

var knex = require('knex')(dbConfig);

knex.schema

    /***** Drop foreign keys *****/
    // Example:
    // .table('transactions', function(table) {
    //     table.dropForeign('account_id');
    // })


    /***** Drop tables *****/
    .dropTableIfExists('productos')


    /***** Create tables (in alphabetic order) *****/
    // Accounts
    .createTable('productos', function(table) {
        table.increments('id');
        table.string('nombre', 64).notNullable().unique();
    })


    /***** Add foreign keys *****/
    // Example:
    // .table('transactions', function(table) {
    //     table.integer('account_id').unsigned().notNullable().references('accounts.id');
    //     table.integer('category_id').unsigned().notNullable().references('categories.id');
    // })


    /***** Destroy the database connection pool *****/
    .then(function() {
        knex.destroy();
    })


    // Finally, add a .catch handler for the promise chain
    .catch(function(e) {
        console.error(e);
    });
