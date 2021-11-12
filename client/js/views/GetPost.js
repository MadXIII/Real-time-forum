import AbstractView from "./AbstractView.js"


export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Post")
    }

    init() { 

        let searchParams = new URLSearchParams(window.location.search)
        console.log(searchParams.get('id'));

    }

    async getHtml() {
        return `
            <div>TEST<div>
        `
    }
}