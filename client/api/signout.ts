import { NuxtHTTPInstance } from '@nuxt/http';

interface DeleteSignoutResponse {}

export const signout = (http: NuxtHTTPInstance): Promise<DeleteSignoutResponse> => http.$delete<DeleteSignoutResponse>('/api/auth/v1/signout');
