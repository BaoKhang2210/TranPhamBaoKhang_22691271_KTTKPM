const http = require('http');
http.createServer((req, res) => {
  res.writeHead(200, {'Content-Type': 'text/plain'});
  res.end('Hello, day la ung dung Nodejs toi uu bang Multi-stage!');
}).listen(3000);
console.log('Server running at http://localhost:3000/');