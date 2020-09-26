import { NuxtAxiosInstance } from '@nuxtjs/axios';

interface DeleteSignoutResponse {}

export const signout = (
  axios: NuxtAxiosInstance
): Promise<DeleteSignoutResponse> =>
  axios.$delete<DeleteSignoutResponse>('/api/auth/v1/signout');
