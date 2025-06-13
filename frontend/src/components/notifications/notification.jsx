"use client";
import Link from "next/link";
import Image from "next/image";
import styles from "./notification.module.css";
import { useState } from "react";
import NotificationDropdown from "./notificationDropdown";

export default function Notification() {
    // const [isNotificationOpen, setIsNotificationOpen] = useState(false);
    const [isOpen, setIsOpen] = useState(false);

    const toggleNotifications = (e) => {
        e.preventDefault();
        setIsOpen((prev) => !prev);
    }

    return (
        
        <div className={styles.notificationWrapper}>
            <div className={styles.image}>
                <Link href={"/groups"} className={styles.link} >
                    <Image
                        src="/images/groupes.png"
                        alt="groups"
                        // className={styles.image}
                        width={30}
                        height={30}
                        title="Groups"
                    />
                </Link>
            </div>
            <div onClick={toggleNotifications} className={styles.image}>
                <Image
                    src="/images/notification.png"
                    alt="notification"
                    className={styles.link}
                    width={40}
                    height={40}
                    title="Notification"
                />
            </div>
            {isOpen && (
                <NotificationDropdown isOpen={isOpen} onClose={() => { setIsOpen(false) }} />
            )}
        </div>
    );
}