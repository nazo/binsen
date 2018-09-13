import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';
import { User } from './types/user';

interface GetMeResponse {
  user: User;
}

export const get = (axios: NuxtAxiosInstance): AxiosPromise<GetMeResponse> =>
  axios.$get<GetMeResponse>('/api/me/v1/get');
