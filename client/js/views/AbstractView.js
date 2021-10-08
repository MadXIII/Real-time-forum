export default class {
    constructor() {

    }
    setTitle(title) {
        document.title = title
    }
    async getHtml() {
        return ""
    }

    isAuth() {
        let temp = ""
        if (document.cookie.indexOf('session') != -1) {
            temp = `
            <a id="logout" href="/logout" data-link>Log Out</a>
            `
        } else {
            temp = `
            <a id="signup" href="/signup" data-link>Sign Up</a>
            <a id="signin" href="/signin" data-link>Sign In</a>
            `
        }
        return temp
    }
}