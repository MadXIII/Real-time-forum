import AbstractView from "./AbstratcView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Post")

    }
    
    async getHtml() {
        return `
            <h1>All posts</h1>
            <button onclick=""> send </button>
        `
            //array posts, each post - click, url/api/postbyid/:id

    }
}