"use client"
import { useState } from "react";
import styles from "./groups.module.css"; // ou ton fichier CSS module

export default function Home() {
  const [groups, setGroups] = useState([]);
  const [error, setError] = useState(null);

  const host = process.env.NEXT_PUBLIC_HOST;
  

  const fetchMyGroups = async () => {
    setError(null);
    try {
      const res = await fetch(`${host}/GetMyGroups`, {
        credentials:"include"
      });
      if (!res.ok) throw new Error("Erreur lors du fetch de mes groupes");
      const data = await res.json();
      console.log(data);
      
      setGroups(data);
    } catch (err) {
      setError(err.message);
      setGroups([]);
    }
  };

  const fetchJoinedGroups = async () => {
    setError(null);
    try {
      const res = await fetch(`${host}/GetJoinedGroups`,{
        credentials:"include"

      });
      if (!res.ok) throw new Error("Erreur lors du fetch des groupes rejoints");
      const data = await res.json();
      setGroups(data);
    } catch (err) {
      setError(err.message);
      setGroups([]);
    }
  };

  const fetchAllGroups = async () => {
    setError(null);
    try {
      const res = await fetch(`${host}/GetGroups`,{
        credentials:"include"
      });
      console.log(res);
      
      if (!res.ok) throw new Error("Erreur lors du fetch des groupes");
      const data = await res.json();
      console.log(data);
      
      setGroups(data);
    } catch (err) {
      setError(err.message);
      setGroups([]);
    }
  };

  return (
    <div>
      <div className={styles.div0}>
        <a href="#" className={styles.lien} onClick={fetchMyGroups}>
          My groups
        </a>
        <a href="#" className={styles.lien} onClick={fetchJoinedGroups}>
          Joined groups
        </a>
        <a href="#" className={styles.lien} onClick={fetchAllGroups}>
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
