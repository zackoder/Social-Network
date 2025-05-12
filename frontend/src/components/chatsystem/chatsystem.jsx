'use client';
import { useState } from 'react';
import styles from './chatsystem.module.css';
import Contacts from '../contacts/contacts';
import ChatBox from '../chatbox/chatbox';

export default function ChatSystem() {
    const [activeContact, setActiveContact] = useState(null);
    const [chatHistory, setChatHistory] = useState({}); // Stores messages per contact

    const handleContactClick = (contact) => {
        // If clicking the same contact, toggle chatbox
        if (activeContact?.id === contact.id) {
            setActiveContact(null);
        } else {
            setActiveContact(contact);
        }
    };

    const handleCloseChatBox = () => {
        setActiveContact(null);
    };

    // Update messages for the current contact
    const updateMessages = (newMessages) => {
        if (!activeContact) return;
        
        setChatHistory(prev => ({
            ...prev,
            [activeContact.id]: newMessages
        }));
    };

    return (
        <div className={styles.chatSystem}>
            <div className={styles.contactsContainer}>
                <Contacts 
                    onContactClick={handleContactClick} 
                    activeContactId={activeContact?.id}
                />
            </div>
            
            <div className={`${styles.chatboxContainer} ${activeContact ? styles.visible : ''}`}>
                {activeContact && (
                    <ChatBox 
                        key={activeContact.id} // Important for resetting state
                        contact={activeContact}
                        messages={chatHistory[activeContact.id] || []}
                        onUpdateMessages={updateMessages}
                        onClickClose={handleCloseChatBox}
                    />
                )}
            </div>
        </div>
    );
}