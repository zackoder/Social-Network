

"use client";

import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import styles from "./post.module.css";
import Link from "next/link";
import { useState, useEffect } from "react";

async function GetData() {
  console.log("rani dakgl hna ");
  
  const host = process.env.NEXT_PUBLIC_HOST;
  return fetch(`${host}/api/posts`).then((response) => {
    if (!response.ok) {
          return []; // retourne un tableau vide au lieu de undefined
      console.log("Failed to Fetch Data");
      return
    }
    console.log("anan");
    
    console.log(response.json());
    
    return response.json();
  });
}

export default function Post({ post, divclass = "container" }) {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  // If a single post is passed as prop, use it instead of fetching
  useEffect(() => {
    if (post) {
      setPosts([post]);
      return;
    }

    setLoading(true);
    GetData()
      .then((data) => {
        setPosts(data);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [post]);

  if (loading) return <div className={styles.container}>Loading posts...</div>;
  if (error) return <div className={styles.container}>Error: {error}</div>;

  return (
    <div className={divclass}>
      {posts.map((post) => (
        <div className={styles.post} key={post.id}>
          <div className={styles.header}>
            <Link href={`/profile?id=${post.poster}&profile=${post.first_name}`}>
              <div className={styles.containerHeader}>
                <div className={styles.imageContainer}>
                  <img
                    src={`http://${post.avatar}`}
                    width={50}
                    height={50}
                    style={{ borderRadius: "100%" }}

                  />
                </div>
                <h2>{post.first_name}</h2>
              </div>
            </Link>
          </div>

          <div className={styles.content}>
            <h3>{post.title}</h3>
            <p>{post.content}</p>
          </div>
          <div className={styles.imagePost}>
            {post.image ? (
              <img
                className={styles.image}
                src={`http://${post.image}`}
                alt="post"
                width={500}
                height={300}
              />
            ) : null}
          </div>

          <div className={styles.reaction}>
            <LikeDislikeComment postId={post.id} />
          </div>
        </div>
      ))}
    </div>
  );
}