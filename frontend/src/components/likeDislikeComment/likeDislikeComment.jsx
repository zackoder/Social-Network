"use client"
// import styles from ""
import { BiLike, BiDislike } from "react-icons/bi";
import React, { useState } from 'react';
// import styles from './likeDislikeComment.modules.css';
import './likeDislikeComment.modules.css';

export default function LikeDislikeComment() {
  const [liked, setLiked] = useState(false);
  const [disliked, setDisliked] = useState(false);
  const [comment, setComment] = useState('');
  const [submittedComment, setSubmittedComment] = useState('');

  const handleLike = () => {
    setLiked(!liked);
    if (disliked) setDisliked(false);
  };

  const handleDislike = () => {
    setDisliked(!disliked);
    if (liked) setLiked(false);
  };

  const handleCommentSubmit = (e) => {
    e.preventDefault();
    setSubmittedComment(comment);
    setComment('');
  };

  return (
    //className={styles.reactionContainer}
    <div style={{minWidth: '100%', margin: '5px auto' }}>
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

      <form className="form" onSubmit={handleCommentSubmit}>
            <textarea
                placeholder="Write a comment..."
                value={comment}
                onChange={(e) => setComment(e.target.value)}
                rows={2}
                style={{ width: '90%', padding: '5px', borderRadius: '4px', resize: 'none' }}
            />
            <button
            type="submit"
            style={{
                marginTop: '0px',
                marginLeft: '5px',
                width: '10%',
                padding: '14px 0',
                backgroundColor: 'var(--color-primary)',  //0070f3
                color: 'white',
                fontWeight: 'bold',
                border: 'none',
                borderRadius: '4px',
                cursor: 'pointer',

            }}
            >
            Send
            </button>
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
