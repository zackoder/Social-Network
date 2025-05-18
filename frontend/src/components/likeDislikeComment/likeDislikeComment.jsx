"use client"
// import styles from ""
import { BiLike, BiDislike } from "react-icons/bi";
import React, { useRef, useState } from 'react';
// import styles from './likeDislikeComment.modules.css';
import './likeDislikeComment.modules.css';
import { FaCloudUploadAlt } from "react-icons/fa";
import { LuSend } from "react-icons/lu";

export default function LikeDislikeComment() {
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);
  const [comment, setComment] = useState('');
  const [image, setImage] = useState("");
  const [submittedComment, setSubmittedComment] = useState('');
  const fileInputRef = useRef(null);
  const host = process.env.NEXT_PUBLIC_HOST


  const handleLike = () => {
    setLiked(!liked);
    if (disliked) setDisliked(false);
  };

  const handleDislike = () => {
    setDisliked(!disliked);
    if (liked) setLiked(false);
  };

  const handleCommentSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    const commentData = {
      liked,
      disliked,
      comment,
    };
    formData.append('postData', JSON.stringify(commentData))
    if (image) {
      formData.append('image', image)
    }
    try {
      const response = fetch(`${host}/end point the comments `, {
        method: "POST",
        body: formData,
        credentials: "include"
      });
      const comment = await response.json();
      if (!response.ok) {
        throw new Error(comment.error);
      } else {
        setSubmittedComment(comment);
        setComment('');
        setLiked(false)
        setDisliked(false)
        setImage("")
      }

      // Reset file input
      if (fileInputRef.current) {
        fileInputRef.current.value = ""
      }
    } catch (error) {
      console.log("Submission Error");
    }

    //end point
  };

  return (
    //className={styles.reactionContainer}
    <div style={{ minWidth: '100%', margin: '5px auto' }}>
      <div className="buttons" style={{ fontSize: '24px', marginBottom: '10px' }}>
        {/* className={styles.button} */}
        <button onClick={handleLike} style={{ color: liked ? 'var(--color-primary)' : 'gray', fontSize: '18px', backgroundColor: 'transparent', border: 'none' }}>
          <BiLike /> <span>Like</span>
        </button>
        <button
          // className={styles.button}
          onClick={handleDislike}
          style={{ color: disliked ? 'red' : 'gray', fontSize: '18px', backgroundColor: 'transparent', border: 'none' }}
        >
          <BiDislike /> <span>Dislike</span>
        </button>
      </div>

      <form className="formComment" onSubmit={handleCommentSubmit}>
        <input
          placeholder="Write a comment..."
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          style={{ width: '87%', padding: '14px 5px', borderRadius: '4px', resize: 'none', border: 'none', outline: 'none' }}
        />

        <input type="file" id="uploadImage" className="uploadImage" onChange={(e) => { setImage(e.target.files[0]) }} />
        <label htmlFor="uploadImage">
          <FaCloudUploadAlt className="uploadIcon" />
        </label>

        <input type="submit" className="submit" />
        <label htmlFor="submitComment">
          <LuSend className="submitIcon" />
        </label>
        {/* </button> */}
      </form>

      {submittedComment && (
        <div style={{ marginTop: '20px', background: '#222', padding: '10px', borderRadius: '4px' }}>
          <strong>Comment:</strong>
          <p>{submittedComment}</p>
        </div>
      )}
    </div>
  );
}
