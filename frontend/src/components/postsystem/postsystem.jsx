"use client";
import { useEffect, useState } from "react";
import CreatePost from "../createPost/createPost";
import Post from "../post/post";

export default function PostSystem() {
    const [posts, setPosts] = useState([]);

    const host = process.env.NEXT_PUBLIC_HOST;
    console.log("Host", host);
    

    const fetchAllPosts = async () => {
        try {
            const response = await fetch(`${host}/api/posts`, {
                cache: "no-store",
            });
            if (!response.ok) {
                throw new Error("Failed to fetch posts");
            }
            const data = await response.json();
            setPosts(data);
        } catch (err) {
            console.error("Fetch error:", err);
        }
    };
    const addNewPost = (newPost) => {
        setPosts((prev) => [newPost, ...prev]);
    };

    useEffect(() => {
        fetchAllPosts();
    }, []);

    return (
        <>
            {/* <CreatePost onPostCreated={addNewPost} /> */}
            <CreatePost onPostCreated={addNewPost}/>
            <Post posts={posts} />
        </>
    );
}
