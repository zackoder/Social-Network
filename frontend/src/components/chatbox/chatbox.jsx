"use client";

import styles from "./chatbox.module.css";
import Image from "next/image";
import { useEffect, useRef, useState } from "react";
import { FaCloudUploadAlt } from "react-icons/fa";
import { IoIosSend } from "react-icons/io";
import { socket } from "../websocket/websocket";
import { isAuthenticated } from "@/app/page";

const host = process.env.NEXT_PUBLIC_HOST;

export default function ChatBox({ contact, onClickClose }) {

  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([]);
  const [image, setImage] = useState(null);
  const [showEmojis, setShowEmojis] = useState(false);
  const bottomRef = useRef(null);
  const userId = parseInt(localStorage.getItem("user-id"));

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

  const handleChange = (e) => setMessage(e.target.value);

  const handleImageChange = (e) => {

    const file = e.target.files[0];
    if (file) {
      setImage(file);
      e.target.value = "";
    }
  };

  useEffect(() => {
    const handleMessage = (event) => {
      try {
        const data = JSON.parse(event.data);

        if (
          (data.receiver_id === userId && data.sender_id === contact.id) ||
          (data.sender_id === userId && data.receiver_id === contact.id)
        ) {
          setMessages((prev) => [...prev, data]);
        }
      } catch (err) {
        console.error("Failed to parse message:", event.data);
      }
    };

    socket.addEventListener("message", handleMessage);

    return () => {
      // 🔥 Clean up the old listener to avoid duplicates
      socket.removeEventListener("message", handleMessage);
    };
  }, [contact.id]);

  useEffect(() => {
    const fetchMessages = async () => {
      try{
        const response = await fetch(`${host}/GetMessages?receiver_id=${contact.id}?offset=${}`);
        const data = await response.json();
        if (!response.ok){
          isAuthenticated(response.status, data.error);
        }
        setMessage(data);
      }catch(err){
        console.log("Error fetching messages", err);
        
      }

    };
    fetchMessages();
  }, [contact.id, userId]);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!message.trim() && image === null) return
    if (image) {
      const reader = new FileReader();
      reader.onload = () => {
        const metadata = {
          sender_id: userId,
          receiver_id: contact.id,
          type: "image",
          token: "online",
          content: message,
          mime: image.type,
          filename: image.name,
        };
        const messageBuffer = buildBinaryMessage(metadata, reader.result);
        socket.send(messageBuffer);
        // Add optimistic update for better UX
      };
      reader.readAsArrayBuffer(image);
      setImage(null);
      setMessage("");
      setImage(null);
    } else {
      const newMsg = {
        sender_id: userId,
        receiver_id: contact.id,
        type: "message",
        content: message,
        token: "online",
      };

      // Optimistic update
      socket.send(JSON.stringify(newMsg));
      setMessage("");
    }
  };

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

      <div className={styles.readmessages}>
        {messages
          .map((msg, index) => (
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
                {msg.filename ? (
                  <div className={styles.imageContainer}>
                    {console.log(
                      `link the image ${process.env.NEXT_PUBLIC_HOST}/uploads/${msg.filename}`
                    )}
                    {console.log(`msg content ${msg.filename}`)}
                    {/* const metadata = {
                                    sender_id: 1,
                                    receiver_id: contact.id,
                                    type: 'image',
                                    token: 'online',
                                    mime: image.type,
                                    filename: image.name,
                                    timestamp: new Date().toISOString()
                                    }; */}
                    <img
                      src={`${process.env.NEXT_PUBLIC_HOST}/uploads/${msg.filename}`}
                      alt="sent-image"
                      width={250} // Set appropriate dimensions
                      height={250}
                      className={styles.imageMessage}
                      onError={(e) => {
                        e.target.src = "/default-error-image.png"; // Add fallback image
                      }}
                    />
                    <span className={styles.timeStamp}>
                      {/* {new Date(msg.timestamp).toLocaleTimeString([], {
                        hour: "2-digit",
                        minute: "2-digit",
                      })} */}
                      {msg.creation_date}
                    </span>
                  </div>
                ) : (
                  <div className={styles.textMessage}>
                    <p>{msg.content}</p>
                    <span className={styles.timeStamp}>
                      {/* {new Date(msg.timestamp).toLocaleTimeString([], {
                        hour: "2-digit",
                        minute: "2-digit",
                      })} */}
                      {msg.creation_date}
                    </span>
                  </div>
                )}
              </div>
              {/* <div className={styles.message}>
                            <p>{msg.content}</p>
                            <span>{new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}</span>
                        </div> */}
              {msg.sender_id === userId && (
                <div className={styles.profileImage}>
                  <Image
                    src="/profile/profile.png"
                    alt="profile"
                    fill
                    style={{ objectFit: "cover", borderRadius: "50%" }}
                  />
                </div>
              )}
            </div>
          ))}
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
              onChange={handleChange}
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
