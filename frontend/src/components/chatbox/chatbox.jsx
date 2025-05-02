import styles from "./chatbox.module.css"
import Image from "next/image";
import { FaCloudUploadAlt } from "react-icons/fa";
import { IoIosSend } from "react-icons/io";

export default function ChatBox() {
    return (
        <div className={styles.chatBox}>
            <div className={styles.header}>
                <div className={styles.imgProfile}>
                    <Image
                        src="/profile/profile.png"
                        alt="image profile"
                        fill
                        style={{ objectFit: 'cover', borderRadius: '50%' }}
                    />
                </div>
                <div className={styles.infoProfile}>
                    <h3>full name</h3>
                    <p>Lorem ipsum dolor sit amet.</p>
                </div>
                <div className={styles.close}>x</div>
            </div>
            <div className={styles.readmessages}>
                <div className={styles.me}>
                    <div className={styles.message}>
                        <p>Lorem ipsum dolor sit amet.</p>
                        <span>8:55Am, Today</span>
                    </div>
                    <div className={styles.profileImage}>
                        <Image
                            src="/profile/profile.png"
                            alt="image profile"
                            fill
                            style={{ objectFit: 'cover', borderRadius: '50%' }}
                        />
                    </div>
                </div>
                <div className={styles.sender}>
                    <div className={styles.profileImage}>
                        <Image
                            src="/profile/profile.png"
                            alt="image profile"
                            fill
                            style={{ objectFit: 'cover', borderRadius: '50%' }}
                        />
                    </div>
                    <div className={styles.message}>
                        <p>Lorem ipsum dolor sit amet.</p>
                        <span>8:55Am, Today</span>
                    </div>
                </div>
            </div>
            <div className={styles.sendmessages}>
                <form action="">
                    <div className={styles.imogie}>
                        😁​😅​🤣​😂​🙂​🙃​🫠​😉​🥰​😍​🤩​☺️​🥲​😛​😜​🤗​🤭​🤫​🤔​🫡​🫥​😒​🙄​🙂‍↔️​🙂‍↕️​🥵​🤯​🥳​😎​😎​🤓​🥺​🥹​😥​😱​😭​👋​👌​🤞​👉​👇​👍​👏
                    </div>
                    <div className={styles.elementsSend}>
                        <input type="file" name="uploadImage" id="uploadImage" className={styles.hiddenInput} />
                        <label htmlFor="uploadImage" className={styles.uploadLabel}>
                            <FaCloudUploadAlt className={styles.iconUpload} />
                        </label>

                        <input type="text" name="message" placeholder="Type your message..." id="" />
                        <input type="submit" name="submit" className={styles.hiddenInput} id="submit" />
                        <label htmlFor="submit" className={styles.labelSend}>
                            <IoIosSend className={styles.iconSend} />
                        </label>
                    </div>
                </form>
            </div>
        </div>
    );
}