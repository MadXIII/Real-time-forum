import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("FORUM")

    }

    async init() {
        let response = await fetch('http://localhost:8383/api/')

        if (response.ok) {
            let result = await response.json()
            result.Categories.forEach(element => {
                cateogriesID.innerHTML += `<option value="${element.id}">${element.category_name}</option>`
            })
            result.Posts.forEach(element => {
                postsID.innerHTML += `<p><div>Title:${element.title}</div><div>Content:${element.content}</div><div>Username:${element.username}</div><div>time:${element.timestamp}</div></p>`
            })
        } else {
            console.log(false)
        }

        submitID.onclick = async () => {
            let category = {
                id: parseInt(cateogriesID.value)
            }
            let response = await fetch('http://localhost:8383/api/', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
                body: JSON.stringify(category)
            })

            if (response.ok) {
                let result = await response.json()
                postsID.innerHTML = ""
                result.forEach(element => {
                    postsID.innerHTML += `<p><div>Title:${element.title}</div><div>Content:${element.content}</div><div>Username:${element.username}</div><div>time:${element.timestamp}</div></p>`
                })
            } else {
                let result = await response.json()
                alert(result)
            }
        }
    }

    async getHtml() {
        let tags = `
            <h1>Home</h1>
            <select id="cateogriesID"></select>
            <button id="submitID">Submit</button>
            <p>Welcome to the Main Page</p>
            <div id="postsID"></div>
        `
        return super.isAuth() + tags
    }
}