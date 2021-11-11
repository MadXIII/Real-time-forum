import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("FORUM")

    }
    init() {
        console.log("ok")
        let response = await fetch('http://localhost:8080')
        // if (response.ok) {
        // }
    }
    async getHtml() {
        let tags = `
            <h1>Home</h1>
            <p>Welcome to the Main Page</p>
            <p><a href="/newpost" datat-link>Create Post</a></p>
        `
        return super.isAuth() + tags
    }
}