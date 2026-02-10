const WebSocket = require('ws');
const net = require('net');

const wss = new WebSocket.Server({ port: 8080 });

// Replace with your ngrok address and port
const TCP_HOST = '0.tcp.ngrok.io';
const TCP_PORT = XXXXX; // your ngrok port

wss.on('connection', (ws) => {
  console.log('WebSocket client connected');

  // Connect to your TCP server
  const tcpSocket = net.createConnection({ host: TCP_HOST, port: TCP_PORT }, () => {
    console.log('Connected to TCP server');
  });

  // Forward TCP data to WebSocket
  tcpSocket.on('data', (data) => {
    ws.send(data.toString());
  });

  // Forward WebSocket messages to TCP server
  ws.on('message', (message) => {
    tcpSocket.write(message + '\n');
  });

  ws.on('close', () => {
    console.log('WebSocket disconnected');
    tcpSocket.end();
  });

  tcpSocket.on('end', () => {
    console.log('TCP connection closed');
    ws.close();
  });
});