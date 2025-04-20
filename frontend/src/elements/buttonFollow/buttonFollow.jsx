import styles from "./buttonFollow.module.css"
import { RiUserFollowFill } from "react-icons/ri";

export default function ButtonFollow(){
    return (
        <button className={styles.button}>
            <RiUserFollowFill /> Follow
        </button>
    );
}