"use client";

import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import style from "./Posts-Groups.module.css";
import Link from "next/link";
import { useState, useEffect, useRef } from "react";
import Modal from "../module/Modal";


const host = process.env.NEXT_PUBLIC_HOST;



export default function Post_Groups({ post, id }) {

  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [isModalOpen, setIsModalOpen] = useState(false);

  const [text, setText] = useState("");
  const [image, setImage] = useState(null);
  const [title, setTitle] = useState("")

  const fileInputRef = useRef(null);
  const handleSubmit = async (e) => {
    console.log(image);

    e.preventDefault();


    if (!text || !title) {
      alert("title and description are required")
      return

    };

    const formData = new FormData();
    const postData = {
      groupe_id: parseInt(id),
      title,
      content: text,
    }
    formData.append("postData", JSON.stringify(postData));
    if (image) {
      formData.append("avatar", image);
    }
    console.log([...formData.entries()]); // pour voir ce qu'il contient


    if (fileInputRef.current) {
      fileInputRef.current.value = "";
    }

    try {
      const response = await fetch(`${host}/addPost`, {
        credentials: "include",
        method: "POST",
        body: formData,
      });

      if (response.ok) {


        console.log("Post created successfully");
        setText("");
        setTitle("")
        setImage(null);
        setIsModalOpen(false)
        handlingdata()
      }
    } catch (err) {
      alert("Failed to create post", err.message);;
    }
  };


  const handleImageChange = (e) => {
    const file = e.target.files?.[0];
    if (file && file.type.startsWith("image/")) {
      setImage(file);
    }
  };
  function handlingdata() {
    GetData().then((data) => {
      setPosts(data);
      setLoading(false);
    })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      })
  }

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
      if (data && data.length > 0) {
        return data;
      } else {
        return [];
      }
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

      ;
  }, [post]);

  if (loading) return <div className={style.container}>Loading posts...</div>;
  if (error) return <div className={style.container}>Error: {error}</div>;

  return (<div>
    <div className={style.divclass}>
      {posts.map((post) => (



        <div className={style.poste} key={post.id}>
          <div className={style.header}>
            <Link href={`/profile?id=${post.poster}&profile=${post.first_name}`}>
              <div className={style.containerHeader} >
                <div className={style.imageContainer}>
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

          <div className={style.content}>
            <h3>{post.title}</h3>
            <p>{post.content}</p>
          </div>
          <div className={style.imagePost}>
            {post.image ? (
              <img
                className={style.imageG}
                src={`${host}${post.image}`}
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
    <div className={style.creatPost}>

      <div className={style.divbutton}>
        <button onClick={() => setIsModalOpen(true)} className={style.addEventButtone}>+Add Post</button>
      </div>
      <Modal isOpen={isModalOpen} onClose={() => setIsModalOpen(false)}>


        <h2 className={style.Createvent}>Create New Post</h2>
        <form onSubmit={handleSubmit}>
          <input
            className={style.input2}
            onChange={(e) => setTitle(e.target.value)}
            value={title}
            placeholder="Titre"
            type="text"
          />
          <textarea
            onChange={(e) => setText(e.target.value)}
            value={text}
            placeholder="Contenu"
            className={style.input3}
            rows={4}
          />
          <input
            type="file"
            className={style.input4}
            onChange={(e) => setImage(e.target.files[0])}
            ref={fileInputRef}
          />
          <button type="submit" className={style.button} >Publier</button>
        </form>

      </Modal>
    </div>
  </div>
  );
}