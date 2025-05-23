"use client";
import Link from "next/link";
import Image from "next/image";
import { links } from "./data";
import styles from "./notification.module.css";
import { useState } from "react";
import NotificationDropdown from "./notificationDropdown";

export default function Notification() {
    const [isNotificationOpen, setIsNotificationOpen] = useState(false);

    const toggleNotifications = (e) => {
        // Prevent navigation for notification icon
        if (e.currentTarget.title === "notification") {
            e.preventDefault();
            setIsNotificationOpen(prev => !prev);
        }
    };

    return (
        <div className={styles.links}>
            {links.map(link => (
                <div key={link.id} className={styles.linkContainer}>
                    <Link 
                        href={link.url} 
                        className={styles.link} 
                        onClick={toggleNotifications}
                    >
                        <Image  
                            className={styles.image}
                            src={`/images/${link.title}.png`}
                            width={30}
                            height={30}
                            title={link.title}
                            alt={link.title}
                        />
                    </Link>
                    {link.title === "notification" && (
                        <NotificationDropdown 
                            isOpen={isNotificationOpen} 
                            onClose={() => setIsNotificationOpen(false)}
                        />                    )}                
                </div>
            ))}
        </div>
    );
}