import Vue from "vue";
import Vuex from "vuex";

import auth from "@/store/modules/auth/auth";
import reference from "@/store/modules/reference/reference";

Vue.use(Vuex);

export default new Vuex.Store({
  strict: process.env.NODE_ENV !== "production",
  modules: { auth, reference }
});
