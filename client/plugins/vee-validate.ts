import Vue from 'vue';
import { localize } from 'vee-validate';
import { ValidationProvider, ValidationObserver } from 'vee-validate';
import ja from "vee-validate/dist/locale/ja.json";
import * as rules from 'vee-validate/dist/rules';

Vue.component('ValidationObserver', ValidationObserver);
Vue.component('ValidationProvider', ValidationProvider);

localize('ja', ja);

Vue.config.productionTip = false;
