import styles from "./page.module.css";
import ChatSystem from "@/components/chatsystem/chatsystem";

import PostSystem from "@/components/postsystem/postsystem";

export default function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.sidebar}>
        <ChatSystem />
      </div>

      <div className={styles.posts}>
        <PostSystem />
        {/* <CreatePost />
        <Post /> */}
      </div>
    </div>
  );
}
