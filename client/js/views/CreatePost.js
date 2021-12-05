import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("New Post")
    }

    init() {
        let response =  fetch('http://localhost:8080/api/newpost')

        if (response.ok) {
            console.log(true)
            let res = response.json()
            console.log(res.id)
        } else {
            console.log(false)
        }

        let submitId = document.getElementById('creatPostBtnID')
        submitId.onclick = async () => {
            let newPost = {
                title: titleID.value,
                content: contentID.value,
            }
            let response = await fetch('http://localhost:8080/api/newpost', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(newPost)
            })
            if (response.ok) {
                let result = await response.json()
                alert(result.notify)
                window.location.replace(`http://localhost:8080/post?id=${result.id}`)
            } else {
                let result = await response.json()
                alert(result)
            }
        }
    }
    async getHtml() {
        return `
            <a id="logout" href="/logout" data-link>Log Out</a>
            <div>Create Post</div>
            <p><input type="text" placeholder="Title" id="titleID"/></p>
            <p><input type="text" placeholder="Your Post" id="contentID"/></p>
            <select id="cateogriesID"></select>
            <button id="creatPostBtnID">Submit</button>
            `
    }
}