import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Log Out");
  }

  init() {
    (async () => {
      let response = await fetch("http://localhost:8080/api/logout", {
        method: "POST",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
        },
      });
      if (response.ok) {
        document.cookie = `${document.cookie}; max-age=0`;
        let result = await response.json();
        alert(result);
        window.location.href = "/";
      } else {
        let result = await response.json();
        alert(result);
      }
    })();
  }

  async getHtml() {
    return `<div></div>`;
  }
}
