import {createStore} from 'vuex'

export default createStore({
    state: {
        userInfo: {
            name: 'az'
        }
    },
    mutations: {
        getUserInfo(state, name) {
            state.userInfo.name = name
        }
    },
    actions: {
        asyncGetUserInfo({commit}) {
            setTimeout(() => {
                commit("getUserInfo", +new Date() + 'action')
            }, 2000)
        }
    },
    modules: {},
    getters: {
        userInfoGetter(state) {
            return state.userInfo.name
        }
    }
})
