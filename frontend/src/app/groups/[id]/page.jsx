"use client";
import GroupChat from "@/components/groupchat/groupchat";
import styles from "./id.module.css";
import { useParams } from "next/navigation";
import { useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

//  const [posts, setPosts] = useState([]);
//   const [events, setEvents] = useState([]);
//   const [postContent, setPostContent] = useState('');

export default function GroupPage() {
  const [showPopup, setShowPopup] = useState(false);

  const [eventTitle, setEventTitle] = useState("");
  const [eventDescription, setEventDescription] = useState("");
  const [eventDatetime, setEventDatetime] = useState("");
  const params = useParams(); // récupère params dynamiques
  const { id } = params; // ici id correspond au nom du fichier dynamique [id]
  const posts = [];
  console.log({ id });

  const host = process.env.NEXT_PUBLIC_HOST;

  // Pour récupérer les query params (comme title, description) passés dans l'URL,
  // useSearchParams est utile:
  const searchParams = useSearchParams();
  const [groupData, setgroupdata] = useState(null);
  useEffect(() => {
    async function getGroupData() {
      try {
        const resp = await fetch(`${host}/group?groupId=${id}`, {
          credentials: "include",
        });
        const data = await resp.json();
        setgroupdata(data);
      } catch {}
    }
    getGroupData();
  }, [id, host]);

  console.log("----------------------------", groupData);
  if (!groupData) {
    return <div>Loading...</div>; // Show loading state while data is being fetched
  }

  const createEvent = () => {
    if (!eventTitle || !eventDatetime) return;
    setEvents([
      ...events,
      {
        title: eventTitle,
        description: eventDescription,
        datetime: eventDatetime,
        createdAt: new Date().toISOString(),
        responses: {},
      },
    ]);
    setEventTitle("");
    setEventDescription("");
    setEventDatetime("");
  };
  const respondToEvent = (index, response) => {
    const user = prompt("Entrez votre nom :");
    if (user) {
      const newEvents = [...events];
      newEvents[index].responses[user] = response;
      setEvents(newEvents);
    }
  };
  return (
    <div className={styles.parant}>
      <div className={styles.left}>
        <div className={styles.soutitre0}>
          <p>All users</p>
        </div>
        <div className={styles.chatbox0}></div>
      </div>
      <div className={styles.divcentral}>
        <div className={styles.supp}>
          <h1 className={styles.header}>{groupData.title}</h1>
          <p className={styles.description}>
            {" "}
            description: {groupData.description}
          </p>
          <small className={styles.creator}>
            Creator : {groupData.first_name} {groupData.last_name}
          </small>
        </div>
        <div className={styles.moyyen}>
          <div className={styles.postsContainer}>
            {posts.map((post) => (
              <div key={post.id} className={styles.postCard}>
                <div> {post.title}</div>
                <div> {post.content}</div>
                <div> {post.image}</div>
                <small className={styles.postDate}>{post.date}</small>
              </div>
            ))}
          </div>
          <div className={styles.creatPost}>
            <div className={styles.textareaWrapper}>
              <textarea
                className={styles.input}
                placeholder="Share something..."
              ></textarea>

              <label htmlFor="imageUpload" className={styles.imageIcon}>
                📷
              </label>
              <input
                type="file"
                accept="image/*"
                id="imageUpload"
                className={styles.hiddenInput}
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
                />
                <br />
                <br />
                <textarea
                  className={styles.input3}
                  value={eventDescription}
                  onChange={(e) => setEventDescription(e.target.value)}
                  placeholder="Description"
                  rows="2"
                ></textarea>
                <br />
                <br />
                <input
                  className={styles.input4}
                  type="datetime-local"
                  value={eventDatetime}
                  onChange={(e) => setEventDatetime(e.target.value)}
                />
                <br />
                <br />
                <button className={styles.button} onClick={createEvent}>
                  Create the Event
                </button>
              </div>
            </div>
          )}
          <button
            className={styles.addEventButton}
            onClick={() => setShowPopup(true)}
          >
            + Add Event
          </button>
        </div>
      </div>
      <div className={styles.right}>
        <GroupChat groupData={groupData}></GroupChat>
        <div className={styles.chatbox1}></div>
      </div>
    </div>
  );
}
