"use client"
// import styles from ""
import { FaCloudUploadAlt } from "react-icons/fa";
import { LuSend } from "react-icons/lu";
import './likeDislikeComment.modules.css';
import { BiLike, BiDislike } from "react-icons/bi";
import React, { useEffect, useRef, useState } from 'react';
// import styles from './likeDislikeComment.modules.css';

export default function LikeDislikeComment({postId}) {
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);
  const [comment, setComment] = useState('');
  const [image, setImage] = useState("");
  const [submittedComment, setSubmittedComment] = useState('');
  const fileInputRef = useRef(null);
  const host = process.env.NEXT_PUBLIC_HOST


  const  handleLike =  async () => {
    try {
      const response = await fetch(`${host}/addReaction`,{
        method:"POST",
        credentials: "include",
        body: JSON.stringify({postId:postId,reactionType:"like"})
      }) 
      if (!response.ok) {
        throw new Error(error);
      } 
      const data = await response.json()
      // check status
      if (await data.message == "Reaction updated" && await data.reaction.reactionType == "like"){
          setLiked(true);
          setDisliked(false);
      } else if (data.message === 'Reaction removed') {
        setLiked(false);
        setDisliked(false);
      } else {
        setLiked(true);
        setDisliked(false);
      }
      // data.message 
      // umdate or remove

    }catch(error){
      console.log();
      
    }
  };
	// Id           int    `json:"id"`
	// PostId       int    `json:"postId"`
	// UserId       int    `json:"userId"`
	// ReactionType string `json:"reactionType"`
	// Date         int64  `json:"date"`
  const handleDislike =  async() => {
      try {
      const response = await fetch(`${host}/addReaction`,{
        method:"POST",
        credentials: "include",
        body: JSON.stringify({postId:postId,reactionType:"dislike"})
      }) 
      if (!response.ok) {
        throw new Error(error);
      } 
      const data = await response.json()
      // check status
      if (await data.message == "Reaction updated" && await data.reaction.reactionType == "dislike"){
          setLiked(false);
          setDisliked(true);
      } else if (data.message === 'Reaction removed') {
        setLiked(false);
        setDisliked(false);
      } else {
        setLiked(false);
        setDisliked(true);
      }
      // data.message 
      // umdate or remove

    }catch(error){
      console.log();
      
    }



    setDisliked(!disliked);
    if (liked) setLiked(false);
  };
  
    const getReactions = async (postId) => {
      try {
      const response = await fetch(`${host}/getReactions?postId=${postId}`, {
        credentials: "include",
      });
      const data = await response.json();
      const userReaction = data.userReaction;
      if (!userReaction || !userReaction.reactionType) {
        setLiked(false);
        setDisliked(false);
      } else if (userReaction.reactionType === "like") {
        setLiked(true);
        setDisliked(false);
      } else if (userReaction.reactionType === "dislike") {
        setLiked(false);
        setDisliked(true);
      }
    } catch (error) {
      console.error("Failed to fetch user reaction:", error);
    }
  };
  
    useEffect(()=> {
          getReactions(postId)
    }, [])
  
  const handleCommentSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    const commentData = {
      content : comment,
    };
    formData.append('commentData', JSON.stringify(commentData))
    if (image) {
      formData.append('avatar', image)
    }
    try {
      const response = fetch(`${host}/addComment`, {
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
