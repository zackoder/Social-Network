"use client"
import { redirect } from "next/navigation";
// import Link from "next/link";
import styles from "./button.module.css";

export default function ButtonLogout(){
    const link = process.env.NEXT_PUBLIC_HOST
    return(
        <button 
            className={styles.logout}
            // href={"/"}
            onClick={()=>{
                fetch(`${link}/api/logout`, {
                    method: "POST",
                    credentials: "include"
                });
                localStorage.removeItem("user-id")
                redirect('/login')
            }}    
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