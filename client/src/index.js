import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import * as serviceWorker from './serviceWorker';

// https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API
const socket = new WebSocket('ws://localhost:8080/ws');
console.log(socket);

socket.onopen = (evt) => {
    console.log("OPEN");
    socket.send('echo test');
}

socket.onclose = (evt) => {
    console.log("CLOSE");
}

socket.onmessage = (evt) => {
    console.log("RESPONSE: " + evt.data);
}

socket.onerror = (evt) => {
    console.log("ERROR: " + evt.data);
}

ReactDOM.render(<App />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();
