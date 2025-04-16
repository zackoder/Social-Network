import LikeDislikeComment from "../likeDislikeComment/likeDislikeComment";
import LikeDislike from "../likeDislike/LikeDislike";
import styles from "./post.module.css"
import Image from "next/image";

export default function Post(){
    return (
        <div className={styles.container}>
            <div className={styles.header}>
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
            <div className={styles.content}>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Totam minima ullam mollitia nesciunt? Cum quia dolorum corrupti ea, magnam voluptas.</p>
            </div>
            <div className={styles.imagePost}>
                <Image
                    className={styles.image}
                    src="/images/post.png"
                    alt="post"
                    // width={500}
                    // height={500}
                    fill={true}
                />
            </div>
            <div className={styles.reaction}>
                <LikeDislikeComment />
            </div>

        </div>
    );
}