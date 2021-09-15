import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("FORUM")

    }
    init() {
    }
    async getHtml() {
        return `
        <p><a href="/signup" data-link>Sign Up</a></p>
        <p><a href="/signin" data-link>Sign In</a></p>
        <h1>Home</h1>
        <p>Welcome to the Main Page</p>
        <p><a href="/post" datat-link>View Posts</a></p>
        `
    }
}