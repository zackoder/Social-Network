"use client"
import { use, useState } from "react";
import "./createPost.modules.css"

export default function CreatePost() {
    
    let [privacy, setPrivacy] = useState("public")
    let [title, setTitle] = useState("")
    let [content, setContent] = useState("")
    let [image, setImage] = useState(null)

    const postData = {
        privacy: privacy,
        title: title,
        content: content
    }
    
    const host = process.env.NEXT_PUBLIC_HOST

    const handleSubmit = async (e) => {
        e.preventDefault();
        const formData = new FormData();
        formData.append('postData', JSON.stringify(postData))
        if (image) {            
            formData.append('avatar', image);
        }
        const response = await fetch(`${host}/addPost`, { //${host}
            method: "POST",
            body: formData
        })

        if (!response.ok) {
            console.log("error not ok");
        }
    }


    return (
        <div className="postContainer">
            <form onSubmit={handleSubmit}>
                <div className="identityProfile">
                    <div className="imageProfile">
                        {/* <Image
                            className={styles.image}
                            src="/images/post.png"
                            alt="post"
                            // width={500}
                            // height={500}
                            fill={true}
                        /> */}
                    </div>
                    <div className="nameProfile">
                        <h3>full name</h3>
                        <select onChange={(e) => { setPrivacy(e.target.value) }} name="friends" id="friends" defaultValue={"public"}>
                            <option value={"public"}>Public</option>
                            <option value={"private"}>Private</option>
                            <option value={"almostPrivate"}>Almost Private </option>
                        </select>
                    </div>
                </div>
                <div className="title">
                    <input onChange={(e) => { setTitle(e.target.value) }} type="text" name="title" placeholder="enter your title" />
                </div>
                <div className="content">
                    {/* <textarea name="content" placeholder="enter your content"></textarea> */}
                    <textarea
                        onChange={(e) => { setContent(e.target.value) }}
                        placeholder="Write a content..."
                        // value={comment}
                        // onChange={(e) => setComment(e.target.value)}
                        rows={4}
                        style={{ width: '90%', padding: '5px', borderRadius: '4px', resize: 'none', outline: 'none', border: 'none', marginLeft: '5%', backgroundColor: '#333' }}
                    />
                </div>
                <div className="uploadImage">
                    <input onChange={(e) => {setImage(e.target.files[0]) }} type="file" name="image" />
                </div>
                <div>
                    <input className="submit" type="submit" value="Publish" />
                </div>
            </form>
        </div>
    );
}

