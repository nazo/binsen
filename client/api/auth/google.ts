import { User } from '../types/user';
import { NuxtHTTPInstance } from '@nuxt/http';

interface GetGoogleAuthResponse {
  redirectUri: string;
}

interface GetGoogleCallbackResponse {
  user: User;
}

export const auth = (
  http: NuxtHTTPInstance
): Promise<GetGoogleAuthResponse> =>
  http.$get<GetGoogleAuthResponse>('/api/auth/v1/google/auth');

export const callback = (
  http: NuxtHTTPInstance,
  { code, state }: { code: string; state: string }
): Promise<GetGoogleCallbackResponse> =>
  http.$get<GetGoogleCallbackResponse>('/api/auth/v1/google/callback', {
    searchParams: { code, state },
  });
