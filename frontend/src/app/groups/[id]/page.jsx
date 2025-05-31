"use client";
import styles from "./id.module.css";
import { useParams, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

export default function GroupPage() {
  const [text, setText] = useState("");
  const [image, setImage] = useState(null);
  const [showPopup, setShowPopup] = useState(false);

  const [eventTitle, setEventTitle] = useState('');
  const [eventDescription, setEventDescription] = useState('');
  const [eventDatetime, setEventDatetime] = useState('');

  const [events, setEvents] = useState([]);
  const [posts, setPosts] = useState([]);
  const [groupData, setGroupData] = useState(null);

  const params = useParams();
  const { id } = params;
  const searchParams = useSearchParams();
  const host = process.env.NEXT_PUBLIC_HOST;

  const handleImageChange = (e) => {
    const file = e.target.files?.[0];
    if (file && file.type.startsWith("image/")) {
      setImage(file);
    }
  };

  const handleSubmit = async () => {
    console.log("hi");
    
    if (!text && !image) return;

    const formData = new FormData();
    formData.append("groupe_id",id);
    formData.append("content", text);
    if (image) {
      formData.append("image", image);
    }
    

    try {
      const response = await fetch(`${host}/addPost`, {
         credentials: "include",
         method: "POST",
        body: formData,
      });

      if (response.ok) {
        console.log("Post created successfully");
        setText("");
        setImage(null);
        // Re-fetch or update posts
      } else {
        alert("Failed to create post");
      }
    } catch (err) {
      console.error("Error:", err);
    }
  };

  useEffect(() => {
    async function getGroupData() {
      try {
        const resp = await fetch(`${host}/group?groupId=${id}`, {
          credentials: "include"
        });
        const data = await resp.json();
        setGroupData(data);
      } catch (err) {
        console.error("Failed to fetch group data:", err);
      }
    }

    getGroupData();
  }, [id, host]);

  const createEvent = () => {
    if (!eventTitle || !eventDatetime) return;

    setEvents([
      ...events,
      {
        title: eventTitle,
        description: eventDescription,
        datetime: eventDatetime,
        createdAt: new Date().toISOString(),
        responses: {}
      }
    ]);

    setEventTitle('');
    setEventDescription('');
    setEventDatetime('');
    setShowPopup(false);
  };

  const respondToEvent = (index, response) => {
    const user = prompt('Entrez votre nom :');
    if (user) {
      const newEvents = [...events];
      newEvents[index].responses[user] = response;
      setEvents(newEvents);
    }
  };

  if (!groupData) {
    return <div>Loading...</div>;
  }

  return (
    <div className={styles.parant}>
      <div className={styles.left}>
        <div className={styles.soutitre0}><p>All users</p></div>
        <div className={styles.chatbox0}></div>
      </div>

      <div className={styles.divcentral}>
        <div className={styles.supp}>
          <h1 className={styles.header}>{groupData.title}</h1>
          <p className={styles.description}>Description: {groupData.description}</p>
          <small className={styles.creator}>
            Creator: {groupData.first_name} {groupData.last_name}
          </small>
        </div>

        <div className={styles.moyyen}>
          <div className={styles.postsContainer}>
            {posts.map((post) => (
              <div key={post.id} className={styles.postCard}>
                <div>{post.title}</div>
                <div>{post.content}</div>
                {post.image && <img src={post.image} alt="Post" />}
                <small className={styles.postDate}>{post.date}</small>
              </div>
            ))}
          </div>

          <div className={styles.creatPost}>
            <div className={styles.textareaWrapper}>
              <textarea
                className={styles.input}
                placeholder="Share something..."
                value={text}
                onChange={(e) => setText(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === "Enter" && !e.shiftKey) {
                    e.preventDefault();
                    handleSubmit();
                  }
                }}
              ></textarea>

              <label htmlFor="imageUpload" className={styles.imageIcon}>📷</label>
              <input
                type="file"
                accept="image/*"
                id="imageUpload"
                className={styles.hiddenInput}
                onChange={handleImageChange}
              />
            </div>
          </div>
        </div>

        <div className={styles.infer}>
          {showPopup && (
            <div className={styles.overlay}>
              <div className={styles.popup}>
                <button
                  className={styles.closeButton}
                  onClick={() => setShowPopup(false)}
                >
                  ×
                </button>
                <h2 className={styles.Createvent}>Create an Event</h2>
                <input
                  className={styles.input2}
                  value={eventTitle}
                  onChange={(e) => setEventTitle(e.target.value)}
                  placeholder="Title"
                /><br /><br />
                <textarea
                  className={styles.input3}
                  value={eventDescription}
                  onChange={(e) => setEventDescription(e.target.value)}
                  placeholder="Description"
                  rows={2}
                ></textarea><br /><br />
                <input
                  className={styles.input4}
                  type="datetime-local"
                  value={eventDatetime}
                  onChange={(e) => setEventDatetime(e.target.value)}
                /><br /><br />
                <button className={styles.button} onClick={createEvent}>Create the Event</button>
              </div>
            </div>
          )}

          <button className={styles.addEventButton} onClick={() => setShowPopup(true)}>
            + Add Event
          </button>
        </div>
      </div>

      <div className={styles.right}>
        <div className={styles.soutitre}><p>Group chat</p></div>
        <div className={styles.chatbox1}></div>
      </div>
    </div>
  );
}
