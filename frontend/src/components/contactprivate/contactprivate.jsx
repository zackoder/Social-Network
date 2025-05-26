"use client";
import React, {useContext, useState, useEffect } from "react";
// import styles from "./contactprivate.module.css"
import { DataContext } from "@/contexts/dataContext";
import { isAuthenticated } from "@/app/page";

export default function ContactsPrivate() {
  const [selectedContacts, setSelectedContacts] = useState([]);
  const {setSelectedContactsIds} = useContext(DataContext);
  const [contacts, setContacts] = useState([]);
  const host = process.env.NEXT_PUBLIC_HOST;
  useEffect(() => {
    const fetchFollowers = async () => {
      try {
        const response = await fetch(`${host}/api/getfollowers`, {
          credentials: "include",
        });
        const data = await response.json();
        setContacts(data)
        // selectedContacts(data);
        if (data && data.error) {          
          throw new Error(data.error);
        }
      } catch (error) {
        console.error("we can't fetch follower", error);
        isAuthenticated(response.status, "you should login first")
      }
    };
    fetchFollowers();
  }, []);
  if (!contacts || contacts.length == 0) {
    return;
  }
  const handleCheckboxChange = (name, id) => {
    setSelectedContacts((prev) =>
      prev.includes(name) ? prev.filter((n) => n !== name) : [...prev, name]
    );
    setSelectedContactsIds((prev)=> 
    prev.includes(id) ? prev.filter((n)=> n !== id) : [...prev,id]
  )
  };

  //  useEffect(() => {
  //   setSelectedContactsIds(selectedContacts); // send updated selection to parent
  // }, [selectedContacts]);
  
  

  return (
    <div style={{ position: "relative", width: "200px" }}>
      <div
        style={{
          border: "1px solid #ccc",
          padding: "5px",
          borderRadius: "4px",
          background: "#777",
          border: "none",
        }}
      >
        {selectedContacts.length > 0
          ? selectedContacts.join(", ")
          : "Select contacts"}
      </div>
      <div
        style={{
          border: "1px solid #ccc",
          padding: "8px",
          position: "absolute",
          background: "#111",
          zIndex: 1,
          borderRadius: "8px",
          display: "flex",
          gap: "10px",
          overflowY: "scroll",
        }}
      >
        {contacts.map((contact) => (          
          <label key={contact.id} style={{ display: "block" }}>
            {console.log("contact", contact)}
            <input
              style={{ marginRight: "10px" }}
              type="checkbox"
              checked={selectedContacts.includes(contact.firstName)}
              onChange={() => handleCheckboxChange(contact.firstName, contact.id)}
            />
            {contact.firstName}
          </label>
        ))}
      </div>
    </div>
  );
}
