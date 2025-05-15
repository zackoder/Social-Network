import styles from "./contactprivate.module.css";

const listContacts = [
    {
        id: 1,
        name: "walid"
    },
    {
        id: 2,
        name: "zaki"
    },
    {
        id: 3,
        name: "ayoub"
    }
];

export default function ContactsPrivate(){
    return (
        <div className={styles.contacts}>
            {listContacts.map(contact => (
                <div className={styles.contact} key={contact.id}>
                    {contact.name}
                </div>
            ))}
        </div>

    );
}