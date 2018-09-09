import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';

interface deleteSignoutResponse {
}

export const signout = (axios: NuxtAxiosInstance): AxiosPromise<deleteSignoutResponse> => axios.$delete<deleteSignoutResponse>('/api/auth/v1/signout');
