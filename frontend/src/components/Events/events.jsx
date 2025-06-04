"use client";



  const [events, setEvents] = useState([]);









export function Events (){










return <div className={styles.EventsCards}>
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
                {event.responce === "" && !answeredEvents.includes(event.id) && (
                  <div className={styles.buttonContainer}>
                    <button
                      className={styles.goenButton}
                      onClick={() => handleResponse("going", event.id)}
                    >
                      Goen
                    </button>
                    <button
                      className={styles.notGoenButton}
                      onClick={() => handleResponse("not going", event.id)}
                    >
                      Not Goen
                    </button>


                  </div>
                )}

              </div>
            ))}
          </div>


}