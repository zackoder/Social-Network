// import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
// import styles from "./post.module.css"
// // import Image from "next/image";
// import Link from "next/link";


// export function getData(data) {
//     console.log("data", data);
//     return data;

// }

// export async function fetchAllPosts() {
//     const host = process.env.NEXT_PUBLIC_HOST;
//     try{
//         const response = await fetch(`${host}/api/posts`);
//         if (!response.ok) {
//             throw new Error('Failed to Fetch Data');
//         }
//         const result = await response.json();
//         return result;
//     }catch (error){
//         console.log("Error loading posts:", error);

//     }
//     // const data = await response.json();
//     // getData(data)
// }



// export default function Post({ posts }) {
//     // const posts = await fetchAllPosts();
//     if (!posts || posts.lenght === 0) {
//         return <p>No posts yet.</p>;
//     }
//     // console.log("posts", posts);

//     return (
//         <div className={styles.container}>
//             {posts.map((post) => (
//                 <div className={styles.post} key={post.id}>
//                     <div className={styles.header}>
//                         <Link href={`/profile?id=${post.poster}&profile=${post.first_name}`}>
//                             <div className={styles.containerHeader}>
//                                 <div className={styles.imageContainer}>
//                                     {/* <img
//                                         className={styles.image}
//                                         src={`http://${post.avatar}`}
//                                         alt={`post.name`}
//                                         fill={false}
//                                     /> */}
//                                 </div>
//                                 <h2>{post.first_name}</h2>

//                             </div>
//                         </Link>
//                     </div>

//                     <div className={styles.content}>
//                         <h3>{post.title}</h3>
//                         <p>{post.content}</p>
//                     </div>
//                     <div className={styles.imagePost}>
//                         {post.image ? (
//                             <img
//                                 className={styles.image}
//                                 src={`http://${post.image}`}
//                                 alt="post"
//                                 width={500}
//                                 height={300}
//                             // fill={true}
//                             />

//                         ) : null}

//                     </div>

//                     <div className={styles.reaction}>
//                         <LikeDislikeComment postId={post.id} />
//                     </div>
//                 </div> //end post
//             ))}
//         </div> // end container
//     );
// }


"use client";

import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import styles from "./post.module.css";
import Link from "next/link";
import { useState, useEffect } from "react";

async function GetData() {
  const host = process.env.NEXT_PUBLIC_HOST;
  return fetch(`${host}/api/posts`).then((response) => {
    if (!response.ok) {
      console.log("Failed to Fetch Data");
      return [];
    }
    return response.json();
  });
}

export default function Post({ post = null, posts: propsPosts = null }) {
  const [posts, setPosts] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    // If an array of posts is passed as a prop
    if (propsPosts && Array.isArray(propsPosts)) {
      setPosts(propsPosts);
      return;
    }

    // If a single post is passed
    if (post) {
      setPosts([post]);
      return;
    }

    // If no props, fetch all posts
    setLoading(true);
    GetData()
      .then((data) => {
        setPosts(Array.isArray(data) ? data : []);
        setLoading(false);
      })
      .catch((err) => {
        setError(err.message);
        setLoading(false);
      });
  }, [post, propsPosts]);

  if (loading) return <div className={styles.container}>Loading posts...</div>;
  if (error) return <div className={styles.container}>Error: {error}</div>;

  return (
    <div className={styles.container}>
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
            {post.image && (
              <img
                className={styles.image}
                src={`http://${post.image}`}
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
