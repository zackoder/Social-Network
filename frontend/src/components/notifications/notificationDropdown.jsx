import { useEffect, useRef, useState } from "react";
import styles from "./notificationDropdown.module.css";
import { socket } from "../websocket/websocket";
import { isAuthenticated } from "@/app/page";
const host = process.env.NEXT_PUBLIC_HOST;

export const fetchNotifications = async () => {
  // if (loading) return;
  // setLoading(true);

  try {
    const response = await fetch(`${host}/getNotifications`, {
      credentials: "include",
    });
    const data = await response.json();
    // if (notifications.length === 0) {
    //   //setNotifications([...data]);

    // } else {
    //   setNotifications((prev) => [...prev, ...data]);
    // }
    return Array.isArray(data) ? data : [];

    // setHasMore(data.hasMore);
    // setOffset((prev) => prev + data.notifications.length);
  } catch (error) {
    console.log("Error fetching notifications:", error);
  }
};

export default function NotificationDropdown({
  isOpen,
  onClose,
  notifications,
}) {
  useEffect(() => {
    console.log("notifications", notifications);
  }, []);
  console.log("notifications", notifications);

  const dropdownRef = useRef(null);
  // const [notifications, setNotifications] = useState([]);
  const [responseData, setResponseData] = useState({ id: "", action: "" });

  // const [offset, setOffset] = useState(0);
  // const [hasMore, setHasMore] = useState(true);
  // const [loading, setLoading] = useState(false);

  // Load more notifications
  // const handleLoadMore = () => {
  //   fetchNotifications();
  // };

  // Close dropdown when clicking outside
  // useEffect(() => {
  //   const handleSocketMessage = (e) => {
  //     try {
  //       const data = JSON.parse(e.data);
  //       setNotifications((prev) => [data, ...prev]);
  //     } catch (err) {
  //       console.log("failed to notification: ", err);
  //     }
  //     socket.addEventListener("message", handleSocketMessage);
  //     if (data === null || Array.isArray(data.notification) ? data : []) {
  //       if (notifications.length === 0) {
  //         setNotifications([...data]);
  //       } else {
  //         setNotifications((prev) => [...prev, ...data]);
  //       }
  //     }

  //     return () => socket.removeEventListener("message", handleSocketMessage);
  //   };
  // }, []);

  // useEffect(() => {
  //   //if (isOpen) {
  //   fetchNotifications();
  //   //}
  // }, []);
  if (!isOpen) return null;

  useEffect(() => {
    const handleClickOutside = (event) => {
      if (dropdownRef.current && !dropdownRef.current.contains(event.target)) {
        onClose(); // call parent to close dropdown
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, [onClose]);

  //send Request
  useEffect(() => {
    const sendRequest = async () => {
      if (!responseData.id) return;
      try {
        console.log(responseData);

        const response = await fetch(`${host}/notiResp`, {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(responseData),
          credentials: "include",
        });
        console.log("response -------", responseData);

        if (!response.ok) {
          console.log("Failed to send notification");
          isAuthenticated(response.status);
          return;
        }
      } catch (err) {
        console.log("failed to send request the notification");
      }
    };
    sendRequest();
  }, [responseData.id]);

  return (
    <div ref={dropdownRef} className={styles.dropdown}>
      {/* <span className={styles.displaynotif}>{notifications.length}</span> */}
      <div className={styles.header}>
        <h3>Notifications</h3>
      </div>
      <div className={styles.notificationsList}>
        {notifications.length === 0 ? (
          <p>No notifications</p>
        ) : notifications ? (
          notifications.map((notification, index) =>
            notification.message === "event" ? (
              <div key={index} className={styles.notificationItem}>
                <p>{notification.type}</p>
                <button
                  onClick={() =>
                    setResponseData({ id: notification.id, message: "going" })
                  }
                  className={styles.btnNotification}
                >
                  Going
                </button>
                <button
                  onClick={() =>
                    setResponseData({
                      id: notification.id,
                      message: "not_going",
                    })
                  }
                  className={styles.btnNotification}
                >
                  Not Going
                </button>
              </div>
            ) : (
              <div key={index} className={styles.notificationItem}>
                <p>{notification.type}</p>
                <button
                  onClick={() =>
                    setResponseData({
                      id: notification.id,
                      message: "accepted",
                    })
                  }
                  className={styles.btnNotification}
                >
                  Accept
                </button>
                <button
                  onClick={() =>
                    setResponseData({
                      id: notification.id,
                      message: "rejected",
                    })
                  }
                  className={styles.btnNotification}
                >
                  Reject
                </button>
              </div>
            )
          )
        ) : (
          ""
        )}
      </div>
    </div>
  );
}
