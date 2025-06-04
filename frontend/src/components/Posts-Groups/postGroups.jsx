"use client";

import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import Post from "../post/post";
import style from "./Posts-Groups.module.css";
import Link from "next/link";
import { useState, useEffect } from "react";




export default function Post_Groups({ post,id }) {
  
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  
  
  async function GetData() {


  const host = process.env.NEXT_PUBLIC_HOST;

  try {
  const response = await fetch(`${host}/api/postsGroups`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ id: parseInt(id) }),
  
    });

    if (!response.ok) {
      console.error("faild to fetch");
      return [];
    }

    const data = await response.json();
    console.log("---------------------datadata-----------------------------------",data);
    return data;
  } catch (error) {
    console.error("Error", error);
    return [];
  }
}

  useEffect(() => {
    if (post) {
      setPosts([post]);
      return;
    }

    setLoading(true);
    GetData().then((data) => {
      setPosts(data);
      setLoading(false);
    })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [post]);

  if (loading) return <div className={style.container}>Loading posts...</div>;
  if (error) return <div className={style.container}>Error: {error}</div>;

  return (
    <div className={style.divclass}>
      {posts.map((post) => (



        <div className={style.poste} key={post.id}>
          <div className={style.header}>
            <Link href={`/profile?id=${post.poster}&profile=${post.first_name}`}>
              <div className={style.containerHeader} >
                <div className={style.imageContainer}>
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

          <div className={style.content}>
            <h3>{post.title}</h3>
            <p>{post.content}</p>
          </div>
          <div className={style.imagePost}>
            {post.image ? (
              <img
                className={style.image}
                src={`http://${post.image}`}
                alt="post"
                width={500}
                height={300}
              />
            ) : null}
          </div>

          <div className={style.reaction}>
            <LikeDislikeComment postId={post.id} />
          </div>
        </div>
      ))}
    </div>
  );
}