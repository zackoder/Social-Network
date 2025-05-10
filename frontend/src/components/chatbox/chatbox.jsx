'use client'

import styles from "./chatbox.module.css"
import Image from "next/image";
import { useEffect, useRef, useState } from "react";
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
    const [image, SetImage] = useState(null);
    const [showEmojis, setShowEmojis] = useState(false);
    const emojis = ['😁', '😅', '​🤣', '😂', '🙂', '🙃', '🫠', '😉', '🥰', '😍', '​🤩', '☺️', '​🥲', '😛', '​😜', '🤗', '🤭', '🤫', '🤔', '🫡', '​🫥', '​😒', '🙄', '🙂‍↔️', '​🙂‍↕️', '🥵', '🤯', '🥳', '😎', '​😎', '🤓', '🥺', '​🥹', '😥​', '😱​', '😭​', '👋​', '👌​', '🤞​', '👉​', '👇​', '👍​', '👏'];
    const addMessage = useRef()
    const handleEmojiClick = (emoji) => {
        setMessage((prev) => prev + emoji);
    };

    const handleChange = (e) => {
        setMessage(e.target.value);
    };

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            SetImage(file);
        }
    }

    const handleSubmit = async (e) => {
        e.preventDefault();

        if (image) {
            const reader = new FileReader();
            reader.onload = () => {
                const metadata = {
                    "sender_id": 1,
                    "reciever_id": 2,
                    "type": "image",
                    "token": "online",
                    // "content": reader.result,
                    "mime": image.type,
                    "filename": image.name,
                }
                const messageBuffer = buildBinaryMessage(metadata, reader.result);
                // socket.send(JSON.stringify(MsgImg));
                socket.send(messageBuffer);
            };
            // reader.readAsDataURL(imageFile);
            // reader.readAsDataURL(image);
            reader.readAsArrayBuffer(image);
            SetImage(null);
        }

        if (message.trim() !== "") {
            const Msg = {
                "sender_id": 1,
                "reciever_id": 2,
                "type": "message",
                "content": message,
                "token": "online"
            }
            console.log("send Message:", message);
            socket.send(JSON.stringify(Msg))
            // socket.send(JSON.stringify({type : "msg" , content: message}))

            setMessage("");
        }

    };

    socket.addEventListener("message", (event) => {
        console.log("Message from server ", event.data);
        let data = JSON.parse(event.data)
        if (data.reciever_id) {
            useEffect(() => {
                console.log(addMessage.current);

                if (addMessage.current) {
                    const newMessage = document.createElement("div")
                    newMessage.className = "me"
                    newMessage.textContent = data.content
                    addMessage.current.append(newMessage)
                }
            }, [])
        }
    });



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
            <div ref={addMessage} className={styles.readmessages}>

                {/* <div className={styles.me}>
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
                </div> */}
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

                        <input type="file" name="uploadImage" id="uploadImage" onChange={handleImageChange} className={styles.hiddenInput} />
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


function buildBinaryMessage(metadata, fileBuffer) {
    const meta = JSON.stringify(metadata) + "::";
    const encoder = new TextEncoder();
    const metaBuffer = encoder.encode(meta);
    // console.log("metaBuffer", metaBuffer);
    const combined = new Uint8Array(
        metaBuffer.length + fileBuffer.byteLength
    );
    combined.set(metaBuffer, 0);
    combined.set(new Uint8Array(fileBuffer), metaBuffer.length);
    return combined;
}
