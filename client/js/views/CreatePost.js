import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("New Post")
    }

    init() {
        let submitId = document.getElementById('creatPostBtnID')
        submitId.onclick = async () => {
            console.log(23)
            let newPost = {
                title: document.getElementById("title").value,
                content: document.getElementById("content").value,
                // timestamp:
            }
            let response = await fetch('http://localhost:8080/newpost', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(newPost)
            })
            if (response.ok) {
                window.location.href = "/"
            }
        }
    }
    async getHtml() {
        return (`
            <a id="logout" href="/logout" data-link>Log Out</a>
            <div>Create Post</div>
            <p><input type="text" placeholder="Title" id="title"/></p>
            <p><input type="text" placeholder="Your Post" id="content"/></p>
            <button id="creatPostBtnID">Submit</button>
        `)
    }
}