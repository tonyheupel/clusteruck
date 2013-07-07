#!/usr/bin/env node

var cluster = require('cluster');
var http = require('http');
var url = require('url');
var numCPUs = require('os').cpus().length;

function slowHandler(req, res) {
  setTimeout(function() {
    res.writeHead(200);
    res.end('hello, world');
  }, 250);
}

function fastHandler(req, res) {
    res.writeHead(200);
    res.end('hello, world');
}

function resizeHandler(req, res) {
  var urlParts = url.parse(req.url, true);
  var workers = parseInt(urlParts.query['workers']);

  forkWorkers(workers);
  res.writeHead(200);
  res.end('Resized to ' + Object.keys(cluster.workers).length);
}

function forkWorkers(numWorkers) {
  var currentWorkers = cluster.workers ? Object.keys(cluster.workers).length : 0;

  if (numWorkers === currentWorkers) {
    console.log('Already at ' + numWorkers +'...');
  } else if (numWorkers > currentWorkers) {
    addWorkers(numWorkers - currentWorkers);
  } else {
    removeWorkers(currentWorkers - numWorkers)
  }
  console.log('Now at ' + Object.keys(cluster.workers).length + ' worker(s)');
}

function addWorkers(numToAdd) {
  for (var i = 0; i < numToAdd; i++) {
    console.log('Forking new worker #' + (i + 1) + '...')
    cluster.fork();
  }
}

function removeWorkers(numToRemove) {
  var workerIds = Object.keys(cluster.workers);
  var numRemoved = 0;

  while (numRemoved < numToRemove) {
    var lastId = workerIds[workerIds.length - (numRemoved + 1)];

    console.log('Killing worker id ' + lastId + '...');
    cluster.workers[lastId].kill();
    delete cluster.workers[lastId];

    numRemoved = numRemoved + 1;
  }
}

if (cluster.isMaster) {
  // Set initial worker size
  forkWorkers(numCPUs);

  cluster.on('exit', function(worker, code, signal) {
    console.log('worker ' + worker.process.pid + ' died');
  });

  http.createServer(resizeHandler).listen(8079);
} else {
  // Workers can share any TCP connection
  // In this case its a HTTP server
  http.createServer(slowHandler).listen(8080);
}
