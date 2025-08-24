//JavaScript-клиент:

const socket = new WebSocket('ws://localhost:8080/ws');

socket.onopen = function(event) {
    console.log('Connected to WebSocket');
    socket.send('Hello from JavaScript!');
};

socket.onmessage = function(event) {
    console.log('Received:', event.data);
};

socket.onclose = function(event) {
    console.log('WebSocket connection closed');
};