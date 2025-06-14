"use client";

import styles from "./navbar.module.css";
import Notification from "../notifications/notification";
import { FaUsers } from "react-icons/fa";
import ButtonLogout from "@/elements/buttonLogout/button";
import Logo from "@/elements/logo/logo";
import ButtonProfile from "@/elements/buttonProfile/buttonProfile";
import { useRouter } from "next/navigation";

export default function Navbar() {
  const router = useRouter();
  return (
    <div className={styles.container}>
      <div className={styles.nav_start}>
        <Logo />
      </div>
      <div className={styles.nav_middle}>
        <Notification />
      </div>
      <div className={styles.membersButtonWrapper}>
        <button
          onClick={() => router.push("/members")}
          className={styles.membersButton}
        >
          <FaUsers style={{ marginRight: "0.5rem" }} />
        </button>
      </div>
      <div className={styles.nav_end}>
        <ButtonProfile />
        <ButtonLogout />
      </div>
    </div>
  );
}
