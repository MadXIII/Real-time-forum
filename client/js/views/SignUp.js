import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Sign Up")
    }

    init() {
        signUpBtnId.onclick = async () => {
            let user = {
                nickname: nickID.value,
                email: emailID.value,
                password: passID.value,
                confirm: confirmID.value,
                first_name: firstnameID.value,
                last_name: lastnameID.value,
                gender: genderID.value,
                age: ageID.value,
            }
            let response = await fetch('http://localhost:8282/api/signup', {
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
        <h1>SIGN UP</h1>
        <p><input type="text" placeholder="Nickname" id="nickID"></p>
        <p><input type="text" placeholder="Email address" id="emailID"></p>
        <p><input type="password" placeholder="Password" id="passID"></p>
        <p><input type="password" placeholder="Confrim Password" id="confirmID"></p>
        <p><input type="text" placeholder="First name" id="firstnameID"></p>
        <p><input type="text" placeholder="Last name" id="lastnameID"></p>
        <p><input type="text" placeholder="Gender" id="genderID"></p>
        <p><input type="text" placeholder="Age" id="ageID"></p>
        <input type="button" onclick="location.href='/signin';" value="Sign In"/>
        or
        <button id="signUpBtnId" type="submit">Sign Up</button>
        `
    }
}