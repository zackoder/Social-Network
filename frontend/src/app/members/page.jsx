"use client";

import styles from "./page.module.css";
import { useEffect, useState } from "react";
import ButtonFollow from "@/elements/buttonFollow/buttonFollow";
import Link from "next/link";

export default function Members() {
  const host = process.env.NEXT_PUBLIC_HOST;
  const [allUsers, setAllUsers] = useState([]);
  const [loading, setLoading] = useState(true);

  const fetchAllUsers = async () => {
    try {
      const response = await fetch(`${host}/api/getallusers`, {
        method: "GET",
        credentials: "include",
      });
      const data = await response.json();

      setAllUsers(Array.isArray(data) ? data : []);
    } catch (error) {
      console.error("Error fetching users:", error);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchAllUsers();
  }, []);

  if (loading) {
    return <h2>Loading members...</h2>;
  }

  if (allUsers.length === 0) {
    return <h2>No members found.</h2>;
  }

  return (
    <div className={styles.container}>
      <h1 className={styles.title}>All Members</h1>
      <ul className={styles.userList}>
        {allUsers.map((user) => (
          <li key={user.ID} className={styles.userItem}>
            <Link href={`/profile?id=${user.ID}&profile=${user.firstname}`}>
              <div className={styles.userInfo}>
                <img
                  src={`${host}${user.avatar}`}
                  alt={`${user.firstname} ${user.lastname}`}
                  className={styles.avatar}
                />
                <div>
                  <div className={styles.name}>
                    {user.firstname} {user.lastname}
                  </div>
                </div>
              </div>
            </Link>
            <ButtonFollow profileId={user.ID} />
          </li>
        ))}
      </ul>
    </div>
  );
}
