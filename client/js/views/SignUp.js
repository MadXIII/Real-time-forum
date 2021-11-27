import AbstractView from "./AbstractView.js";
import { navigateTo } from "../index.js";

export default class extends AbstractView {
  constructor() {
    super();
    this.setTitle("Sign Up");
  }

  init() {
    if (document.cookie.includes("session")) {
      navigateTo("/");
    }

    signUpForm.onsubmit = async (e) => {
      e.preventDefault();
      e.stopPropagation();
      const formData = new FormData(signUpForm);
      const response = await fetch("http://localhost:8080/api/signup", {
        method: "POST",
        body: JSON.stringify({
          nickname: formData.get("nickname"),
          email: formData.get("email"),
          password: formData.get("password"),
          confirm: formData.get("confirm"),
          first_name: formData.get("firstname"),
          last_name: formData.get("lastname"),
          gender: formData.get("gender"),
          age: formData.get("age"),
        }),
      });
      if (response.ok) {
        navigateTo("/");
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
            <div class="flex flex-col items-center gap-1 border-b p-7">
                <span class="text-lg font-bold">Registration</span>
            </div>
            <div class="flex flex-col items-center gap-1 border-b p-7">
                <form class="flex flex-col w-full gap-3" id="signUpForm">
                    <div>
                        <label for="nickname" class="block text-sm font-medium text-gray-700">Nickname</label>
                        <input type="text" name="nickname" id="nickname" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="email" class="block text-sm font-medium text-gray-700">Email address</label>
                        <input type="text" name="email" id="email" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="password" class="block text-sm font-medium text-gray-700">Password</label>
                        <input type="password" name="password" id="password" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="confirm" class="block text-sm font-medium text-gray-700">Confirm password</label>
                        <input type="password" name="confirm" id="confirm" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="firstname" class="block text-sm font-medium text-gray-700">First name</label>
                        <input type="text" name="firstname" id="firstname" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="lastname" class="block text-sm font-medium text-gray-700">Last name</label>
                        <input type="text" name="lastname" id="lastname" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="gender" class="block text-sm font-medium text-gray-700">Gender</label>
                        <input type="text" name="gender" id="gender" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <div>
                        <label for="age" class="block text-sm font-medium text-gray-700">Age</label>
                        <input type="text" name="age" id="age" autocomplete="given-name"
                            class="mt-1 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 border rounded-md py-1 px-2">
                    </div>
                    <button type="submit"
                        class="inline-flex justify-center py-2 px-4 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">Submit</button>
                    <p class="text-xs text-gray-500 self-center">Already have an account <a href="/signin" data-link
                            class="text-indigo-600"">Sign in</a></p>
                </form>
            </div>
        </div>
    </div>
        `;
  }
}
