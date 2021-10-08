import AbstractView from "./AbstractView.js"

export default class extends AbstractView{
    constructor(){
        super()
        this.setTitle("New Post")
    }

    init() {
    }
    async getHtml() {
        return (`
            <div>Emtpy Page</div>
        `)
    }
}