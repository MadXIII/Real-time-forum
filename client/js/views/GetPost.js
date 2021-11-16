import AbstractView from "./AbstractView.js"


export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Post")
    }

    async init() {
        let searchParams = new URLSearchParams(window.location.search)
        let obj = {
            id: searchParams.get('id'),
        }
        let response = await fetch('http://localhost:8080/api/post', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(obj)
        })
        let submitID = document.getElementById('creatCommentBtnID')
        submitID.onclick = async () => {
            let comment
        }
        // comment: document.getElementById('newComment').value,
        if (response.ok) {
            let res = await response.json()
            let divUsername = document.getElementById("PostDataUsername");
            let divTitle = document.getElementById("PostDataTitle");
            let divContent = document.getElementById("PostDataContent");
            let divTime = document.getElementById("PostDataTime");
            divUsername.innerText = `Username : ${res.username}`;
            divTitle.innerHTML = `Title: ${res.title}`;
            divContent.innerHTML = `Content: ${res.content}`;
            divTime.innerHTML = `Time: ${res.timestamp}`;
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
        <h3>Comments</h3>
        <p><input type="text" placeholder="Comment" id="newComment"/></p>
        <button id="creatCommentBtnID">Submit</button>
        `
    }
}