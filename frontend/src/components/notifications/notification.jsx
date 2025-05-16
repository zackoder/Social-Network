import Link from "next/link";
import Image from "next/image";
import { links } from "./data";
import styles from "./notification.module.css"

export default function Notification(){
    return (
        <div className={styles.links}>
            {/* {links.map(link => <Link key={link.id} href={link.url} className={styles.link}>{link.title}</Link>)} */}
            {links.map(link => <Link key={link.id} href={link.url} className={styles.link}>
                {/* {link.title} */}
                <Image  
                    className={styles.image}
                    src = {`/images/${link.title}.png`}
                    width={30}
                    height={30}
                    title={link.title}
                    alt={link.title}
                />
            </Link>)}
        </div>
    );
}