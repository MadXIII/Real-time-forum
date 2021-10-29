import AbstractView from "./AbstractView.js"

export default class extends AbstractView{
    constructor() {
        super()
        this.setTitle("Post")
    }
    
    init(){
        console.log(123);
    }
    
    async getHtml(){
        return `
            <div>TEST<div>
        `
    }
}