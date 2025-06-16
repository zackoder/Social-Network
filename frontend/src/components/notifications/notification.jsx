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
  // const [isNotificationOpen, setIsNotificationOpen] = useState(false);
  const [isOpen, setIsOpen] = useState(false);
  const [notifications, setNotifications] = useState([]);
  const [notisnumber, setnotisnumber] = useState("");

  const toggleNotifications = (e) => {
    e.preventDefault();
    setIsOpen((prev) => !prev);
  };

  useEffect(() => {
    (async function () {
      const data = fetchNotifications();
      const notis = await data;
      setNotifications(notis);
      setnotisnumber(notis.length);
    })();

    //if (isOpen) {

    //}
  }, []);

  useEffect(() => {
    const handleSocketMessage = (e) => {
      try {
        const data = JSON.parse(e.data);

        // setNotifications((prev) => [data, ...prev]);
        if (notifications.length === 0) {
          setNotifications(Array.isArray(data) ? data : []);
        } else {
          setNotifications((prev) => [...prev, ...data]);
        }

        setnotisnumber(notisnumber + 1);

        console.log("notisnumber", notisnumber);
      } catch (err) {
        console.log("failed to notification: ", err);
      }

      return () => socket.removeEventListener("message", handleSocketMessage);
    };
    socket.addEventListener("message", handleSocketMessage);
  }, [notisnumber]);

  return (
    <div className={styles.notificationWrapper}>
      <div className={styles.image}>
        <Link href={"/groups"} className={styles.link}>
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
        <span className={styles.displaynotif}>{notifications.length}</span>
        <Image
          src="/images/notification.png"
          alt="notification"
          className={styles.link}
          width={40}
          height={40}
          title="Notification"
        />
      </div>
      <span className={styles.displaynotif}>{notisnumber}</span>
      {isOpen && (
        <NotificationDropdown
          isOpen={isOpen}
          notifications={notifications}
          onClose={() => {
            setIsOpen(false);
          }}
        />
      )}
    </div>
  );
}
