import styles from './register.module.css'

export default function Register() {
    return (
        <main className={styles.main}>
            <div className={styles.containerRegister}>
                <h2 className={styles.title}>Register+"1231"</h2>
                <form action="">
                    <div className={styles.name}>
                        <input type="text" className={styles.firstname} name='firstname' placeholder='First Name' />
                        <input type="text" className={styles.lastname} name='lastname' placeholder='Last Name' />
                    </div>
                    <div className={styles.nicknameAndAge}>
                        <input type="text" className={styles.nickname} name='nickname' placeholder='Nickname' />
                        <input type="date" className={styles.age} name='age' title='Date of Birth' />
                    </div>
                    <div className={styles.genderAndAvatar}>
                        <select name='gender'>
                            <option value={"male"}>Male</option>
                            <option value={"female"}>Female</option>
                        </select>
                        <input type="file" className={styles.avatar} name='avatar' title='upload your Image' placeholder='upload your Image' />
                    </div>
                    <div className={styles.email}>
                        <input type="email" className={styles.email} name='email' placeholder='Email' />
                    </div>
                    <div className={styles.password}>
                        <input type="password" className={styles.passwordInput} name='password' placeholder='password' />
                        <input type="password" className={styles.confirmPassword} name='confirmPassword' placeholder='Confirm Password' />
                    </div>
                    <div className={styles.aboutme}>
                        <textarea rows={3} className={styles.aboutme} name='aboutme' placeholder='What About Me' >
                        </textarea>
                    </div>
                    <div className={styles.submit}>
                        <input type="submit" className={styles.submit} value="Register" />
                    </div>
                </form>
            </div>
        </main>
    );
}