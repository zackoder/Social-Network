'use client'
import { useEffect } from "react"
import { Websocket } from "./websocket"

export default function WebsocketInitializer(){
    useEffect(() => {
        Websocket()
    }, []);
    return null
}