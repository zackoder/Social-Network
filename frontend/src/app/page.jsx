import Image from "next/image";
import styles from "./page.module.css";
import Post from "@/components/post/post";
import CreatePost from "@/components/createPost/createPost";
import Contacts from "@/components/contacts/contacts";

export default function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.sidebar}>
        <Contacts />
      </div>
      
      <div className={styles.posts}>
        <CreatePost />
        <Post />
      </div>
    </div>
  );
}
