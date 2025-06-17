"use client";
import Link from "next/link";
import Image from "next/image";
import styles from "./notification.module.css";
import { useEffect, useState } from "react";
import NotificationDropdown, {
  fetchNotifications,
} from "./notificationDropdown";
import { socket } from "../websocket/websocket";

export default function Notification() {
  const [isOpen, setIsOpen] = useState(false);

  const toggleNotifications = (e) => {
    e.preventDefault();
    setIsOpen((prev) => !prev);
  };

  return (
    <div className={styles.notificationWrapper}>
      <div className={styles.image}>
        <Link href={"/groups"} className={styles.link}>
          <Image
            src="/images/groupes.png"
            alt="groups"
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
      <NotificationDropdown isOpen={isOpen} />
    </div>
  );
}
