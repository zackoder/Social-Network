"use client";
import styles from "./id.module.css";
import Post_Groups from "@/components/Posts-Groups/postGroups";
import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import { Events } from "@/components/Events/events";
import GroupChat from "@/components/groupchat/groupchat";
import ContactsPrivate from "@/components/contactprivate/contactprivate";

export default function GroupPage() {

  const [groupData, setGroupData] = useState(null);
  const params = useParams();
  const { id } = params;
  const host = process.env.NEXT_PUBLIC_HOST;





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



  useEffect(() => {


    getGroupData();
  }, [id, host]);






  if (!groupData) {
    return <div>Loading...</div>;
  }

  return (
    <div className={styles.parant}>
      <div className={styles.left}>
        <div className={styles.soutitre0}>
          <p>All users</p>
          <ContactsPrivate/>
        </div>
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
          <Post_Groups id={id} />
        </div>

        <div className={styles.infer}>
          <Events id={id} />

        </div>
      </div>

      <div className={styles.right}>
        <GroupChat groupData={groupData}></GroupChat>

        <div className={styles.chatbox1}></div>
      </div>
    </div>
  );
}
