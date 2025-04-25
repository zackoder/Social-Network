import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import styles from "./post.module.css"
// import Image from "next/image";
import Link from "next/link";

async function getData() {
    const host = process.env.NEXT_PUBLIC_HOST;
    const response = await fetch(`${host}/api/posts`);
    if (!response.ok) {
        throw new Error('Failed to Fetch Data');
    }
    return response.json();
}

export default async function Post() {
    const posts = await getData();
    console.log(posts);

    return (

        <div className={styles.container}>
            {posts.map((post) => (
                <div className={styles.post} key={post.id}>
                    <div className={styles.header}>
                        <Link href={"/profile?id=1&profile=zack"}>
                            <div className={styles.containerHeader}>
                                <div className={styles.imageContainer}>
                                    {/* <Image
                                className={styles.image}
                                src=""
                                alt=""
                                fill={false}
                            /> */}
                                </div>
                                <h2>Name</h2>

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