import ButtonFollow from "@/elements/buttonFollow/buttonFollow";
import styles from "./profile.module.css"
import Image from "next/image";
import Post from "@/components/post/post";

export default function page(props){
    let idUser = props.params.id
    //`/api/${post.poster}?id=${post.poster}`
    return(
        <div className={styles.container}>
            <header className={styles.header}>
                <div className={styles.info}>
                    <div className={styles.boxImage}>
                        <Image
                            className={styles.image}
                            src="/profile/profile.png"
                            alt="Image Profile"
                            fill={true}
                        />
                    </div>
                    <div className={styles.name}>
                        <h3>full Name</h3>
                        <p>description Lorem ipsum dolor sit amet consectetur adipisicing.</p>
                        <ButtonFollow />
                    </div>
                </div>
                <div className={styles.follow}>
                    <div className={styles.followers}>
                        <p>Followers</p>
                        <h3>2,985</h3>
                    </div>
                    <div className={styles.following}>
                        <p>Following</p>
                        <h3>132</h3>
                    </div>
                </div>
            </header>
            <main>
                <Post />
            </main>
        </div>
    );
}