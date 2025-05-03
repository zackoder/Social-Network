import Link from 'next/link';
import styles from './contacts.module.css'
import Image from 'next/image';

const dummyContact = {
    id: 1,
    name: 'Zack',
    image: '/images/profile.png',
};

export default function Contacts({onContactClick}) {
    return (
        <div className={styles.container}>
            {/* <Link href=""> */}
                <div className={styles.profile} onClick={() => onContactClick(dummyContact)}>
                    <div className={styles.imgProfile}>
                        <Image
                            src="/images/profile.png"
                            className={styles.image}
                            width={50}
                            height={50}
                            alt="profile"
                        />
                    </div>
                    <div className={styles.name}>
                        <h3>{dummyContact.name}</h3>
                    </div>
                </div>
            {/* </Link> */}
        </div>
    );
}