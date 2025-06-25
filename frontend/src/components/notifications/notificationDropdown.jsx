import { useEffect, useRef, useState } from "react";
import styles from "./notificationDropdown.module.css";
import { socket } from "../websocket/websocket";
import { isAuthenticated } from "@/app/page";
const host = process.env.NEXT_PUBLIC_HOST;

export const fetchNotifications = async () => {
  try {
    const response = await fetch(`${host}/getNotifications`, {
      credentials: "include",
    });
    const data = await response.json();
    return Array.isArray(data) ? data : [];
  } catch (error) {
    console.log("Error fetching notifications:", error);
  }
};

export default function NotificationDropdown({ isOpen }) {
  // const [isOpen, setIsOpen] = useState(false);
  const [notifications, setNotifications] = useState([]);
  const [notisnumber, setnotisnumber] = useState("");

  useEffect(() => {
    const handleSocketMessage = (e) => {
      try {
        const data = JSON.parse(e.data);

        if (data && !data.message) return;
        // setNotifications((prev) => [data, ...prev]);

        if (notifications.length === 0) {
          setNotifications([data]);
        } else {
          setNotifications([data, ...notifications]);
        }

        setnotisnumber(notisnumber + 1);
      } catch (err) {
        console.log("failed to notification: ", err);
      }

      return () => socket.removeEventListener("message", handleSocketMessage);
    };
    socket.addEventListener("message", handleSocketMessage);
  }, [notisnumber]);

  useEffect(() => {
    (async function () {
      const data = await fetchNotifications();
      setNotifications(Array.isArray(data) ? data : []);
      setnotisnumber(data.length);
      console.log("notifications", data);
    })();
  }, []);
  const dropdownRef = useRef(null);

  const [responseData, setResponseData] = useState({ id: "", action: "" });

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

        if (!response.ok) {
          console.log("Failed to send notification");
          isAuthenticated(response.status);
          return;
        }
        setnotisnumber(notisnumber - 1);
        setNotifications(() => {
          if (!notifications && notifications.length <= 1) {
            return [];
          }
          return notifications.filter((not) => {
            return not.id != responseData.id;
          });
        });
      } catch (err) {
        console.log("failed to send request the notification");
      }
    };
    sendRequest();
  }, [responseData.id]);

  return (
    <>
      <span className={styles.displaynotif}>{notisnumber}</span>
      {isOpen && (
        <div ref={dropdownRef} className={styles.dropdown}>
          <div className={styles.header}>
            <h3>Notifications</h3>
          </div>
          <div className={styles.notificationsList}>
            {notifications.length === 0 ? (
              <p>No notifications</p>
            ) : (
              notifications &&
              notifications.map((notification, index) =>
                notification.message === "event" ? (
                  <div key={index} className={styles.notificationItem}>
                    <p>{notification.type}</p>
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
            )}
          </div>
        </div>
      )}
    </>
  );
}
