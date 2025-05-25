"use client"
import { useState , useEffect} from "react";
import styles from "./groups.module.css";

export default function Home() {
  const [groups, setGroups] = useState([]);
  const [error, setError] = useState(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  const host = process.env.NEXT_PUBLIC_HOST;
  

  const fetchGroups = async (url) => {

    setError(null);
    try {
      const res = await fetch(`${host}${url}`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Erreur lors du fetch des groupes");
      const data = await res.json();
      if (!data || data.length === 0) {
        setGroups([]);
        setError("No available Groups");
        return;
      }
      setGroups(data);
    } catch (err) {
      setGroups([]);
      setError(err.message || "Erreur inconnue");
    }
  };
    useEffect(() => {
    fetchGroups("/GetGroups");
  }, []);
   const handleCreateGroup = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`${host}/creategroup`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ title, description }),
      });
      if (!res.ok) throw new Error("Group creation failed");
      const newGroup = await res.json();
      setIsPopupOpen(false);
      setTitle("");
      setDescription("");
      setGroups((prev) => [...prev, newGroup]);

    } catch (err) {
      alert(err.message || "Error creating group");
    }
  };

  return (
    
    <div>
      <div className={styles.div0}>
        <a
          href="#"
          className={styles.lien}
          onClick={(e) => {
            e.preventDefault();
            fetchGroups("/GetMyGroups");
          }}
        >
          My groups
        </a>
        <a
          href="#"
          className={styles.lien}
          onClick={(e) => {
            e.preventDefault();
            fetchGroups("/GetJoinedGroups");
          }}
        >
          Joined groups
        </a>
        <a
          href="#"
          className={styles.lien}
          onClick={(e) => {
            e.preventDefault();
            fetchGroups("/GetGroups");
          }}
        >
          All groups
        </a>
        <div className={styles.div3}>
          <button
           className={styles.button} 
           onClick={() => {
             console.log("Popup ouverte !");
             setIsPopupOpen(true);
          }}>Creat groupe</button>
        </div>
      </div>
        {/* POPUP MODAL */}
      {isPopupOpen && (
  <div className={styles.modalOverlay} onClick={() => setIsPopupOpen(false)}>
    <div className={styles.modal} onClick={(e) => e.stopPropagation()}>
      <h2>Create a group</h2>
      <form onSubmit={handleCreateGroup}>
        <label>Title :</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          required
        />
        <label>Description :</label>
        <textarea
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          required
        ></textarea>
        <div className={styles.actions}>
          <button type="submit" className={styles.submitBtn}>Create</button>
          <button type="button" onClick={() => setIsPopupOpen(false)} className={styles.cancelBtn}>Cancel</button>
        </div>
      </form>
    </div>
  </div>
)}


      <div className={styles.contenu} style={{ marginTop: "20px" }}>
        {error && <p style={{ color: "red" }}>{error}</p>}

        {groups.length > 0 ? (
          <ul>
            {groups.map((groupe, i) => (
              <li key={i}>
                <p>{groupe.title}</p>
              </li>
            ))}
          </ul>
        ) : (
          !error && <p>Aucun groupe à afficher.</p>
        )}
      </div>
    </div>
  );
}
