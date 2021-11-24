export default class {
    constructor() {

    }
    setTitle(title) {
        document.title = title
    }

    isAuth() {
        let tags = ""
        if (document.cookie.indexOf('session') != -1) {
            tags = `
            <a id="logout" href="/logout" data-link>Log Out</a>
            <p><a href="/newpost" data-link>Create Post</a></p>
            `
        } else {
            tags = `
            <a id="signup" href="/signup" data-link>Sign Up</a>
            <a id="signin" href="/signin" data-link>Sign In</a>
            `
        }
        return tags
    }
}