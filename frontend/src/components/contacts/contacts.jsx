import Link from 'next/link';
import styles from './contacts.module.css'
import Image from 'next/image';

export default function Contacts() {
    return (
        <div className={styles.container}>
            <Link href="/profile">
                <div className={styles.profile}>
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
                        <h3>full Name</h3>
                    </div>
                </div>
            </Link>
        </div>
    );
}