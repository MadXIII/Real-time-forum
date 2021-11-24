import AbstractView from "./AbstractView.js"


export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Post")
    }

    async init() {
        let urlID = new URLSearchParams(window.location.search).get('id')

        let submitID = document.getElementById('creatCommentBtnID')
        submitID.onclick = async () => {
            let obj = {
                post_id: parseInt(`${urlID}`),
                content: document.getElementById('newComment').value,
            }
            let response = await fetch('http://localhost:8080/api/post', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(obj)
            })
            if (response.ok) {
                let result = await response.json()
                alert(result)
                window.location.href = `/post?id=${urlID}`
            } else {
                let result = await response.json()
                alert(result)
            }
        }

        let response = await fetch(`http://localhost:8080/api/post?id=${urlID}`)

        if (response.ok) {
            let res = await response.json()
            let divPostUsername = document.getElementById("PostDataUsername");
            let divPostTitle = document.getElementById("PostDataTitle");
            let divPostContent = document.getElementById("PostDataContent");
            let divPostTime = document.getElementById("PostDataTime");

            let divCommentUsername = document.getElementById("comment1Username");
            let divCommentTimestamp = document.getElementById("comment1Timestamp");
            let divCommentContent = document.getElementById("comment1Content");

            divPostUsername.innerText = `Username: ${res.Post.username}`;
            divPostTitle.innerHTML = `Title: ${res.Post.title}`;
            divPostContent.innerHTML = `Content: ${res.Post.content}`;
            divPostTime.innerHTML = `Time: ${res.Post.timestamp}`;
            if (res.Comments != null) {
                divCommentUsername.innerHTML = `CommentUsername: ${res.Comments[0].username}`;
                divCommentTimestamp.innerHTML = `CommentTimestamp: ${res.Comments[0].timestamp}`;
                divCommentContent.innerHTML = `CommentContent: ${res.Comments[0].content}`;
            }
        } else {
            let res = await response.json()
            alert(res)
        }
    }

    async getHtml() {
        return `
        <p><a href="/" data-link>Home</a></p>
        <h1>Post</h1>
        <div id="PostDataUsername"></div>
        <div id="PostDataTitle"></div>
        <div id="PostDataContent"></div>
        <div id="PostDataTime"></div>
        <button id="likeBtnID">Like</button>
        <button id="dislikeBtnID">Dislike</button>
        <h3>Comments</h3>
        <p><input type="text" placeholder="Comment" id="newComment"/></p>
        <button id="creatCommentBtnID">Submit</button>
        <div id="comment1Username"></div>
        <div id="comment1Timestamp"></div>
        <div id="comment1Content"></div>
        `
    }
}