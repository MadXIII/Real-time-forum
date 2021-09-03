import AbstractView from "./AbstratcView.js"

export default class extends AbstractView {
    constructor() {
        super()
        this.setTitle("Home")

    }
    init(){
        console.log('init')
        super.showPussy()
    }

    async getHtml() {
        return `
            <h1>Welcome my friend</h1>
            <p>
                This is the Real fucking time Forum!!!
            </p>
            <p>
                <a href="/post" datat-link>View Posts</a>
            </p>
        `
    }
}