import AbstractView from "./AbstractView.js";
import { navigateTo } from "../index.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Sign In");
  }

  init() {
    if (document.cookie.includes("session")) {
      navigateTo("/");
    }

    signInForm.onsubmit = async (e) => {
      e.preventDefault();
      e.stopPropagation();
      const formData = new FormData(signInForm);
      const response = await fetch("http://localhost:8080/api/signin", {
        method: "POST",
        body: JSON.stringify({
          login: formData.get("login"),
          password: formData.get("password"),
        }),
      });
      if (response.ok) {
        navigateTo(location.pathname);
      } else {
        let result = await response.json();
        alert(result);
      }
    };
  }

  async getHtml() {
    return `
    <div class="w-full h-full flex justify-center items-center">
        <div class="w-96 bg-white shadow-xl rounded-lg overflow-hidden">
            <div class="flex flex-col items-center gap-1 border-b p-7"><span
                    class="text-lg font-bold">Welcome!</span><span class="wt-margin-bottom-20">Sign in to
                    continue</span></div>
            <div class="flex flex-col items-center gap-1 border-b p-7">
                <form class="flex flex-col w-full gap-3" id="signInForm">
                    <div>
                        <label for="login" class="block text-sm font-medium text-gray-700">Username or
                            email</label>
                        <input type="text" name="login" id="login" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
                        <input type="password" name="password" id="password" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <button type="submit"
                        class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Submit</button>
                    <p class="text-xs text-gray-500 self-center">Don't have an account <a href="/signup" data-link
                            class="text-indigo-600"">Sign up</a></p>
                </form>
            </div>
        </div>
    </div>
        `;
  }
}
