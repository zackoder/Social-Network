"use client"
import { useState , useEffect} from "react";
import styles from "./groups.module.css";

export default function Home() {
  const [groups, setGroups] = useState([]);
  const [error, setError] = useState(null);

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
      </div>

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
