import { navigateTo } from "../index.js";
import AbstractView from "./AbstractView.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("New Post");
  }

  init() {
    let submitId = document.getElementById("creatPostBtnID");
    newPostForm.onsubmit = async (e) => {
      e.preventDefault();
      e.stopPropagation();
      const formData = new FormData(newPostForm);
      let response = await fetch("http://localhost:8080/api/newpost", {
        method: "POST",
        headers: {
          "Content-Type": "application/json;charset=utf-8",
        },
        body: JSON.stringify({
          title: formData.get("title"),
          content: formData.get("content"),
        }),
      });
      if (response.ok) {
        let result = await response.json();
        navigateTo(`http://localhost:8080/post?id=${result.id}`);
      } else {
        let result = await response.json();
        alert(result);
      }
    };
  }
  async getHtml() {
    return (
      super.header() +
      `
        <div class="w-full h-full flex justify-center items-center">
            <div class="bg-white shadow-xl rounded-lg overflow-hidden p-7" style="width:500px">
                <form class="flex flex-col w-full gap-3" id="newPostForm">
                    <div>
                        <label for="title" class="block text-sm font-medium text-gray-700">Title</label>
                        <input type="text" name="title" id="title" autocomplete="given-name" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="content" class="block text-sm font-medium text-gray-700">Post</label>
                        <textarea name="content" id="content" rows="5" class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2"></textarea>
                    </div>
                    <button type="submit" class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Submit</button>
                </form>
            </div>
        </div>
      `
    );
  }
}
