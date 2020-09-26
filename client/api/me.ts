import { NuxtAxiosInstance } from '@nuxtjs/axios';
import { User } from './types/user';

interface GetMeResponse {
  user: User;
}

export const get = (axios: NuxtAxiosInstance): Promise<GetMeResponse> =>
  axios.$get<GetMeResponse>('/api/me/v1/get');
