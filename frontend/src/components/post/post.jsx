 

"use client";

import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import styles from "./post.module.css";
import Link from "next/link";
import { useState, useEffect } from "react";
const host = process.env.NEXT_PUBLIC_HOST;

export default  function Post({ post = null, posts: propsPosts = null }) {
  const [posts, setPosts] = useState([]);
  useEffect(() => {
    // If an array of posts is passed as a prop
    if (propsPosts && Array.isArray(propsPosts)) {
      setPosts(propsPosts);
      return;
    }
 
    if (post) {
      setPosts([post]);
      return;
    }
 
        
      }, [posts, propsPosts]);
      
 
  return (
    <div className={styles.container}>
      {posts.map((post, index) => (
        <div className={styles.post} key={index}>
          <div className={styles.header}>
            <Link
              href={`/profile?id=${post.poster}&profile=${post.first_name}`}
            >
              <div className={styles.containerHeader}>
                <div className={styles.imageContainer}>
                  <img
                    src={`${host}${post.avatar}`}
                    width={50}
                    height={50}
                    style={{ borderRadius: "50%" }}
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
            {post.image && (
              <img
                className={styles.image}
                src={`${host}${post.image}`}
                alt="post"
                width={500}
                height={300}
             
              />
            )}
          </div>

          <div className={styles.reaction}>
            <LikeDislikeComment postId={post.id} />
          </div>
        </div>
      ))}
    </div>
  );

}
