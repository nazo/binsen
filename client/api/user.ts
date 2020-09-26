import { User } from './types/user';
import { NuxtAxiosInstance } from '@nuxtjs/axios';

interface GetUsersResponse {
  users: Array<User>;
}

interface CreateUserResponse {
  user: User;
}

interface UpdateUserResponse {
  user: User;
}

interface DeleteUserResponse {}

export const list = (
  axios: NuxtAxiosInstance
): Promise<GetUsersResponse> =>
  axios.$get<GetUsersResponse>('/api/user/v1/list', {});

export const create = (
  axios: NuxtAxiosInstance,
  { name, email }: { name: string; email: string }
): Promise<CreateUserResponse> =>
  axios.$put<CreateUserResponse>('/api/user/v1/create', { name, email });

export const update = (
  axios: NuxtAxiosInstance,
  { id, name, email }: { id: number; name: string; email: string }
): Promise<UpdateUserResponse> =>
  axios.$patch<UpdateUserResponse>('/api/user/v1/update', { id, name, email });

export const destroy = (
  axios: NuxtAxiosInstance,
  { id }: { id: number }
): Promise<DeleteUserResponse> =>
  axios.$delete<DeleteUserResponse>('/api/user/v1/delete', { params: { id } });
