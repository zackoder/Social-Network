"use client";

import styles from "./navbar.module.css";
import Notification from "../notifications/notification";
// import { FaUsers } from "react-icons/fa";
import ButtonLogout from "@/elements/buttonLogout/button";
import Logo from "@/elements/logo/logo";
import ButtonProfile from "@/elements/buttonProfile/buttonProfile";
import { useRouter } from "next/navigation";
import { IoMdContacts } from "react-icons/io";

 

export default function Navbar() {
  const router = useRouter();
  return (
    <div className={styles.container}>
      <div className={styles.nav_start}>
        <Logo />
      </div>
      <div className={styles.nav_middle}>
        <Notification />
        <div className={styles.membersButtonWrapper}>
          <button
            onClick={() => router.push("/members")}
            className={styles.membersButton}
          >
            <IoMdContacts/>
          </button>
        </div>
      </div>
      <div className={styles.nav_end}>
        <ButtonProfile />
        <ButtonLogout />
      </div>
    </div>
  );
}
