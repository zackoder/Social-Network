"use client";
import { socket, Websocket } from "../websocket/websocket";
import "./groupChat.css";
import { useEffect, useState } from "react";
import { FaCloudUploadAlt } from "react-icons/fa";

export function displayChatbox() {
  const button = document.querySelector(".soutitre");
  const container = document.querySelector(".chatcontainer");
  // const formSubmit = document.querySelector(".submitForm");

  button?.addEventListener("click", () => {
    if (container?.classList.contains("showw")) {
      container.classList.remove("showw");
      container.classList.add("hide");
      // formSubmit?.classList.add("hide");

      // After animation ends, hide the element
      // button.addEventListener(
      //   "click",
      //   () => {
      //     if (container?.classList.contains("hide")) {
      //       container.style.display = "none";
      //       // formSubmit.style.display = "none";
      //     }
      //   },

      // );
    } else {
      container.classList.remove("hide");
      container.classList.add("showw");
      // formSubmit?.classList.add("showw");
      container.style.display = "block";
    }
  });
}

export default function GroupChat({ groupData }) {
  const host = process.env.NEXT_PUBLIC_HOST;
  const [offset, setoffset] = useState(0);
  const [messages, setmessages] = useState([]);
  const [newmessage, setmessage] = useState("");
  const [image, setImage] = useState(null);
  // const user_id = parseInt(localStorage.getItem("user-id"));

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
      setmessages(Array.isArray(data) ? data : []);
    }
    fetchdata();
  }, [groupData.Id, offset, host]);

  function handleMessage(e) {
    e.preventDefault();
    const content = e.target.children[0].value.trim();

    if (!content && image === null) return;
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
        Websocket().then((socket) => {
          if (socket.readyState === WebSocket.OPEN) socket.send(messageBuffer);
        });
      };
      reader.readAsArrayBuffer(image);
    } else {
      console.log(message);
      Websocket()
        .then((socket) => {
          if (socket.readyState === WebSocket.OPEN)
            socket.send(JSON.stringify(message));
          else console.log("hadshi ma5damsh");
        })
        .catch((err) => {
          console.log("ash hadshi ", err);
        });
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
    <>
      <button className="soutitre" onClick={displayChatbox}>
        <p className="titleGroup">Group chat</p>
      </button>
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
                  <div className="content">
                    {/* {message.content} */}
                    {message.filename !== "" || message.content !== "" ? (
                      // <img src={message.filename} alt="Image" />
                      <div>
                        {message.content !== "" ? <p>{message.content}</p> : ""}
                        {message.filename !== "" ? (
                          <img
                            src={`${process.env.NEXT_PUBLIC_HOST}${message.filename}`}
                            alt="Image"
                            width={250} // Set appropriate dimensions
                            height={250}
                          // className={styles.imageGroupChat}
                          />
                        ) : (
                          ""
                        )}
                      </div>
                    ) : (
                      ""
                    )}
                  </div>
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
            <input
              type="file"
              id="uploadGroupImage"
              onChange={handleImageChange}
              className="hiddenInput"
            />
            <label htmlFor="uploadGroupImage" className="uploadLabel">
              <FaCloudUploadAlt className="iconUpload" />
            </label>
            <button type="submit" className="submit">
              submit
            </button>
          </div>
        </form>
      </div>
    </>
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
