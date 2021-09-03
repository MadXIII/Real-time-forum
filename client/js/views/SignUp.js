import AbstractView from "./AbstratcView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("SignUp")

    }

    async getHtml() {
        return `
        <h1>Real Time Forum</h1>
        <p><input type="text" placeholder="Nickname" id="nick"></p>
        <p><input type="text" placeholder="Email address" id="email"></p>
        <p><input type="text" placeholder="Password" id="pass"></p>
        <p><input type="text" placeholder="First name" id="firstname"></p>
        <p><input type="text" placeholder="Last name" id="lastname"></p>
        <p><input type="text" placeholder="Gender" id="gender"></p>
        <p><input type="text" placeholder="Age" id="age"></p>
        <button type="submit">Sing Up</button>
        or
        <button type="submit">Sing In</button>
        `
    }

}