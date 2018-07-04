'use strict';

var domain = require('../../domain');

var resultMaps = [
    {
        mapId: 'accountMap',
        createNew: function() {
            return new domain.Account();
        },
        properties: [
            {name: 'nombre', column: 'nombre'},
	    {name: 'precio', column: 'precio'},
	    {name: 'tipo', column: 'tipo'}
        ]
    }
];

module.exports = resultMaps;
