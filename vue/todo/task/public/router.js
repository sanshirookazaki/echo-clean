var router = new VueRouter({
    routes: [
        {
            path: '/login',
            component: Login
        }
    ]
  })

var Login = {
    template: '#login'
}

//var app = new Vue({
//    router: router
//}).$mount('#app')