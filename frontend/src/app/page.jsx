'use server'

import styles from "./page.module.css";
import ChatSystem from "@/components/chatsystem/chatsystem";

import PostSystem from "@/components/postsystem/postsystem";
import { redirect } from "next/navigation";
import { revalidatePath } from "next/cache";

export async function isAuthenticated(codeError, msgError) {
  console.log(codeError);

  if (codeError === 401) {
    // revalidatePath('/')
    redirect("/login")
  }
  if (codeError === 500){
    redirect("/error")
  }
}

export default async function Home() {
  return (
    <div className={styles.container}>
      <div className={styles.sidebar}>
        <ChatSystem />
      </div>

      <div className={styles.posts}>
        <PostSystem />
      </div>
    </div>
  );
}
