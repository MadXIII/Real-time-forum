import HomePage from "./views/HomePage.js"
import Signup from "./views/SignUp.js"
import Signin from "./views/SignIn.js"
import LogOut from "./views/LogOut.js"
import CreatePost from "./views/CreatePost.js"
import GetPost from "./views/GetPost.js"

const navigateTo = url => {
    history.pushState(null, null, url)
    router()
}

export const router = async () => {
    const routes = [
        { path: "/", view: HomePage },
        { path: "/signup", view: Signup },
        { path: "/signin", view: Signin },
        { path: "/logout", view: LogOut },
        { path: "/newpost", view: CreatePost },
        { path: "/post", view: GetPost },
    ];
    
    const potentialMatches = routes.map(route => {
        return {
            route: route,
            isMatch: location.pathname === route.path,
        }
    })
    let match = potentialMatches.find(potentialMatche => potentialMatche.isMatch)

    if (!match) {
        match = {
            route: routes[0],
            isMatch: true
        }
    }

    const view = new match.route.view()

    document.querySelector("#app").innerHTML = await view.getHtml();
    view.init();
}

window.addEventListener("popstate", router)

document.addEventListener("DOMContentLoaded", () => {
    document.body.addEventListener("click", e => {
        if (e.target.matches("[data-link]")) {
            e.preventDefault()
            navigateTo(e.target.href)
        }
    })
    router()
})


