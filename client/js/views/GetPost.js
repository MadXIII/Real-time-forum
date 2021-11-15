import AbstractView from "./AbstractView.js"


export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Post")
    }

    async init() {
        let searchParams = new URLSearchParams(window.location.search)
        let obj = { id: searchParams.get('id') }

        let response = await fetch('http://localhost:8080/api/post', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=utf-8'
            },
            body: JSON.stringify(obj)
        })
        if (response.ok) {
            let res = await response.json()
            let divUsername = document.getElementById("PostDataUsername");
            let divTitle = document.getElementById("PostDataTitle");
            let divContent = document.getElementById("PostDataContent");
            let divTime = document.getElementById("PostDataTime");
            divUsername.innerText = `Post Username : ${res.username}`;
            divTitle.innerHTML = `Post Title: ${res.title}`;
            divContent.innerHTML = `Post Content: ${res.content}`;
            divTime.innerHTML = `Post Time: ${res.timestamp}`;
        } else {
        }
    }

    async getHtml() {
        return `
        <div id="PostDataUsername"></div>
        <div id="PostDataTitle"></div>
        <div id="PostDataContent"></div>
        <div id="PostDataTime"></div>
        `
    }
}