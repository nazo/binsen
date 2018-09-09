import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';
import { User } from './types/user';

interface getMeResponse {
    user: User;
}

export const get = (axios: NuxtAxiosInstance): AxiosPromise<getMeResponse> => axios.$get<getMeResponse>('/api/me/v1/get');
