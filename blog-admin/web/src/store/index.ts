import {createStore} from 'vuex'

export default createStore({
  // all data save in app
  state: {
    user: {
      token: ''
    }
  },
  // all functions change state
  mutations: {
    SET_TOKEN(state, token: string) {
      state.user.token = token
    }
  },
  // commit mutations
  actions: {
    setToken({commit}, token) {
      return new Promise((resolve) => {
        commit('SET_TOKEN', token)
        resolve()
      })
    }
  },
  // sub modules
  modules: {},
  // get all states
  getters: {
    token(state) {
      return state.user.token
    }
  }
})
