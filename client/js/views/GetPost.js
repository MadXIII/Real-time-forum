import AbstractView from "./AbstractView.js"


export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Post")
    }

    async init() {
        let urlID = new URLSearchParams(window.location.search).get('id')

        let response = await fetch(`http://localhost:8282/api/post?id=${urlID}`)

        if (response.ok) {
            let res = await response.json()

            PostDataUsername.innerText = `Username: ${res.Post.username}`;
            PostDataTitle.innerHTML = `Title: ${res.Post.title}`;
            PostDataContent.innerHTML = `Content: ${res.Post.content}`;
            PostDataTime.innerHTML = `Time: ${res.Post.timestamp}`;
            PostLikeCount.innerHTML = `${res.Post.likes}`;
            if (res.Comments != null) {
                comment1Username.innerHTML = `CommentUsername: ${res.Comments[0].username}`;
                comment1Timestamp.innerHTML = `CommentTimestamp: ${res.Comments[0].timestamp}`;
                comment1Content.innerHTML = `CommentContent: ${res.Comments[0].content}`;
            }
        } else {
            let res = await response.json()
            alert(res)
        }

        creatCommentBtnID.onclick = async () => {
            let obj = {
                cpost_id: parseInt(urlID),
                content: newComment.value,
            }
            let response = await fetch('http://localhost:8282/api/post', {
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
                let res = await response.json()
                alert(res)
                window.location.href = `/signin`
            }
        }


        likeBtnID.onclick = async () => {
            let obj = {
                post_id: parseInt(urlID),
                vote_type: "like",
            }
            let response = await fetch('http://localhost:8282/api/post', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(obj)
            })
            if (response.ok) {
                window.location.href = `/post?id=${urlID}`
            } else {
                let res = await response.json()
                alert(res)
                window.location.href = `/signin`
            }
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
        <button id="likeBtnID">&#128402</button>
        <div id="PostLikeCount">0</div>
        <h3>Comments</h3>
        <p><input type="text" placeholder="Comment" id="newComment"/></p>
        <button id="creatCommentBtnID">Submit</button>
        <div id="comment1Username"></div>
        <div id="comment1Timestamp"></div>
        <div id="comment1Content"></div>
        `
    }
}