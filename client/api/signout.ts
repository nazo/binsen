import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';

interface DeleteSignoutResponse {
}

export const signout = (axios: NuxtAxiosInstance): AxiosPromise<DeleteSignoutResponse> => axios.$delete<DeleteSignoutResponse>('/api/auth/v1/signout');
