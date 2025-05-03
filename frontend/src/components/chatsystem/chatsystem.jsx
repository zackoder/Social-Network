'use client'
import { useState } from "react"
import Contacts from "@/components/contacts/contacts.jsx"
import ChatBox from "../chatbox/chatbox";


export default function ChatSystem(){

    const [activeContact, setActiveContact] = useState(null);

    const handleContactClick = (contact) => {
        setActiveContact(contact);
    }

    const handleCloseChatBox = () => {
        setActiveContact(null);
    }

    return (
        <>
            <Contacts onContactClick={handleContactClick} />
            {activeContact && (<ChatBox contact={activeContact} onClickClose={() => setActiveContact(null)} />)}
        </>
    );

}