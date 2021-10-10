import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("New Post")
    }

    init() {
        let submitId = document.getElementById('submitBtnId')
        submitId.onclick = async () => {
            let newPost = {
                title: getElementById("title").value,
                content: getElementById("content").value,
                // timestamp:
            }
        }
    }
    async getHtml() {
        return (`
            <a id="logout" href="/logout" data-link>Log Out</a>
            <div>Create Post</div>
            <p><input type="text" placeholder="Title" id="title"/></p>
            <p><input type="text" placeholder="Your Post" id="content"/></p>
            
        `)
    }
}