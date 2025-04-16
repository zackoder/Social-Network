import "./createPost.modules.css"

export default function CreatePost(){
    return (
        <div className="postContainer">
            <form>
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
                        <select name="friends">
                            <option value={"public"}>Public</option>
                            <option value={"private"}>Private</option>
                            <option value={"almostPrivate"}>Almost Private </option>
                        </select>
                    </div>
                </div>
                <div className="title">
                    <input type="text" name="title" placeholder="enter your title" />
                </div>
                <div className="content">
                    {/* <textarea name="content" placeholder="enter your content"></textarea> */}
                    <textarea
                        placeholder="Write a content..."
                        // value={comment}
                        // onChange={(e) => setComment(e.target.value)}
                        rows={4}
                        style={{ width: '90%', padding: '5px', borderRadius: '4px', resize: 'none', outline: 'none', border: 'none', marginLeft: '5%', backgroundColor: '#333'}}
                    />
                </div>
                <div className="uploadImage">
                    <input type="file" name="image" id="" />
                </div>
                <div>
                    <input className="submit" type="submit" value="Publish" />
                </div>
            </form>
        </div>
    );
}