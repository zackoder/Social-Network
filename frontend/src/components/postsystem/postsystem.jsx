"use client";
import { useEffect, useState, useCallback } from "react";
import CreatePost from "../createPost/createPost";
import { debounce } from "@/utils/debounce";
import Post from "../post/post";

export default function PostSystem() {
  const [posts, setPosts] = useState([]);
  const [offset, setOffset] = useState(10);
  const [hasMore, setHasMore] = useState(true);
  const [loading, setLoading] = useState(false);
  const host = process.env.NEXT_PUBLIC_HOST;
  const LIMIT = 10;

  const fetchAllPosts = async () => {
    console.log(loading, hasMore);

    if (loading || !hasMore) return;
    setLoading(true);
    try {
      const response = await fetch(
        `${host}/api/posts?offset=${offset}&limit=${LIMIT}`,
        {
          cache: "no-store",
        }
      );
      if (!response.ok) {
        // throw new Error("Failed to fetch posts");
      }
      const data = await response.json();
      console.log("2222222222222222222222222222222");

      //   console.log(data.length, LIMIT);
      if (posts.length === 0 && data === null) {
        setPosts(Array.isArray(data) ? data : []);
      }
      if (posts.length < offset && data === null) {
        setHasMore(false); // No more posts available
        return;
      }
      setPosts((prev) => [...data, ...prev]);
      setOffset((prev) => prev + LIMIT);
    } catch (err) {
      console.error("Fetch error:", err);
    } finally {
      console.log("3333333333333333333333333333333333333333");
      console.log(loading, hasMore);
      // setHasMore(true)
      setLoading(false);
    }
  };
  const addNewPost = (newPost) => {
    if (!posts || posts.length === 0) {
      setPosts([newPost]);
    } else {
      setPosts((prev) => [newPost, ...prev]);
    }
  };

  useEffect(() => {
    fetchAllPosts();
  }, []);

  const debouncedFetchPosts = useCallback(debounce(fetchAllPosts, 300), [
    offset,
    hasMore,
    loading,
  ]);
  return (
    <>
      {/* <CreatePost onPostCreated={addNewPost} /> */}
      <CreatePost onPostCreated={addNewPost} />
      <Post posts={posts} />
      {hasMore ? (
        <button onClick={debouncedFetchPosts} disabled={loading}>
          {loading ? "Loading..." : "Load More"}
        </button>
      ) : (
        <button  >
          {  "there are no more posts"}
        </button>
      )}
    </>
  );
}
