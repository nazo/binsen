import { NuxtHTTPInstance } from '@nuxt/http';
import { User } from './types/user';

interface GetMeResponse {
  user: User;
}

export const get = (http: NuxtHTTPInstance): Promise<GetMeResponse> =>
  http.$get<GetMeResponse>('/api/me/v1/get');
