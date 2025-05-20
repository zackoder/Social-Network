"use client";
import { useState } from 'react';
import styles from "./groups.module.css"




export default function Home() {
  const [content, setContent] = useState("C'est le contenu par défaut");

  const changeContent = (newContent) => {
    setContent(newContent);
  };

  return (
    <div>
      {/* Conteneur pour les liens */}
      <div className={styles.div0}>
        <a className={styles.lien} href="#" onClick={() =>
          fetchGroups()

        }>My groups</a>
        <a className={styles.lien} href="#" onClick={() => changeContent("Contenu du deuxième lien")} >Joined groups</a>
        <a className={styles.lien} href="#" onClick={() => changeContent("Contenu du troisième lien")}>All the groups</a>
      </div>

      {/* Contenu changeable affiché sous les liens */}
      <div style={{ marginTop: '20px' }}>
        <h2>{content}</h2>
      </div>
    </div>
  );
}
async function fetchGroups() {
  const host = process.env.NEXT_PUBLIC_HOST;

  try {
    const response = await fetch(`${host}/GetGroups`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const groups = await response.json();
    console.log(groups);

  } catch (error) {
    console.log("Erreur lors du fetch des groupes :", error);
  }
}

