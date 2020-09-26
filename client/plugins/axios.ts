import { NuxtAxiosInstance } from '@nuxtjs/axios';
import { AxiosError } from 'axios';
import { Route } from 'vue-router';

export default ({ $axios, redirect }: { $axios: NuxtAxiosInstance, redirect(path: string, query?: Route['query']): void }) => {
  $axios.onError((error: AxiosError) => {
    const code = error.response && error.response!.status;
    if(code === 401) {
      redirect('/signin');
    }
  })
}
