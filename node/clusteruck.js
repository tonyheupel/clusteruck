#!/usr/bin/env node

var http = require('http');
var url = require('url');

function slowHandler(req, res) {
  setTimeout(function() {
    res.writeHead(200, { "Server": "Node.js/http.Server" });
    res.end('hello, world');
  }, 250);
}

http.createServer(slowHandler).listen(8081);
