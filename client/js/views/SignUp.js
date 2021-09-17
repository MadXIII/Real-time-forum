import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Sign Up")
    }

    init() {
        let signUpId = document.getElementById('signUpBtnId')
        signUpId.onclick = async () => {
            let user = {
                nickname: document.getElementById('nick').value,
                email: document.getElementById('email').value,
                password: document.getElementById('pass').value,
                confirm: document.getElementById('confirm').value,
                first_name: document.getElementById('firstname').value,
                last_name: document.getElementById('lastname').value,
                gender: document.getElementById('gender').value,
                age: parseInt(document.getElementById('age').value),
            }
            let response = await fetch('http://localhost:8080/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(user)
            })
            if (response.ok) {
                window.location.href = "/"
            } else {
                let result = await response.json()
                alert(result['notify'])
            }
        }
    }

    async getHtml() {
        return `
        <h1>SIGN UP</h1>
        <p><input type="text" placeholder="Nickname" id="nick"></p>
        <p><input type="text" placeholder="Email address" id="email"></p>
        <p><input type="text" placeholder="Password" id="pass"></p>
        <p><input type="text" placeholder="Confrim Password" id="confirm"></p>
        <p><input type="text" placeholder="First name" id="firstname"></p>
        <p><input type="text" placeholder="Last name" id="lastname"></p>
        <p><input type="text" placeholder="Gender" id="gender"></p>
        <p><input type="text" placeholder="Age" id="age"></p>
        <input type="button" onclick="location.href='/signin';" value="Sign In"/>
        or
        <button id="signUpBtnId" type="submit">Sign Up</button>
        `
    }
}