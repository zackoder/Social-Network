"use client"
import styles from "./id.module.css"

export default function GroupPage({ params }) {
  // params est injecté automatiquement en mode serveur dans /app
  console.log(params);


  return (
    <div className={styles.parant}>

         <div className={styles.left}>
      
         </div>
         <div className={styles.divcentral}>
              <div className={styles.supp}>

              </div>
              <div className={styles.moyyen}>

              </div>
              <div className={styles.infer}>
        
              </div>
      
           </div>
          <div className={styles.right}>
      
          </div>
   </div>
    
  );
}