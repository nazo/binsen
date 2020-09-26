import { User } from '../types/user';
import { NuxtAxiosInstance } from '@nuxtjs/axios';

interface GetGoogleAuthResponse {
  redirectUri: string;
}

interface GetGoogleCallbackResponse {
  user: User;
}

export const auth = (
  axios: NuxtAxiosInstance
): Promise<GetGoogleAuthResponse> =>
  axios.$get<GetGoogleAuthResponse>('/api/auth/v1/google/auth');

export const callback = (
  axios: NuxtAxiosInstance,
  { code, state }: { code: string; state: string }
): Promise<GetGoogleCallbackResponse> =>
  axios.$get<GetGoogleCallbackResponse>('/api/auth/v1/google/callback', {
    params: { code, state },
  });
