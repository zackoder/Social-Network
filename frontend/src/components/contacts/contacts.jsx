"use client";

import { useEffect, useState } from "react";
import Image from "next/image";
import { isAuthenticated } from "@/app/page";
import styles from "./contacts.module.css";

export default function Contacts({ onContactClick, activeContactId }) {
  const [contacts, setContacts] = useState([]);
  const host = process.env.NEXT_PUBLIC_HOST;

  useEffect(() => {
    // This runs when the component is mounted
    const fetchContacts = async () => {
      try {
        const response = await fetch(`${host}/api/getuserfriends`, {
          credentials: "include",
        });
        const data = await response.json();
        setContacts(data);
        if (data &&data.error) {
          // throw new Error(data.error);
          console.log(data.error);
        }
        console.log(data);
      } catch (error) {
        console.error("Failed to fetch contacts:", error);
        isAuthenticated(response.status, "you should login first")
        
      }
    };
    fetchContacts();
  }, []);
  if (!contacts || contacts.length == 0) {
    return;
  }

  return (
    <div className={styles.container}>
      {contacts.map((contact) => (
        <div
          key={contact.id}
          className={`${styles.profile} ${
            activeContactId === contact.id ? styles.active : ""
          }`}
          onClick={() => onContactClick(contact)}
          role="button"
          tabIndex={0}
          onKeyDown={(e) => e.key === "Enter" && onContactClick(contact)}
        >
          <div className={styles.imgProfile}>
            <Image
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
      ))}
    </div>
  );
}
