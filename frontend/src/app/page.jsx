import Image from "next/image";
import styles from "./page.module.css";
import Post from "@/components/post/post";
import CreatePost from "@/components/createPost/createPost";
import ChatBox from "@/components/chatbox/chatbox.jsx";
import Contacts from "@/components/contacts/contacts.jsx"
import ChatSystem from "@/components/chatsystem/chatsystem";

export default function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.sidebar}>
        <ChatSystem />
      </div>
      
      <div className={styles.posts}>
        <CreatePost />
        <Post />
      </div>
    </div>
  );
}
