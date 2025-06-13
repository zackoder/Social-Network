"use client";
import { useState, useEffect } from "react";
import styles from "./groups.module.css";
import { useRouter } from "next/router";
import GroupCard from "@/components/groupcard/groupCard";
import Link from "next/link";

export default function Home() {
  const [groups, setGroups] = useState([]);
  const [error, setError] = useState(null);
  const [isPopupOpen, setIsPopupOpen] = useState(false);
  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");
  const [activeTab, setActiveTab] = useState("all");

  const host = process.env.NEXT_PUBLIC_HOST;

  const fetchGroups = async (url) => {
    setError(null);
    try {
      const res = await fetch(`${host}${url}`, { credentials: "include" });
      if (!res.ok) throw new Error("An error occurred while fetching groups");
      const data = await res.json();
      // if (!data || data.length === 0) {
      //   setGroups([]);

      //   setError("No available Groups");
      //   return;
      // }
      // console.log(data);

      setGroups(Array.isArray(data) ? data : []);

      if (!data.length) setError("No available Groups");
    } catch (err) {
      setGroups([]);
      setError(err.message || "Unknown error");
    }
  };

  useEffect(() => {
    fetchGroups("/GetGroups");
  }, []);

  const handleTabClick = (type) => {
    setActiveTab(type);
    const endpoints = {
      all: "/GetGroups",
      my: "/GetMyGroups",
      joined: "/GetJoinedGroups",
    };
    fetchGroups(endpoints[type]);
  };

  const handleCreateGroup = async (e) => {
    e.preventDefault();
    try {
      const res = await fetch(`${host}/creategroup`, {
        method: "POST",
        credentials: "include",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title, description }),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error || "Unknown error");
      }

      setGroups((prev) => [...prev, data]);
      setIsPopupOpen(false);
      setTitle("");
      setDescription("");
      setError("");
    } catch (err) {
      alert(err.message || "Error creating group");
    }
  };

  return (
    <>
      <div>
        <div className={styles.div0}>
          <a
            href="#"
            className={`${styles.lien} ${
              activeTab === "my" ? styles.active : ""
            }`}
            onClick={(e) => {
              e.preventDefault();
              handleTabClick("my");
            }}
          >
            My groups
          </a>
          <a
            href="#"
            className={`${styles.lien} ${
              activeTab === "all" ? styles.active : ""
            }`}
            onClick={(e) => {
              e.preventDefault();
              handleTabClick("all");
            }}
          >
            All groups
          </a>
          <a
            href="#"
            className={`${styles.lien} ${
              activeTab === "joined" ? styles.active : ""
            }`}
            onClick={(e) => {
              e.preventDefault();
              handleTabClick("joined");
            }}
          >
            Joined groups
          </a>
          <div className={styles.div3}>
            <button
              className={styles.button}
              onClick={() => setIsPopupOpen(true)}
            >
              Create group
            </button>
          </div>
        </div>
        {isPopupOpen && (
          <div
            className={styles.modalOverlay}
            onClick={() => setIsPopupOpen(false)}
          >
            <div className={styles.modal} onClick={(e) => e.stopPropagation()}>
              <h2>Create a group</h2>
              <form onSubmit={handleCreateGroup}>
                <label>Title :</label>
                <input
                  type="text"
                  value={title}
                  onChange={(e) => setTitle(e.target.value)}
                  required
                />
                <label>Description :</label>
                <textarea
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  required
                />
                <div className={styles.actions}>
                  <button type="submit" className={styles.submitBtn}>
                    Create
                  </button>
                  <button
                    type="button"
                    className={styles.cancelBtn}
                    onClick={() => setIsPopupOpen(false)}
                  >
                    Cancel
                  </button>
                </div>
              </form>
            </div>
          </div>
        )}
        {error && <p style={{ color: "red" }}>{error}</p>}
        <GroupCard groups={groups} />;
        {/* <div className={styles.contenu} style={{ marginTop: "20px" }}>
        {error && <p style={{ color: "red" }}>{error}</p>}
        {groups.length > 0 ? (
          <ul>
            {groups.map((groupe, i) => (
              <li key={i}>
                <Link
                  href={{
                    pathname: `/groups/${groupe.Id}`,
                    query: {
                      Id: groupe.Id,
                     
                    },
                  }}
                >
                  <p>{groupe.title}</p>
                </Link>
              </li>
            ))}
          </ul>
        ) : (
          !error && <p>No groups to display.</p>
        )}
      </div> */}
      </div>
    </>
  );
}
