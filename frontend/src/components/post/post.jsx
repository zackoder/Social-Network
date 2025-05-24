"use client";

import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import styles from "./post.module.css";
import Link from "next/link";
import { useState, useEffect } from "react";

function getData() {
  const host = process.env.NEXT_PUBLIC_HOST;
  return fetch(`${host}/api/posts`).then((response) => {
    console.log("response ------", response.ok);
    if (!response.ok) {
      // throw new Error("Failed to Fetch Data");
    }
    return response.json();
  });
}

export default function Post({ post }) {
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
    getData()
      .then((data) => {
        console.log(data);
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
    <div className={styles.container}>
      {posts.map((post) => (
        <div className={styles.post} key={post.id}>
          <div className={styles.header}>
            <Link href={`/profile?id=${post.poster}&profile=${post.name}`}>
              <div className={styles.containerHeader}>
                <div className={styles.imageContainer}></div>
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
            <LikeDislikeComment postId={post.id}/>
          </div>
        </div>
      ))}
    </div>
  );
}
