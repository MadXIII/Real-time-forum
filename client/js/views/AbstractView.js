export default class {
  constructor() {}
  setTitle(title) {
    document.title = title;
  }

  header() {
    const isAuthorized = document.cookie.includes("session");
    return `
        <nav class="bg-white py-2 md:py-4">
            <div class="container px-4 mx-auto md:flex md:items-center">

                <div class="flex justify-between items-center">
                    <a href="/" class="font-bold text-xl text-indigo-600" data-link>FRM</a>
                    <button
                        class="border border-solid border-gray-600 px-3 py-1 rounded text-gray-600 opacity-50 hover:opacity-75 md:hidden"
                        id="navbar-toggle">
                        <i class="fas fa-bars"></i>
                    </button>
                </div>
                <div class="hidden md:flex flex-col md:flex-row md:ml-auto mt-3 md:mt-0" id="navbar-collapse">
                    <a href="/create-post"
                        class="p-2 lg:px-4 md:mx-2 text-gray-600 rounded hover:bg-gray-200 hover:text-gray-700 transition-colors duration-300" data-link>Create Post</a>
                    ${
                      isAuthorized
                        ? '<a href="/logout" class="p-2 lg:px-4 md:mx-2 text-indigo-600 text-center border border-solid border-indigo-600 rounded hover:bg-indigo-600 hover:text-white transition-colors duration-300 mt-1 md:mt-0 md:ml-1" data-link>Logout</a>'
                        : '<a href="/signin" class="p-2 lg:px-4 md:mx-2 text-indigo-600 text-center border border-transparent rounded hover:bg-indigo-100 hover:text-indigo-700 transition-colors duration-300" data-link>Login</a><a href="/signup" class="p-2 lg:px-4 md:mx-2 text-indigo-600 text-center border border-solid border-indigo-600 rounded hover:bg-indigo-600 hover:text-white transition-colors duration-300 mt-1 md:mt-0 md:ml-1" data-link>Signup</a>'
                    }
                </div>
            </div>
        </nav>
    `;
  }

  postView(post) {
    return `
    <a href="/post?id=${post.id}" class="flex bg-white shadow-lg rounded-lg w-full">
    <div class="flex flex-col px-4 py-6 w-full">
        <div class="flex items-center justify-between">
            <h2 class="text-lg font-semibold text-gray-900 -mt-1">${post.username}</h2>
            <small class="text-sm text-gray-700">${post.timestamp}</small>
        </div>
        <p class="mt-3 text-gray-700 text-sm">
            ${post.content}
        </p>
        <div class="mt-4 flex items-center">
            <div class="flex mr-2 text-gray-700 text-sm mr-3">
                <svg fill="none" viewBox="0 0 24 24" class="w-4 h-4 mr-1" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M4.318 6.318a4.5 4.5 0 000 6.364L12 20.364l7.682-7.682a4.5 4.5 0 00-6.364-6.364L12 7.636l-1.318-1.318a4.5 4.5 0 00-6.364 0z" />
                </svg>
                <span>12</span>
            </div>
            <div class="flex mr-2 text-gray-700 text-sm mr-8">
                <svg fill="none" viewBox="0 0 24 24" class="w-4 h-4 mr-1" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M17 8h2a2 2 0 012 2v6a2 2 0 01-2 2h-2v4l-4-4H9a1.994 1.994 0 01-1.414-.586m0 0L11 14h4a2 2 0 002-2V6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2v4l.586-.586z" />
                </svg>
                <span>8</span>
            </div>
        </div>
    </div>
</a>
      `;
  }
}
