"use client";
import { useEffect, useState } from "react";
import CreatePost from "../createPost/createPost";
import Post from "../post/post";
import { isAuthenticated } from "@/app/page";

export default function PostSystem() {
    const [posts, setPosts] = useState([]);

    const host = process.env.NEXT_PUBLIC_HOST;

    const fetchAllPosts = async () => {
        try {
            const response = await fetch(`${host}/api/posts`, {
                cache: "no-store",
            });
            if (!response.ok) {
                // throw new Error("Failed to fetch posts");
            }
            const data = await response.json();            
            setPosts(data);
        } catch (err) {
            // console.error("Fetch error:", err);
        }
    };
    const addNewPost = (newPost) => {
        if (!posts || posts.length === 0) {
            setPosts([newPost])
        } else {
            setPosts((prev) => [newPost, ...prev]);
        }
    };

    useEffect(() => {
        fetchAllPosts();
    }, []);

    return (
        <>
            {/* <CreatePost onPostCreated={addNewPost} /> */}
            <CreatePost onPostCreated={addNewPost} />
            <Post posts={posts} />
        </>
    );
}
