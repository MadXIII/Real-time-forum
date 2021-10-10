import AbstractView from "./AbstractView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Log Out")
    }

    init(){
        let logoutID = document.getElementById('logoutBtnId')
        logoutID.onclick = async () => {
            let response = await fetch('http://localhost:8080/logout', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json;charset=utf-8'
                },
            })
            if (response.ok) {
                window.location.href = "/"
            } else {
                let result = await response.json()
                alert(result['notify'])
            }
        }

    
    }

    async getHtml() {
        return `
            <h2>are you sure?</h2>
            <button id="logoutBtnId" >Yes</button >
            <input type="button" onclick="location.href='/'"value="No" >
        `
    }
}