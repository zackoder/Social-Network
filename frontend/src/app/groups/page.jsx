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
        <a className={styles.lien} href="#" onClick={() => changeContent("Contenu du premier lien")}>My groups</a>
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
