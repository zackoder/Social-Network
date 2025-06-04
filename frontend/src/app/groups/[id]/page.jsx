"use client";
import styles from "./id.module.css";
import Modal from "@/components/module/Modal";
import Post_Groups from "@/components/Posts-Groups/postGroups";
import { useParams, useSearchParams } from "next/navigation";
import { useEffect, useState, useRef } from "react";
import GroupChat from "@/components/groupchat/groupchat";

export default function GroupPage() {
  const [text, setText] = useState("");
  const [image, setImage] = useState(null);
  const [showPopup, setShowPopup] = useState(false);
  const [title, setTitle] = useState("")
  const [isModalOpen, setIsModalOpen] = useState(false);

  const [eventTitle, setEventTitle] = useState('');
  const [eventDescription, setEventDescription] = useState('');
  const [eventDatetime, setEventDatetime] = useState('');
  const fileInputRef = useRef(null);
  const [events, setEvents] = useState([]);
  const [posts, setPosts] = useState([]);
  const [groupData, setGroupData] = useState(null);

  const params = useParams();
  const { id } = params;
  const searchParams = useSearchParams();
  const host = process.env.NEXT_PUBLIC_HOST;




const handleResponse = async (responseValue, eventId) => {
    try {
      const res = await fetch(`${host}/api/event-response`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          event_id: eventId,
          responce: responseValue,
        }),
      });

      if (!res.ok) {
        throw new Error('Erreur lors de la requête');
      }

      const data = await res.json();
      console.log('Réponse backend:', data);

      setAnsweredEvents(prev => [...prev, eventId]);

    } catch (error) {
      console.error('Erreur:', error);
    }
  };





  async function GetEvents() {
    try {
      const response = await fetch(`${host}/api/getevents`, {
        credentials: "include",
        method: "POST",
        body: parseInt(id)
      });

      if (response.ok) {
        const data = await response.json();
        console.log("---------------------------------------------------------------------", data);
        console.log(convertEventTimes(data));

        setEvents((data));
      }

    } catch (err) {
      console.error("Erreur de requête:", err);
      alert("Error: " + err.message);
    }
  }


  const handleImageChange = (e) => {
    const file = e.target.files?.[0];
    if (file && file.type.startsWith("image/")) {
      setImage(file);
    }
  };
  function convertEventTimes(events) {
    return events.map(event => {
      return {
        ...event,
        event_time: new Date(event.event_time).toString()

      };
    });
  }

  const handleSubmit = async (e) => {
    e.preventDefault();


    if (!text && !image) return;

    const formData = new FormData();
    const postData = {
      groupe_id: parseInt(id),
      title,
      content: text,
    }
    formData.append("postData", JSON.stringify(postData));
    formData.append("postData", text);
    if (image) {
      formData.append("image", image);
    }

    if (fileInputRef.current) {
      fileInputRef.current.value = "";
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
        setTitle("")
        setImage(null);
        setIsModalOpen(false)

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
    GetEvents();

    getGroupData();
  }, [id, host]);

  const createEvent = async () => {
    console.log("hyhy");

    if (!eventTitle || !eventDatetime || !eventDescription) return;
    const eventT = new Date(eventDatetime)
    console.log("___________________________________----------------", eventT.getTime());
    const response = await fetch(`${host}/CreatEvent`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        groupe_id: parseInt(id),
        title: eventTitle,
        description: eventDescription,
        event_time: (eventT.getTime() / 1000)
      })
    });
    console.log(response.ok);

    if (!response.ok) {

      setShowPopup(false);

      alert("failed to creat event")
      return

    }
    // data = awai response.JSON()
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
          {/* <div className={styles.postsContainer}> */}
          <Post_Groups />



          <div className={styles.creatPost}>
            <button onClick={() => setIsModalOpen(true)} className={styles.addEventButtone}>+Add Post</button>

            <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>


              <h2 className={styles.Createvent}>Create New Post</h2>
              <form onSubmit={handleSubmit}>
                <input
                  className={styles.input2}
                  onChange={(e) => setTitle(e.target.value)}
                  value={title}
                  placeholder="Titre"
                  type="text"
                />
                <textarea
                  onChange={(e) => setText(e.target.value)}
                  value={text}
                  placeholder="Contenu"
                  className={styles.input3}
                  rows={4}
                />
                <input
                  type="file"
                  className={styles.input4}
                  onChange={(e) => setImage(e.target.files[0])}
                  ref={fileInputRef}
                />
                <button type="submit" className={styles.button} >Publier</button>
              </form>

            </Modal>

          </div>
        </div>

        <div className={styles.infer}>
          <div className={styles.EventsCards}>
            {events.map((event) => (



              <div className={styles.event} key={event.id}>

                <div className={styles.Title}>
                  {event.title}
                </div>
                <div className={styles.description} >
                  {event.description}
                </div>
                <div className={styles.event_time}>
                  {new Date(event.event_time).toString()}

                </div>
                {event.responce === "" && (
                  <div className={styles.buttonContainer}>
                    <button
                      className={styles.goenButton}
                      onClick={() => handleResponse("Goen", event.id)}
                    >
                      Goen
                    </button>
                    <button
                      className={styles.notGoenButton}
                      onClick={() => handleResponse("Not Goen", event.id)}
                    >
                      Not Goen
                    </button>


                  </div>
                )}

              </div>
            ))}
          </div>

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
        <GroupChat groupData={groupData}></GroupChat>
        {/* <GroupChat groupData={groupData}></GroupChat> */}
        {/* <div className={styles.soutitre}><p>Group chat</p></div> */}
        <div className={styles.chatbox1}></div>
      </div>
    </div>
  );
}
