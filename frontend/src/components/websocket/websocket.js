export let socket;
export let websocketConnection = null;

export function Websocket() {

    if (websocketConnection) {
        return websocketConnection;
    }
    websocketConnection = new Promise((resolve, reject) => {
        // Create WebSocket connection.
        socket = new WebSocket("ws://localhost:8080/ws");

        // Connection opened
        socket.addEventListener("open", (event) => {
            // socket.send("Hello Server!");
            console.log("Websocket Connected");
            resolve(socket);

        });

        //Websocket Closed
        socket.addEventListener('close', (event) => {
            console.log("WebSocket closed");
            websocketConnection = null;
            reject(new Error('websocket connection closed'));
        });

        //Websocket error
        socket.addEventListener('error', (err) => {
            console.log("WebSocket error", err);
            reject(err);
        });

        // Listen for messages
        // socket.addEventListener("message", (event) => {
        //     console.log("Message from server ", event.data);
        //     let data = JSON.parse(event.data)
        //     if (data.reciever_id){
                
        //     }
        // });
    });
    return websocketConnection;

}


// type Message struct {
// 	Sender_id   int    `json:"sender_id"`
// 	Reciever_id int    `json:"reciever_id"`
// 	Type        string `json:"type"`
// 	Group_id    int    `json:"group_id"`
// 	Content     string `json:"content"`
// 	Mime        string `json:"mime"`
// 	Filename    string `json:"filename"`
// }