"use client";
import styles from "./id.module.css";
import { useParams } from "next/navigation";
import { useSearchParams } from 'next/navigation';
import { useEffect, useState } from "react";


export default function GroupPage() {
  const params = useParams();  // récupère params dynamiques
  const { id } = params;       // ici id correspond au nom du fichier dynamique [id]
  console.log({ id });

  const host = process.env.NEXT_PUBLIC_HOST;

  // Pour récupérer les query params (comme title, description) passés dans l'URL,
  // useSearchParams est utile:
  const searchParams = useSearchParams();
  const [groupData, setgroupdata] = useState(null)
  useEffect(() => {
    async function getGroupData() {
      try {
        const resp = await fetch(`${host}/group?groupId=${id}`, {
          credentials: "include"
        })
        const data = await resp.json()
        setgroupdata(data)
      } catch {

      }
    }
    getGroupData()
  }, [id, host])

  console.log("----------------------------", groupData);
  if (!groupData) {
    return <div>Loading...</div>; // Show loading state while data is being fetched
  }
  return (
    <div className={styles.parant}>
      <div className={styles.left}></div>
      <div className={styles.divcentral}>
        <div className={styles.supp}>
          <h1>{groupData.title}</h1>
          <p>{groupData.description}</p>
          <small>Créateur : {groupData.first_name} {groupData.last_name}</small>
        </div>
        <div className={styles.moyyen}></div>
        <div className={styles.infer}></div>
      </div>
      <div className={styles.right}></div>
    </div>
  );
}
