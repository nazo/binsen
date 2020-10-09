import { NuxtHTTPInstance } from '@nuxt/http';
import { Route } from 'vue-router';

export default ({ $http, redirect }: { $http: NuxtHTTPInstance, redirect(path: string, query?: Route['query']): void }) => {
  $http.onError(error => {
    const code = error.response && error.response!.status;
    if(code === 401) {
      redirect('/signin');
    }
  })
}