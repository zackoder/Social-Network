"use client";
import styles from "./id.module.css";
import Post_Groups from "@/components/Posts-Groups/postGroups";
import { useEffect, useState } from "react";
import { useParams } from "next/navigation";
import { Events } from "@/components/Events/events";
import GroupChat from "@/components/groupchat/groupchat";
import InviteUsers from "@/components/inviteFollower/inviteFollowers";
export default function GroupPage() {

  const [groupData, setGroupData] = useState(null);
  const [error, setError] = useState("");

  const params = useParams();
  const { id } = params;
  const host = process.env.NEXT_PUBLIC_HOST;





  async function getGroupData() {
    try {
      console.log(id);

      const resp = await fetch(`${host}/group?groupId=${id}`, {
        credentials: "include",
      });

      const data = await resp.json();

      if (!resp.ok) {
        throw new Error(data.error || "Unknown error");
      }

      setGroupData(data);
    } catch (err) {
      console.error("Failed to fetch group data:", err.message);
      setGroupData(null); // facultatif si tu veux afficher une erreur
      setError(err.message || "This group does not exist.");
    }
  }



  useEffect(() => {


    getGroupData();
  }, [id, host]);



  if (error) {
    return <div className={styles.error}>{error}</div>;
  }




  if (!groupData) {
    return <div>Loading...</div>;
  }

  return (

    <div className={styles.parant}>
      <div className={styles.left}>

        <InviteUsers group_id={id}></InviteUsers>


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
          <div className={styles.PostsTitle}>
            <header className={styles.textTitle}>Posts</header>

          </div>
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
