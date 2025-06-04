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




    //   const getImageLink = async (path) => {
//     try {
//       const res = await fetch(`http://localhost:3000/images/${path}.png`);
//       if (res.ok) {
//         const blob = await res.blob();
//         const imageURL = URL.createObjectURL(blob);
//         return imageURL;
//       } else {
//         throw new Error(`Failed to fetch image: ${res.status}`);
//       }
//     } catch (error) {
//       console.error("Error fetching image:", error);
//       return null;
//     }
//   };

//   useEffect(() => {
//     links.map( (link) => {
//         console.log("sssss hada ana",link);
//     //   const url = getImageLink(link.title);
//       link.title = url;
      
//     });
//   });

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
                            title= {link.title}
                            alt="Groups"
                        />
                    </Link>
                     {link.title === "notification" && ( 
                        <NotificationDropdown 
                            isOpen={isNotificationOpen} 
                            onClose={() => setIsNotificationOpen(false)}
                        />                     )}
                </div>
             ))}
        </div>
    );
}