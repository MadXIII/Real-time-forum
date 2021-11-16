import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("FORUM")

    }
    init() {
        let response = fetch('http://localhost:8080/api/')
        if (response.ok) {
            
        }
    }
    async getHtml() {
        let tags = `
            <h1>Home</h1>
            <p>Welcome to the Main Page</p>
            <p><a href="/newpost" data-link>Create Post</a></p>
        `
        return super.isAuth() + tags
    }
}