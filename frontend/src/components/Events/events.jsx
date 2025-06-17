"use client";
import { useEffect, useState } from "react";
import styles from "./Events.module.css";

const host = process.env.NEXT_PUBLIC_HOST;
const isDateInFuture = (selectedDate) => {
  const now = new Date();
  return selectedDate > now;
};

export function Events({ id }) {
  const [events, setEvents] = useState([]);
  const [answeredEvents, setAnsweredEvents] = useState([]);
  const [showPopup, setShowPopup] = useState(false);
  const [eventTitle, setEventTitle] = useState("");
  const [eventDescription, setEventDescription] = useState("");
  const [eventDatetime, setEventDatetime] = useState("");
  const [participationStatus, setParticipationStatus] = useState(null);


  const createEvent = async () => {
    if (!eventTitle || !eventDatetime || !eventDescription) return;
    const eventT = new Date(eventDatetime);

    if (isNaN(eventT)) {
      alert("Please enter a valid date!");
      return;
    }
    if (!isDateInFuture(eventT)) {
      alert("Please choose a date in the future!");
      return;
    }

    console.log(
      "___________________________________----------------",
      eventT.getTime()
    );
    const response = await fetch(`${host}/CreatEvent`, {
      method: "POST",
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        groupe_id: parseInt(id),
        title: eventTitle,
        description: eventDescription,
        event_time: eventT.getTime() / 1000,
        action :participationStatus
      }),
    });
    console.log(response.ok);

    if (!response.ok) {
      setShowPopup(false);

      alert("failed to creat event");
      return;
    }
    // data = awai response.JSON()
    setEventTitle("");
    setEventDescription("");
    setEventDatetime("");
    setShowPopup(false);
    GetEvents();
  };

  const handleResponse = async (responseValue, eventId) => {
    try {
      const res = await fetch(`${host}/api/event-response`, {
        credentials: "include",
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          groupe_id: parseInt(id),
          event_id: eventId,
          responce: responseValue,
        }),
      });

      if (!res.ok) {
        throw new Error("Erreur lors de la requête");
      }

      const data = await res.json();

      setAnsweredEvents((prev) => [...prev, eventId]);
    } catch (error) {
      console.error("Erreur:", error);
    }
  };

  async function GetEvents() {
    try {
      const response = await fetch(`${host}/api/getevents`, {
        credentials: "include",
        method: "POST",
        body: parseInt(id),
      });

      if (response.ok) {
        const data = await response.json();
        setEvents(Array.isArray(data) ? data : []);
      }
    } catch (err) {
      console.log(err);
      console.error("Erreur de requête:", err);
    }
  }

  useEffect(() => {
    GetEvents();
  }, []);

  return (
    <div>
      <div className={styles.EventTitle}>
        <header className={styles.TextHedear}>Events</header>
      </div>
      <div className={styles.EventsCards}>
        {events.map((event) => (
          <div className={styles.event} key={event.id}>
            <div className={styles.Title}>{event.title}</div>
            <div className={styles.description}>{event.description}</div>
            <div className={styles.event_time}>
              {new Date(event.event_time * 1000).toISOString()}
            </div>
            {event.responce === "" && !answeredEvents.includes(event.id) && (
              <div className={styles.buttonContainer}>
                <button
                  className={styles.goenButton}
                  onClick={() => handleResponse("going", event.id)}
                >
                  Going
                </button>
                <button
                  className={styles.notGoenButton}
                  onClick={() => handleResponse("not going", event.id)}
                >
                  Not Going
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
              X
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
              rows={2}
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
            <div className={styles.action}>
              <button
                className={`${styles.goenButton} ${participationStatus === 'going' ? styles.selected : ''}`}
                onClick={() => setParticipationStatus('going')}

              >
                Going
              </button>
              <button
               className={`${styles.notGoenButton} ${participationStatus === 'not going' ? styles.selected : ''}`}
               onClick={() => setParticipationStatus('not going')}
                
              >
                Not Going
              </button>

            </div>

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
        +Add Event
      </button>

    
    </div>
  );
}
