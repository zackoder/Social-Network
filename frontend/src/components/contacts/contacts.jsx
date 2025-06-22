"use client";

import { useEffect, useRef, useState } from "react";
import { isAuthenticated } from "@/app/page";
import styles from "./contacts.module.css";
import { socket } from "../websocket/websocket";

export default function Contacts({ onContactClick, activeContactId }) {
  const [contacts, setContacts] = useState([]);
  const currentContactRef = useRef(activeContactId);
  const host = process.env.NEXT_PUBLIC_HOST;

  useEffect(() => {
    currentContactRef.current = activeContactId;
  }, [activeContactId]);

  useEffect(() => {
    const fetchContacts = async () => {
      try {
        const response = await fetch(`${host}/api/getuserfriends`, {
          credentials: "include",
        });
        const data = await response.json();

        if (!response.ok) {
          isAuthenticated(response.status, "you should login first");
        }

        if (Array.isArray(data)) {
          // Add 'notis' field to each contact
          const contactsWithNotifications = data.map((contact) => ({
            ...contact,
            notis: 0,
          }));
          setContacts(contactsWithNotifications);
        } else {
          console.log(
            "Failed to load contacts:",
            data?.error || "Unknown error"
          );
        }
      } catch (error) {
        console.log("Failed to fetch contacts:", error);
      }
    };
    fetchContacts();
  }, []);

  function handleMessageNotification(e) {
    try {
      const data = JSON.parse(e.data);

      setContacts((prevContacts) =>
        prevContacts.map((contact) => {
          if (contact.id === data.sender_id) {
            const isCurrent = currentContactRef.current === data.sender_id;
            return {
              ...contact,
              notis: isCurrent ? 0 : (contact.notis || 0) + 1,
            };
          }
          return contact;
        })
      );
    } catch (err) {
      console.log("Failed to handle message notification:", err);
    }
  }

  useEffect(() => {
    setContacts((prevContacts) =>
      prevContacts.map((contact) =>
        contact.id === activeContactId ? { ...contact, notis: 0 } : contact
      )
    );
  }, [activeContactId]);

  useEffect(() => {
    socket.addEventListener("message", handleMessageNotification);
    return () => {
      socket.removeEventListener("message", handleMessageNotification);
    };
  }, []);

  return (
    <div className={styles.container}>
      {contacts.map((contact) => (
        <div className={styles.contactContainer} key={contact.id}>
          {contact.notis > 0 && (
            <span className={styles.displaynotif}>{contact.notis}</span>
          )}
          <div
            className={`${styles.profile} ${
              activeContactId === contact.id ? styles.active : ""
            }`}
            onClick={() => onContactClick(contact)}
            role="button"
            tabIndex={0}
            onKeyDown={(e) => e.key === "Enter" && onContactClick(contact)}
          >
            <div className={styles.imgProfile}>
              <img
                src={`http://${contact.avatar}`}
                width={50}
                height={50}
                style={{ borderRadius: "100%" }}
                alt={`${contact.firstName} ${contact.lastName}'s profile`}
              />
            </div>
            <div className={styles.contactInfo}>
              <h3>{`${contact.firstName} ${contact.lastName}`}</h3>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
