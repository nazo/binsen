import { User } from '../types/user';
import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';

interface GetGoogleAuthResponse {
  redirectUri: string;
}

interface GetGoogleCallbackResponse {
  user: User;
}

export const auth = (
  axios: NuxtAxiosInstance
): AxiosPromise<GetGoogleAuthResponse> =>
  axios.$get<GetGoogleAuthResponse>('/api/auth/v1/google/auth');

export const callback = (
  axios: NuxtAxiosInstance,
  { code, state }: { code: string; state: string }
): AxiosPromise<GetGoogleCallbackResponse> =>
  axios.$get<GetGoogleCallbackResponse>('/api/auth/v1/google/callback', {
    params: { code, state },
  });
