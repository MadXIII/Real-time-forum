import AbstractView from "./AbstractView.js"
import {router} from "../index.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("New Post")
    }

    init() {
        let submitId = document.getElementById('creatPostBtnID')
        submitId.onclick = async () => {
            let newPost = {
                title: document.getElementById("title").value,
                content: document.getElementById("content").value,
            }
            let response = await fetch('http://localhost:8080/newpost', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(newPost)
            })
            if (response.ok) {
                let postID = await response.json()
                // window.location.href = "/post/" + postID
                window.location.replace("http://localhost:8080/post/" + postID)
                             //   window.location.replace("http://localhost:8080/post/" + postID)export const redirect = (endpoint) => {
    //   history.pushState(null, "", `http://localhost:8080/post/${postID}`);
    //     window.addEventListener("popstate", router());
                
            } else {
                let result = await response.json()
                alert(result)
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