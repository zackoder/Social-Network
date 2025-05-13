import Link from 'next/link';
import styles from './contacts.module.css'
import Image from 'next/image';

// const dummyContact = {
//         id: 1,
//         name: 'Zack',
//         image: '/images/profile.png',
// }

const dummyContact = [
    {
        id: 1,
        name: 'Zack',
        image: '/images/profile.png',
    },
    {
        id: 2,
        name: 'Walid',
        image: '/images/profile.png',
    },
    {
        id: 3,
        name: 'Med',
        image: '/images/profile.png',
    }
];

export default function Contacts({ onContactClick, activeContactId }) {
    return (
        <div className={styles.container}>
            {dummyContact.map(contact => (
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