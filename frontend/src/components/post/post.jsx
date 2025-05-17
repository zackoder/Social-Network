import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import styles from "./post.module.css"
// import Image from "next/image";
import Link from "next/link";


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



export default function Post({posts}) {
    // const posts = await fetchAllPosts();
    if (!posts || posts.lenght === 0) {
        return <p>No posts yet.</p>;
    }

    return (
        <div className={styles.container}>
            {posts.map((post) => (
                <div className={styles.post} key={post.id}>
                    <div className={styles.header}>
                         <Link href={`/profile?id=${post.poster}&profile=${post.name}`}>
                            <div className={styles.containerHeader}>
                                <div className={styles.imageContainer}>
                                    {/* <Image
                                className={styles.image}
                                src=""
                                alt=""
                                fill={false}
                            /> */}
                                </div>
                                <h2>{post.name}</h2>

                            </div>
                        </Link>
                    </div>

                    <div className={styles.content}>
                        <h3>{post.title}</h3>
                        <p>{post.content}</p>
                    </div>
                    <div className={styles.imagePost}>
                        {post.image ? (
                            <Image
                                className={styles.image}
                                src={`http://${post.image}`}
                                alt="post"
                                width={500}
                                height={300}
                            // fill={true}
                            />

                        ) : null}

                    </div>

                    <div className={styles.reaction}>
                        <LikeDislikeComment />
                    </div>
                </div> //end post
            ))}
        </div> // end container
    );
}