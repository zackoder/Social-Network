'use client'

import styles from "./chatbox.module.css"
import Image from "next/image";
import { useState } from "react";
import { FaCloudUploadAlt } from "react-icons/fa";
import { IoIosSend } from "react-icons/io";
// import { socket } from "../websocket/websocket.js";
import { socket } from "../websocket/websocket";

// const Msg = {
//     "sender_id": 1,
//     "reciever_id": ,
//     "type": ,
//     "group_id": ,
//     "content": ,
//     "mime": ,
//     "filename": ,
// }

export default function ChatBox({ contact, onClickClose }) {
    const [message, setMessage] = useState('');
    const [showEmojis, setShowEmojis] = useState(false);
    const emojis = ['😁', '😅', '​🤣', '😂', '🙂', '🙃', '🫠', '😉', '🥰', '😍', '​🤩', '☺️', '​🥲', '😛', '​😜', '🤗', '🤭', '🤫', '🤔', '🫡', '​🫥', '​😒', '🙄', '🙂‍↔️', '​🙂‍↕️', '🥵', '🤯', '🥳', '😎', '​😎', '🤓', '🥺', '​🥹', '😥​', '😱​', '😭​', '👋​', '👌​', '🤞​', '👉​', '👇​', '👍​', '👏'];

    const handleEmojiClick = (emoji) => {
        setMessage((prev) => prev + emoji);
    };

    const handleChange = (e) => {
        setMessage(e.target.value);
    };


    const handleSubmit = (e) => {
        e.preventDefault();
        const Msg = {
            "sender_id": 1,
            "reciever_id": 2,
            "type": "message",
            "content": message,
            // "mime": ,
            // "filename": ,
        }
        console.log("send Message:", message);
        socket.send(JSON.stringify(Msg))
        // socket.send(JSON.stringify({type : "msg" , content: message}))

        setMessage("");
    };

    return (
        <div className={styles.chatBox}>
            <div className={styles.header}>
                <div className={styles.imgProfile}>
                    <Image
                        src="/profile/profile.png"
                        alt="image profile"
                        fill
                        style={{ objectFit: 'cover', borderRadius: '50%' }}
                    />
                </div>
                <div className={styles.infoProfile}>
                    <h3>{contact.name}</h3>
                    <p>Lorem ipsum dolor sit amet.</p>
                </div>
                <div className={styles.close} onClick={onClickClose}>x</div>
            </div>
            <div className={styles.readmessages}>
                <div className={styles.me}>
                    <div className={styles.message}>
                        <p>Lorem ipsum dolor sit amet.</p>
                        <span>8:55Am, Today</span>
                    </div>
                    <div className={styles.profileImage}>
                        <Image
                            src="/profile/profile.png"
                            alt="image profile"
                            fill
                            style={{ objectFit: 'cover', borderRadius: '50%' }}
                        />
                    </div>
                </div>
                <div className={styles.sender}>
                    <div className={styles.profileImage}>
                        <Image
                            src="/profile/profile.png"
                            alt="image profile"
                            fill
                            style={{ objectFit: 'cover', borderRadius: '50%' }}
                        />
                    </div>
                    <div className={styles.message}>
                        <p>Lorem ipsum dolor sit amet.</p>
                        <span>8:55Am, Today</span>
                    </div>
                </div>
            </div>
            <div className={styles.sendmessages}>
                <form onSubmit={handleSubmit}>
                    {/* Emoji picket */}
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
                            ))};
                        </div>
                    )};
                    <div className={styles.elementsSend}>
                        <button
                            type="button"
                            className={styles.emojiButton}
                            onClick={() => setShowEmojis(!showEmojis)}
                        >
                            😀
                        </button>
                        <input type="text" name="message" placeholder="Type your message..." value={message} onChange={handleChange} id="" />

                        <input type="file" name="uploadImage" id="uploadImage" className={styles.hiddenInput} />
                        <label htmlFor="uploadImage" className={styles.uploadLabel}>
                            <FaCloudUploadAlt className={styles.iconUpload} />
                        </label>

                        <input type="submit" name="submit" className={styles.hiddenInput} id="submit" />
                        <label htmlFor="submit" className={styles.labelSend}>
                            <IoIosSend className={styles.iconSend} />
                        </label>
                    </div>
                </form>
            </div>
        </div>
    );
}