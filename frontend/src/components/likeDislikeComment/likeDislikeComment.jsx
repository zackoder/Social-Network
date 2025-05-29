"use client";
// import styles from ""
import { FaCloudUploadAlt } from "react-icons/fa";
import { LuSend } from "react-icons/lu";
import "./likeDislikeComment.modules.css";
import { BiLike, BiDislike } from "react-icons/bi";
import React, { useEffect, useRef, useState } from "react";
import { isAuthenticated } from "@/app/page";
// import styles from './likeDislikeComment.modules.css';

export default function LikeDislikeComment({ postId }) {
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);
  const [likeNumber, setLikeNbr] = useState(0);
  const [disLikeNumber, setDisLikeNbr] = useState(0);
  const [comment, setComment] = useState("");
  const [comments, setComments] = useState([]);
  const [image, setImage] = useState("");
  const [showComments, setShowComments] = useState(false);

  // const [submittedComment, setSubmittedComment] = useState("");
  const fileInputRef = useRef(null);
  const host = process.env.NEXT_PUBLIC_HOST;

  const getReactions = async (postId) => {
    try {
      const response = await fetch(`${host}/getReactions?postId=${postId}`, {
        credentials: "include",
      });
      const data = await response.json();
      setDisLikeNbr(data.counts.dislike);
      setLikeNbr(data.counts.like);

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

  const handleLike = async () => {
    try {
      const response = await fetch(`${host}/addReaction`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({ postId: postId, reactionType: "like" }),
      });
      const data = await response.json();
      if (!response.ok) {
        // throw new Error(error);
        console.log(error);
        isAuthenticated(response.status, data.error);
      }
      // check status

      if (
        (await data.message) == "Reaction updated" &&
        (await data.reaction.reactionType) == "like"
      ) {
        setLiked(true);
        setDisliked(false);
        // setLikeNbr(likeNumber+1)
        // setDisLikeNbr(disLikeNumber-1)
      } else if (data.message === "Reaction removed") {
        // setLikeNbr(likeNumber-1)
        setLiked(false);
        setDisliked(false);
      } else {
        // setLikeNbr(likeNumber+1)
        setLiked(true);
        setDisliked(false);
      }
      // data.message
      // umdate or remove
    } catch (error) {
      console.log();
    }
    await getReactions(postId);

    setLiked(!liked);
    if (disliked) {
      setDisliked(false);
    }
  };
  // Id           int    `json:"id"`
  // PostId       int    `json:"postId"`
  // UserId       int    `json:"userId"`
  // ReactionType string `json:"reactionType"`
  // Date         int64  `json:"date"`
  const handleDislike = async () => {
    try {
      const response = await fetch(`${host}/addReaction`, {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({ postId: postId, reactionType: "dislike" }),
      });
      const data = await response.json();
      if (!response.ok) {
        isAuthenticated(response.status, data.error);
        return;
        // throw new Error(error);
      }

      // check status
      if (
        (await data.message) == "Reaction updated" &&
        (await data.reaction.reactionType) == "dislike"
      ) {
        setLiked(false);
        setDisliked(true);
      } else if (data.message === "Reaction removed") {
        setLiked(false);
        setDisliked(false);
      } else {
        setLiked(false);
        setDisliked(true);
      }
      // data.message
      // umdate or remove
    } catch (error) {
      console.log();
    }
    await getReactions(postId);

    setDisliked(!disliked);
    if (liked) {
      setLiked(false);
    }
  };

  useEffect(() => {
    getReactions(postId);
  }, []);

  /* handling show comments */
  const handleClick = async () => {
    if (showComments) {
      setShowComments(false);
      return;
    }
    // if (comments.length === 0) {

    try {
      const response = await fetch(`${host}/getComments?postId=${postId}`, {
        // add an end point
        method: "GET",
        credentials: "include",
      });

      console.log("fetching data response1111111", response);

      if (!response.ok) {
        console.log(`Error: ${response.status}`);
        isAuthenticated(response.status);
        return;
      }
      const data = await response.json();
      console.log("data the fetch posts------ ", data);
      if (data.content === ""){
        return
      }

      setComments(data);
      setShowComments(true);
      console.log("comments---", comments);

      console.log("fetch comments: ", data); // for testing fetching
    } catch (err) {
      console.log("error", err);
    // }
  }
  };

  const handleCommentSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    const commentData = {
      content: comment,
      postId: postId,
    };
    formData.append("commentData", JSON.stringify(commentData));
    if (image) {
      formData.append("avatar", image);
    }
    try {
      const response = await fetch(`${host}/addComment`, {
        method: "POST",
        body: formData,
        credentials: "include",
      });
      const comment = await response.json();
      //  const result = await response.json();
      if (!response.ok) {
        isAuthenticated(response.status, comment.error);
        return;
      }
      // setComments((prevComments) => [comment, ...prevComments]);
      setComments((prev) => [comment, ...prev]);
      setShowComments(true);
      setComment("");
      setImage("");

      // Reset file input
      if (fileInputRef.current) {
        fileInputRef.current.value = "";
      }
    } catch (error) {
      console.log("Submission Error", error);
    }
    //end point
  };
  return (
    //className={styles.reactionContainer}
    <div style={{ minWidth: "100%", margin: "5px auto" }}>
      <div
        className="buttons"
        style={{ fontSize: "24px", marginBottom: "10px" }}
      >
        {/* className={styles.button} */}
        <button
          onClick={handleLike}
          style={{
            color: liked ? "var(--color-primary)" : "gray",
            fontSize: "18px",
            backgroundColor: "transparent",
            border: "none",
          }}
        >
          <BiLike /> <span>Like {likeNumber}</span>
        </button>
        <button
          // className={styles.button}
          onClick={handleDislike}
          style={{
            color: disliked ? "red" : "gray",
            fontSize: "18px",
            backgroundColor: "transparent",
            border: "none",
          }}
        >
          <BiDislike /> <span>Dislike {disLikeNumber}</span>
        </button>
      </div>

      <form className="formComment" onSubmit={handleCommentSubmit}>
        <input
          aria-required
          placeholder="Write a comment..."
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          style={{
            width: "87%",
            padding: "14px 5px",
            borderRadius: "4px",
            resize: "none",
            border: "none",
            outline: "none",
          }}
        />

        <input
          type="file"
          id="uploadImage"
          className="uploadImage"
          onChange={(e) => {
            setImage(e.target.files[0]);
          }}
        />
        <label htmlFor="uploadImage">
          <FaCloudUploadAlt className="uploadIcon" />
        </label>

        <input type="submit" className="submit" id="submitComment" />
        <label htmlFor="submitComment">
          <LuSend className="submitIcon" />
        </label>
        {/* </button> */}
      </form>

      {/* this is button the show Comments and display comments*/}
      <button className="show" onClick={handleClick}>
        {showComments ? "Hide Comments" : "Show Comments"}
      </button>

      {showComments && (
        <div className="comments-container">
          {comments.length === 0 ? (
            <p style={{ color: "#aaa", textAlign: "center" }}>
              No comments yet.
            </p>
          ) : (
            comments.map((comment, index) => (
              <div className="comment" key={index}>
                <div className="comment-header">
                  <span className="comment-author">{comment.userName}</span>
                  <span className="comment-date">
                    {new Date(comment.date * 1000).toLocaleString()}
                  </span>
                </div>
                <div className="comment-content">
                  <p>{comment.content}</p>
                </div>
              </div>
            ))
          )}
        </div>
      )}
    </div>
  );
}
