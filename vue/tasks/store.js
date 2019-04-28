import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
    state: {
        tasks: [
            {
                id: 1,
                name: 'buy milk',
                done: false
        },
        {
            id: 2,
            name: 'buy vuejs book',
            done: true
        }],
    },
})

export default store