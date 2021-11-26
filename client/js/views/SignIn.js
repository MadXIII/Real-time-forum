import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Sign In")
    }

    init() {
        signInBtnId.onclick = async () => {
            let user = {
                login: loginID.value, 
                password: passID.value,
            }
            let response = await fetch('http://localhost:8080/api/signin', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
            })
            if (response.ok) {
                let result = await response.json()
                alert(result)
                window.location.href = "/"
            } else {
                let result = await response.json()
                alert(result)
            }
        }
    }

    async getHtml() {
        return `
            <h1>SIGN IN</h1>
            <p><input type="text" placeholder="Nickname or Email" id="loginID"</p>
            <p><input type="password" placeholder="Password" id="passID"</p>
            <input type="button" onclick="location.href='/signup';" value="Sign Up"/>
            or
            <button id="signInBtnId" type="submit">Sign In</button>
        `
    }
}