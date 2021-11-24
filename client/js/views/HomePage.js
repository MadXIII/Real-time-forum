import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("FORUM")

    }
    
    init() {
    }

    async getHtml() {
        let tags = `
            <h1>Home</h1>
            <p>Welcome to the Main Page</p>
        `
        return super.isAuth() + tags
    }
}