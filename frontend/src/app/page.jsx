import Image from "next/image";
import styles from "./page.module.css";

export default function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.sidebar}>
        <div>profile</div>
        <div>contacts</div>
      </div>
      <div className={styles.posts}>posts</div>
    </div>
  );
}
