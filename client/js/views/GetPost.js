import AbstractView from "./AbstractView.js"

export default class extends AbstractView{
    constructor() {
        super()
        this.setTitle("Post")
    }
    
    init(){

    }

    async getHtml(){
        return (`
            <div>TEST<div>
        `)
    }
}