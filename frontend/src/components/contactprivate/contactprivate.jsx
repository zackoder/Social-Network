// import styles from "./contactprivate.module.css";

// const listContacts = [
//     {
//         id: 1,
//         name: "walid"
//     },
//     {
//         id: 2,
//         name: "zaki"
//     },
//     {
//         id: 3,
//         name: "ayoub"
//     }
// ];

// export default function ContactsPrivate() {
//     return (
//         <div className={styles.contacts}>
//             <select name="" id="">
//                 {listContacts.map(contact => (
//                     <option value={`${contact.name}`} key={contact.id}>{contact.name}</option>
//                 ))}
//             </select>
//         </div>

//     );
// }

'use client';
import React, { useState } from 'react';

const listContacts = [
  { id: 1, name: "walid" },
  { id: 2, name: "zaki" },
  { id: 3, name: "ayoub" },
];

export default function ContactsPrivate() {
  const [selectedContacts, setSelectedContacts] = useState([]);

  const handleCheckboxChange = (name) => {
    setSelectedContacts(prev =>
      prev.includes(name)
        ? prev.filter(n => n !== name)
        : [...prev, name]
    );
  };

  return (
    <div style={{ position: 'relative', width: '200px' }}>
      <div style={{ border: '1px solid #ccc', padding: '5px', borderRadius: '8px', background: '#777', border: 'none' }}>
        {selectedContacts.length > 0 ? selectedContacts.join(', ') : 'Select contacts'}
      </div>
      <div style={{ border: '1px solid #ccc', padding: '8px', position: 'absolute', background: '#111', zIndex: 1,  borderRadius: '8px' }}>
        {listContacts.map(contact => (
          <label key={contact.id} style={{ display: 'block'}}>
            <input
              style={{marginRight: '10px'}}
              type="checkbox"
              checked={selectedContacts.includes(contact.name)}
              onChange={() => handleCheckboxChange(contact.name)}
            />
            {contact.name}
          </label>
        ))}
      </div>
    </div>
  );
}
