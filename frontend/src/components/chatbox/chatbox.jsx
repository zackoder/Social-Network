"use client";

import styles from "./chatbox.module.css";
import Image from "next/image";
import { useEffect, useRef, useState } from "react";
import { FaCloudUploadAlt } from "react-icons/fa";
import { IoIosSend } from "react-icons/io";
import { socket } from "../websocket/websocket";
import { isAuthenticated } from "@/app/page";

const host = process.env.NEXT_PUBLIC_HOST;

function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`${name}=`);
  // console.log("------------------------------", parts.pop().split(";").shift());
  if (parts.length === 2) return parts.pop().split(";").shift();
}
const oldToken = getCookie("token");
console.log("--------------", oldToken);

export default function ChatBox({ contact, onClickClose }) {
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);
  const [image, setImage] = useState(null);
  const [showEmojis, setShowEmojis] = useState(false);
  const bottomRef = useRef(null);
  const scrollContainerRef = useRef(null);
  const [offset, setOffset] = useState(0);
  const limit = 10;
  const [userId, setUserId] = useState(null);

  const [error, seterror] = useState("");

  useEffect(() => {
    const fetchUserId = async () => {
      const response = await fetch(`${host}/userData`, {
        credentials: "include",
      });
      if (!response.ok) {
        isAuthenticated(response.status, "Please login first");
        return;
      }
      const data = await response.json();
      setUserId(data.id);
    };
    fetchUserId();
  }, []);

  useEffect(() => {
    if (!userId) return;

    const handleMessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        if (
          (data.sender_id === contact.id && data.receiver_id === userId) ||
          (data.sender_id === userId && data.receiver_id === contact.id)
        ) {
          setMessages((prev) => [...prev, data]);
          setOffset((prevOffset) => prevOffset + 1);
        }
        if (data.error) {
          seterror(data.error);
        }
      } catch (err) {
        console.log("Failed to parse message:", event.data);
      }
    };

    socket.addEventListener("message", handleMessage);
    return () => socket.removeEventListener("message", handleMessage);
  }, [contact.id, userId]);

  // const loadingRef = useRef(false);
  const offsetRef = useRef(0);
  const fetchMessages = async (offsetValue = 0) => {
    // if (loadingRef.current) return;
    // loadingRef.current = true;

    const container = scrollContainerRef.current;
    const previousScrollHeight = container?.scrollHeight || 0;

    try {
      const response = await fetch(
        `${host}/GetMessages?receiver_id=${contact.id}&offset=${offsetValue}`,
        {
          credentials: "include",
        }
      );

      const data = await response.json();

      if (!response.ok) {
        isAuthenticated(response.status, data.error);
        return;
      }

      if (Array.isArray(data) && data.length > 0) {
        setMessages((prev) => [...data.reverse(), ...prev]);
        // setOffset(offsetValue + limit);
        offsetRef.current = offsetValue + limit;

        setTimeout(() => {
          const newScrollHeight = container?.scrollHeight || 0;
          if (container) {
            container.scrollTop = newScrollHeight - previousScrollHeight;
          }
        }, 0);
      }
    } catch (err) {
      console.log("Error fetching messages", err);
    } finally {
      // loadingRef.current = false;
    }
  };

  useEffect(() => {
    setMessages([]);
    // setOffset(0);
    fetchMessages(0);
    offsetRef.current = 0;
  }, [contact.id]);

  const handleScroll = (e) => {
    if (e.target.scrollTop === 0) {
      //&& !loadingRef.current
      // console.log("Fetching with offset:", offsetRef.current);
      fetchMessages(offsetRef.current);
    }
  };

  useEffect(() => {
    const container = scrollContainerRef.current;
    if (!container) return;

    container.addEventListener("scroll", handleScroll);
    return () => container.removeEventListener("scroll", handleScroll);
  }, [offset]);

  useEffect(() => {
    bottomRef.current?.scrollIntoView({ behavior: "auto" });
  }, [messages]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    const token = getCookie("token");
    if (token !== oldToken) {
      window.location.reload();
      return;
    }
    if (!message.trim() && image === null) return;

    if (image) {
      const reader = new FileReader();
      reader.onload = () => {
        const metadata = {
          sender_id: userId,
          receiver_id: contact.id,
          type: "image",
          token,
          content: message,
          mime: image.type,
          filename: image.name,
        };
        const messageBuffer = buildBinaryMessage(metadata, reader.result);
        if (socket.readyState === WebSocket.OPEN) socket.send(messageBuffer);
      };
      reader.readAsArrayBuffer(image);
      setMessage("");
      setImage(null);
    } else {
      const newMsg = {
        sender_id: userId,
        receiver_id: contact.id,
        type: "message",
        content: message,
        token,
      };
      socket.send(JSON.stringify(newMsg));
      setMessage("");
    }
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      setImage(file);
      e.target.value = "";
    }
  };

  const emojis = [
    "😁",
    "😅",
    "🤣",
    "😂",
    "🙂",
    "🙃",
    "🫠",
    "😉",
    "🥰",
    "😍",
    "🤩",
    "☺️",
    "🥲",
    "😛",
    "😜",
    "🤗",
    "🤭",
    "🤫",
    "🤔",
    "🫡",
    "🫥",
    "😒",
    "🙄",
    "🙂‍↔️",
    "🙂‍↕️",
    "🥵",
    "🤯",
    "🥳",
    "😎",
    "🤓",
    "🥺",
    "🥹",
    "😥",
    "😱",
    "😭",
    "👋",
    "👌",
    "🤞",
    "👉",
    "👇",
    "👍",
    "👏",
  ];
  const handleEmojiClick = (emoji) => setMessage((prev) => prev + emoji);

  return (
    <div className={styles.chatBox}>
      <div className={styles.header}>
        <div className={styles.imgProfile}>
          <img
            src={`http://${contact.avatar}`}
            alt="image profile"
            width={50}
            height={50}
            style={{ objectFit: "cover", borderRadius: "50%" }}
          />
        </div>
        <div className={styles.infoProfile}>
          <h3>{`${contact.firstName} ${contact.lastName}`}</h3>
          <p>Chat started...</p>
        </div>
        <div className={styles.close} onClick={onClickClose}>
          x
        </div>
      </div>

      <div className={styles.readmessages} ref={scrollContainerRef}>
        {messages.map((msg, index) => (
          <div
            key={index}
            className={msg.sender_id === userId ? styles.me : styles.sender}
          >
            {msg.sender_id !== userId && (
              <div className={styles.profileImage}>
                <img
                  src={`http://${contact.avatar}`}
                  alt="profile"
                  width={50}
                  height={50}
                  style={{ objectFit: "cover", borderRadius: "50%" }}
                />
              </div>
            )}
            <div className={styles.message}>
              {msg.content && (
                <div className={styles.textMessage}>
                  <p>{msg.content}</p>
                </div>
              )}
              {msg.filename && (
                <div className={styles.imageContainer}>
                  <img
                    src={`http://${msg.filename}`}
                    alt="sent-image"
                    width={250}
                    height={250}
                    className={styles.imageMessage}
                  />
                </div>
              )}
              <span className={styles.timeStamp}>
                {formatDate(msg.creation_date)}
              </span>
            </div>
            {msg.sender_id === userId && (
              <div className={styles.profileImage}>
                <img
                  src="/profile/profile.png"
                  alt="profile"
                  width={50}
                  height={50}
                  style={{ objectFit: "cover", borderRadius: "50%" }}
                />
              </div>
            )}
          </div>
        ))}
        {error && <p className={styles.msgerror}>{error}</p>}
        <div ref={bottomRef} />
      </div>

      <div className={styles.sendmessages}>
        <form onSubmit={handleSubmit}>
          {showEmojis && (
            <div className={styles.emojiPicker}>
              {emojis.map((emoji, index) => (
                <button
                  key={index}
                  type="button"
                  className={styles.emojiButton}
                  onClick={() => handleEmojiClick(emoji)}
                >
                  {emoji}
                </button>
              ))}
            </div>
          )}
          <div className={styles.elementsSend}>
            <button
              type="button"
              className={styles.emojiButton}
              onClick={() => setShowEmojis(!showEmojis)}
            >
              😀
            </button>
            <input
              type="text"
              name="message"
              placeholder="Type your message..."
              value={message}
              onChange={(e) => setMessage(e.target.value)}
            />
            <input
              type="file"
              id="uploadImage"
              onChange={handleImageChange}
              className={styles.hiddenInput}
            />
            <label htmlFor="uploadImage" className={styles.uploadLabel}>
              <FaCloudUploadAlt className={styles.iconUpload} />
            </label>
            <input type="submit" id="submit" className={styles.hiddenInput} />
            <label htmlFor="submit" className={styles.labelSend}>
              <IoIosSend className={styles.iconSend} />
            </label>
          </div>
        </form>
      </div>
    </div>
  );
}

function buildBinaryMessage(metadata, fileBuffer) {
  const meta = JSON.stringify(metadata) + "::";
  const encoder = new TextEncoder();
  const metaBuffer = encoder.encode(meta);
  const combined = new Uint8Array(metaBuffer.length + fileBuffer.byteLength);
  combined.set(metaBuffer, 0);
  combined.set(new Uint8Array(fileBuffer), metaBuffer.length);
  return combined;
}

const oneday = 60 * 60 * 24;
const onehour = 60 * 60;
const oneminute = 60;

export function formatDate(time) {
  if (!time) time = Date.now() / 1000 - 1;
  const now = Date.now() / 1000;
  const elapsed = now - time;
  const days = Math.floor(elapsed / oneday);
  const hours = Math.floor((elapsed % oneday) / onehour);
  const minutes = Math.floor((elapsed % onehour) / oneminute);
  const seconds = Math.floor(elapsed % oneminute);

  if (days > 0) return `${days}d`;
  if (hours > 0) return `${hours}h`;
  if (minutes > 0) return `${minutes}min`;
  return `${seconds}s`;
}
