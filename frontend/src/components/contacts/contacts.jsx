// import Link from 'next/link';
// import styles from './contacts.module.css'
// import Image from 'next/image';

// // const dummyContact = {
// //         id: 1,
// //         name: 'Zack',
// //         image: '/images/profile.png',
// // }

// const dummyContact = [
//     {
//         id: 1,
//         name: 'Zack',
//         image: '/images/profile.png',
//     },
//     {
//         id: 2,
//         name: 'Walid',
//         image: '/images/profile.png',
//     },
//     {
//         id: 3,
//         name: 'Med',
//         image: '/images/profile.png',
//     }
// ];

// export default function Contacts({ onContactClick, activeContactId }) {
//     return (
//         <div className={styles.container}>
//             {dummyContact.map(contact => (
//                 <div 
//                     key={contact.id}
//                     className={`${styles.profile} ${activeContactId === contact.id ? styles.active : ''}`}
//                     onClick={() => onContactClick(contact)}
//                     role="button"
//                     tabIndex={0}
//                     onKeyDown={(e) => e.key === 'Enter' && onContactClick(contact)}
//            
//                     <div className={styles.imgProfile}>
//                         <Image
//                             src={contact.image}
//                             width={50}
//                             height={50}
//                             alt={`${contact.name}'s profile`}
//                         />
//                     </div>
//                     <div className={styles.contactInfo}>
//                         <h3>{contact.name}</h3>
//                         <p className={styles.lastMessage}>{contact.lastMessage}</p>
//                         <span className={styles.time}>{contact.lastMessageTime}</span>
//                     </div>
//                 </div>
//             ))}
//         </div>
//     );
// }




import { useEffect, useState } from 'react';
import Image from 'next/image';
import styles from './contacts.module.css';

export default function Contacts({ onContactClick, activeContactId }) {
    const [contacts, setContacts] = useState([]);
    const host = process.env.NEXT_PUBLIC_HOST;

    useEffect(() => {
        // This runs when the component is mounted
        const fetchContacts = async () => {
            try {
                const response = await fetch(`${host}/api/getuserfriends`);
                const data = await response.json();
                // console.log(data.avatar);
                // Optional: format to match your current structure
                const formattedContacts = data.map((friend, index) => ({
                    id: index + 1, // or use friend.id if your API returns it
                    name: `${friend.firstName} ${friend.lastName}`,
                    image: `/${friend.avatar}`, // assuming avatar is like "uploads/layla.png"
                    lastMessage: "Hey there!", // placeholder
                    lastMessageTime: "2h ago", // placeholder
                }));

                setContacts(formattedContacts);
            } catch (error) {
                console.error('Failed to fetch contacts:', error);
            }
        };
        // console.log(contacts);

        fetchContacts();
    }, []);
    return (
        <div className={styles.container}>
            {contacts.map(contact => (
                <div 
                    key={contact.id}
                    className={`${styles.profile} ${activeContactId === contact.id ? styles.active : ''}`}
                    onClick={() => onContactClick(contact)}
                    role="button"
                    tabIndex={0}
                    onKeyDown={(e) => e.key === 'Enter' && onContactClick(contact)}
                >
                    <div className={styles.imgProfile}>
                        <Image
                            src={contact.image}
                            width={50}
                            height={50}
                            alt={`${contact.name}'s profile`}
                        />
                    </div>
                    <div className={styles.contactInfo}>
                        <h3>{contact.name}</h3>
                        <p className={styles.lastMessage}>{contact.lastMessage}</p>
                        <span className={styles.time}>{contact.lastMessageTime}</span>
                    </div>
                </div>
            ))}
        </div>
    );
}
