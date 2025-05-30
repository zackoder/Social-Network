"use client";
import { socket } from "../websocket/websocket";
import "./groupChat.css";
import { useEffect, useState } from "react";
import { FaCloudUploadAlt } from "react-icons/fa";

export default function GroupChat({ groupData }) {
  const host = process.env.NEXT_PUBLIC_HOST;
  const [offset, setoffset] = useState(0);
  const [messages, setmessages] = useState([]);
  const [newmessage, setmessage] = useState("");
  const [image, setImage] = useState(null);
  const user_id = parseInt(localStorage.getItem("user-id"));

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) setImage(file);
    e.target.value = "";
  };

  useEffect(() => {
    async function fetchdata() {
      const resp = await fetch(
        `${host}/groupmessages?groupId=${groupData.Id}&offset=${offset}`,
        {
          credentials: "include",
        }
      );
      const data = await resp.json();
      console.log("data ----------------------------------------", data);
      setmessages(Array.isArray(data) ? data : []);
    }
    fetchdata();
  }, [groupData.Id, offset, host]);

  function handleMessage(e) {
    e.preventDefault();
    const content = e.target.children[0].value.trim();
    if (!content && image !== null) return;
    const message = {
      group_id: groupData.Id,
      content: content,
    };
    if (image) {
      message.mime = image.type;
      message.filename = image.name;
      const reader = new FileReader();
      reader.onload = () => {
        const messageBuffer = buildBinaryMessage(message, reader.result);
        socket.send(messageBuffer);
      };
      reader.readAsArrayBuffer(image);
    } else {
      console.log(message);
      socket.send(JSON.stringify(message));
    }
    setmessage("");
    setImage(null);
  }

  useEffect(() => {
    function handlereceivedmsgs(event) {
      const data = JSON.parse(event.data);
      if (data.group_id == groupData.Id) {
        setmessages((prev) => [...prev, data]);
      }
    }
    socket.addEventListener("message", handlereceivedmsgs);
    return () => {
      socket.removeEventListener("message", handlereceivedmsgs);
    };
  }, [groupData.Id]);

  return (
    <div className="chatcontainer">
      <div className="msgsContainer">
        {messages.length !== 0 ? (
          messages.map((message, index) => {
            return (
              <div key={index} className="message">
                <div className="infoMsg">
                  <img
                    className="avatar"
                    src={`http://${message.avatar}`}
                    alt="image Profile"
                  />
                  <h3 className="firstlastname">{`${message.first_name} ${message.last_name}`}</h3>
                </div>
                <div className="content">{message.content}</div>
              </div>
            );
          })
        ) : (
          <p className="empty">no messages yet</p>
        )}
      </div>
      <form className="submitForm" onSubmit={handleMessage}>
        <input
          value={newmessage}
          onChange={(e) => {
            setmessage(e.target.value);
          }}
          type="text"
        />
        <div className="buttons">
          <button type="submit" className="submit">
            submit
          </button>
          <input
            type="file"
            id="uploadImage"
            onChange={handleImageChange}
            className="hiddenInput"
          />
          <label htmlFor="uploadImage" className="uploadLabel">
            <FaCloudUploadAlt className="iconUpload" />
          </label>
        </div>
      </form>
    </div>
  );
}
