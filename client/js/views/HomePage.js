import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Forum");
  }

  init() {}

  async getHtml() {
    return (
      super.header() +
      `
          <h1>Home</h1>
          <p>Welcome to the Main Page</p>
        `
    );
  }
}
