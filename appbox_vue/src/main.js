import Vue from 'vue'
import App from './App.vue'
import Notifications from 'vue-notification'

Vue.use(Notifications)
import "bootstrap";

import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.min.js";

Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')

