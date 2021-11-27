import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Forum");
  }

  init() {}

  async getHtml() {
    const data = await fetch("/api/posts")
      .then((res) => res.json())
      .then((data) => data);

    let posts = "";
    data?.map((post) => (posts += super.postView(post)));

    return (
      super.header() +
      `<div class="flex flex-col w-full items-center max-w-2xl gap-4 self-center mt-10">${posts}</div>`
    );
  }
}
