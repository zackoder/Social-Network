"use client"
// import Link from "next/link";
import styles from "./button.module.css";

export default function ButtonLogout(){
    return(
        <button 
            className={styles.logout}
            // href={"/"}
            onClick={()=>{console.log("Logout")}}    
        >
            <img
                src="/images/logout.png" 
                alt="Logout"
                title="Logout"
                style={{width: "40px", height: "40px"}}  //border: "2px solid var(--color-primary)", borderRadius: "40px"
            />
        </button>
    );
}