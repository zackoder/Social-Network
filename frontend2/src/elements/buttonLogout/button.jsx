"use client"
// import Link from "next/link";
import styles from "./button.module.css";

export default function ButtonLogout(){
    return(
        <button 
            className={styles.logout}
            // href={"/"}
            onClick={()=>{console.log("Logout")}}    
        >Logout</button>
    );
}