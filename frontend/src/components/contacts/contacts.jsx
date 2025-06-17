"use client";

import { useEffect, useRef, useState } from "react";
import Image from "next/image";
import { isAuthenticated } from "@/app/page";
import styles from "./contacts.module.css";
import { socket } from "../websocket/websocket";

export default function Contacts({ onContactClick, activeContactId }) {
  const [contacts, setContacts] = useState([]);
  const host = process.env.NEXT_PUBLIC_HOST;
  const contactref = useRef([]);
  const currentContactRef = useRef(activeContactId);
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
        setContacts(Array.isArray(data) ? data : []);
        contactref.current = Array.isArray(data) ? data : [];
        if (data && data.error) {
          // throw new Error(data.error);
          console.log(data.error);
        }
        if (!response.ok) {
          isAuthenticated(response.status, "you should login first");
        }
      } catch (error) {
        console.log("Failed to fetch contacts:", error);
      }
    };
    fetchContacts();
  }, []);

  function handlemessageNotification(e) {
    const data = JSON.parse(e.data);
    console.log("contacts", data);
    contactref.current.forEach((contact) => {
      if (
        currentContactRef.current === data.sender_id &&
        data.sender_id === contact.id
      ) {
        console.log("hello", data);
      } else if (contact.id === data.sender_id) {
        console.log("create a notification");
      }
    });
  }
  useEffect(() => {
    socket.addEventListener("message", handlemessageNotification);
    return () => {
      socket.removeEventListener("message", handlemessageNotification);
    };
  }, []);

  return (
    <div className={styles.container}>
      {contacts.map((contact) => (
        <div key={contact.id}>
          <span className={styles.displaynotif}>{}</span>
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
                alt={`${contact.name}'s profile`}
              />
            </div>
            <div className={styles.contactInfo}>
              <h3>{`${contact.firstName} ${contact.lastName} `}</h3>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
