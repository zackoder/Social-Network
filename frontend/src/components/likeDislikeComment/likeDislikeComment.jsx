"use client";
// import styles from ""
import { FaCloudUploadAlt } from "react-icons/fa";
import { LuSend } from "react-icons/lu";
import "./likeDislikeComment.modules.css";
import { BiLike, BiDislike } from "react-icons/bi";
import React, { useEffect, useRef, useState, useCallback } from "react";
import { isAuthenticated } from "@/app/page";
import { debounce } from "@/utils/debounce";
// import styles from './likeDislikeComment.modules.css';
const host = process.env.NEXT_PUBLIC_HOST;
const LIMIT = 10;

export default function LikeDislikeComment({ postId }) {
  const [offset, setOffset] = useState(0);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);
  const [likeNumber, setLikeNbr] = useState(0);
  const [disLikeNumber, setDisLikeNbr] = useState(0);
  const [comment, setComment] = useState("");
  const [comments, setComments] = useState([]);
  const [image, setImage] = useState(null);
  const [showComments, setShowComments] = useState(false);
  const fileInputRef = useRef(null);

  // *********************** comments logic ***********************//
  const fetchComments = async () => {
    if (loading || !hasMore) return;

    setLoading(true);
    try {
      const res = await fetch(
        `${host}/getComments?postId=${postId}&offset=${offset}&limit=${LIMIT}`,
        { method: "GET", credentials: "include" }
      );

      if (!res.ok) {
        isAuthenticated(res.status);
        return;
      }

      const data = await res.json();
        
      if (!data || data.length === 0) {
        setHasMore(false);
        return;
      }

      setComments((prev) => [...prev, ...data]);
      setOffset((prev) => prev + LIMIT);
    } catch (err) {
      console.error("Fetching comments error:", err);
    } finally {
      setLoading(false);
      setShowComments(true);
    }
  };

  const handleToggleComments = async () => {
    if (showComments) {
      // Hide and reset
      setShowComments(false);
      setComments([]);
      setOffset(0);
      setHasMore(true);
    } else {
      setShowComments(true);
      await fetchComments();
    }
  };
  const debouncedFetchComments = useCallback(debounce(fetchComments, 300), [
    offset,
    hasMore,
    loading,
  ]);
  const handleCommentSubmit = async (e) => {
    e.preventDefault();
    const formData = new FormData();
    const commentData = {
      content: comment,
      postId: postId,
    };
    formData.append("commentData", JSON.stringify(commentData));
      // if ()
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
      if (comments.length === 0) {
        await fetchComments();
      } else {
        setComments((prev) => [comment,...prev]);
      }
      // setComments((prevComments) => [comment, ...prevComments]);
      setShowComments(true);
      setComment("");
      setImage(null);

      // Reset file input
      if (fileInputRef.current) {
        fileInputRef.current.value = "";
      }
    } catch (error) {
      console.log("Submission Error", error);
    }
  };

  // *****************reaction logic ****************** //

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

  useEffect(() => {
    getReactions(postId);
  }, []);

  const handleReaction = async (reactionType) => {
    try {
      const response = await fetch(`${host}/addReaction`, {
        method: "POST",
        credentials: "include",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ postId, reactionType }),
      });

      const data = await response.json();

      if (!response.ok) {
        isAuthenticated(response.status, data.error);
        return;
      }

      const isLike = reactionType === "like";
      const isDislike = reactionType === "dislike";
      console.log(data);

      if ((await data.message) === "Reaction updated") {
        if ((await data.type) === "like") {
          setLiked(true);
          setDisliked(false);
        } else if (data.type === "dislike") {
          setLiked(false);
          setDisliked(true);
        }
      } else if (data.message === "Reaction removed") {
        setLiked(false);
        setDisliked(false);
      } else {
        // Default: new reaction
        setLiked(isLike);
        setDisliked(isDislike);
      }
    } catch (error) {
      console.error("Reaction error:", error);
    }

    await getReactions(postId);
  };

  /* handling image comments */
  const handleImageChange = (e) => {

    const file = e.target.files[0];
    if (file) {
      setImage(file);
    }
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
          onClick={() => handleReaction("like")}
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
          onClick={() => handleReaction("dislike")}
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
          // aria-required
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
          id={`uploadImage${postId}`}
          className="uploadImage"
          onChange={handleImageChange}
          ref={fileInputRef}
        />
        <label htmlFor={`uploadImage${postId}`}>
          <FaCloudUploadAlt className="uploadIcon" />
        </label>

        <input type="submit" className="submit" id="submitComment" />
        <label htmlFor="submitComment">
          <LuSend className="submitIcon" />
        </label>
        {/* </button> */}
      </form>

      {/* this is button the show Comments and display comments*/}

      <button className="show" onClick={handleToggleComments}>
        {showComments ? "Hide Comments" : "Show Comments"}
      </button>

      {showComments && (
        <div className="comments-container">
          {comments.length === 0 && !loading ? (
            <p className="no-comments">No comments yet.</p>
          ) : (
            <>
              {comments.map((c, i) => (
                <div className="comment" key={i}>
                  <div className="comment-header">
                    <div className="comment-image">
                      <img src={`${host}${c.userAvatar}`} alt="profile" />
                    </div>
                    <span className="comment-author">{c.userName}</span>
                    <span className="comment-date">
                      {new Date(c.date * 1000).toLocaleString()}
                    </span>
                  </div>
                  <div className="comment-content">
                    <p>{c.content}</p> {c.imagePath && 
                      <img src={`${host}${c.imagePath}`} 
                    alt="image" /> 
                  }

                  </div>
                </div>
              ))}
              {hasMore && (
                <button onClick={debouncedFetchComments} disabled={loading}>
                  {loading ? "Loading..." : "Load More"}
                </button>
              )}
              {!hasMore && (
                <p className="no-more-comments">No more comments to load.</p>
              )}
            </>
          )}
        </div>
      )}
    </div>
  );
}
