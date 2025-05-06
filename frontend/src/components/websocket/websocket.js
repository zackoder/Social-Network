export let socket;

export function Websocket() {
    // if (!socket || socket.readyState === WebSocket.CLOSED) {

        // Create WebSocket connection.
        socket = new WebSocket("ws://localhost:8080/ws")

        // Connection opened
        socket.addEventListener("open", (event) => {
            // socket.send("Hello Server!");
            console.log("Websocket Connected");
            
        });

        // Listen for messages
        socket.addEventListener("message", (event) => {
            console.log("Message from server ", event.data);
        });
    
        //Websocket Closed
        socket.addEventListener('close', () => {
            console.log("WebSocket closed");
        });

        //Websocket error
        socket.addEventListener('error', (err) => {
            console.log("WebSocket error", err);
        });

    // }

}
