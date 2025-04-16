'use client'
import styles from "./LikeDislike.module.css"
import { BiLike, BiDislike } from "react-icons/bi";

import React, { useState } from 'react';

export default function LikeDislike() {
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);

  const handleLike = () => {
    setLiked(!liked);
    if (disliked) setDisliked(false); // undo dislike if like is clicked
  };

  const handleDislike = () => {
    setDisliked(!disliked);
    if (liked) setLiked(false); // undo like if dislike is clicked
  };

  return (
    // style={{ fontSize: '24px' }}
    <div className={styles.reactionContainer} style={{ fontSize: '24px' }}> 
      <div className={styles.like}>
        <button className={styles.button} onClick={handleLike} style={{ color: liked ? 'var(--color-primary)' : 'gray', fontSize: '18px' }}>
        <BiLike /> <span>Like</span>
          
          {/* <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M709.23-140H288.46v-480l265.39-263.84L587.69-850q6.23 6.23 10.35 16.5 4.11 10.27 4.11 19.35V-804l-42.46 184h268q28.54 0 50.42 21.89Q900-576.23 900-547.69v64.61q0 6.23-1.62 13.46-1.61 7.23-3.61 13.47L780.15-185.69q-8.61 19.23-28.84 32.46T709.23-140Zm-360.77-60h360.77q4.23 0 8.65-2.31 4.43-2.31 6.74-7.69L840-480v-67.69q0-5.39-3.46-8.85t-8.85-3.46H483.85L534-779.23 348.46-594.46V-200Zm0-394.46V-200v-394.46Zm-60-25.54v60H160v360h128.46v60H100v-480h188.46Z" /></svg>
          <span>Like</span> */}
        </button>
      </div>
      <div className={styles.dislike}>
        <button className={styles.button} onClick={handleDislike} style={{ color: disliked ? 'red' : 'gray', marginLeft: '10px', fontSize: '18px' }}>
          <BiDislike /> <span>Dislike</span>
          {/* <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M250.77-803.84h420.77v479.99L406.15-60l-33.84-33.85q-6.23-6.23-10.35-16.5-4.11-10.27-4.11-19.34v-10.16l42.46-184h-268q-28.54 0-50.42-21.88Q60-367.62 60-396.15v-64.62q0-6.23 1.62-13.46 1.61-7.23 3.61-13.46l114.62-270.46q8.61-19.23 28.84-32.46t42.08-13.23Zm360.77 60H250.77q-4.23 0-8.65 2.3-4.43 2.31-6.74 7.7L120-463.84v67.69q0 5.38 3.46 8.84 3.46 3.47 8.85 3.47h343.84L426-164.61l185.54-184.77v-394.46Zm0 394.46v-394.46 394.46Zm60 25.53v-59.99H800v-360H671.54v-60H860v479.99H671.54Z" /></svg>
          <span>Dislike</span> */}
        </button>
      </div>

    </div>
  );
}



{/* <div className={styles.dislike}>
  <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M250.77-803.84h420.77v479.99L406.15-60l-33.84-33.85q-6.23-6.23-10.35-16.5-4.11-10.27-4.11-19.34v-10.16l42.46-184h-268q-28.54 0-50.42-21.88Q60-367.62 60-396.15v-64.62q0-6.23 1.62-13.46 1.61-7.23 3.61-13.46l114.62-270.46q8.61-19.23 28.84-32.46t42.08-13.23Zm360.77 60H250.77q-4.23 0-8.65 2.3-4.43 2.31-6.74 7.7L120-463.84v67.69q0 5.38 3.46 8.84 3.46 3.47 8.85 3.47h343.84L426-164.61l185.54-184.77v-394.46Zm0 394.46v-394.46 394.46Zm60 25.53v-59.99H800v-360H671.54v-60H860v479.99H671.54Z" /></svg>
  <span>Dislike</span>
</div> */}


//--------


// return(
//     <div className={styles.like}>
//         <svg xmlns="http://www.w3.org/2000/svg" height="24px" viewBox="0 -960 960 960" width="24px" fill="#e3e3e3"><path d="M709.23-140H288.46v-480l265.39-263.84L587.69-850q6.23 6.23 10.35 16.5 4.11 10.27 4.11 19.35V-804l-42.46 184h268q28.54 0 50.42 21.89Q900-576.23 900-547.69v64.61q0 6.23-1.62 13.46-1.61 7.23-3.61 13.47L780.15-185.69q-8.61 19.23-28.84 32.46T709.23-140Zm-360.77-60h360.77q4.23 0 8.65-2.31 4.43-2.31 6.74-7.69L840-480v-67.69q0-5.39-3.46-8.85t-8.85-3.46H483.85L534-779.23 348.46-594.46V-200Zm0-394.46V-200v-394.46Zm-60-25.54v60H160v360h128.46v60H100v-480h188.46Z"/></svg>
//         <span>Like</span>
//     </div>
// );