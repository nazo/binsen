import { User } from './types/user';
import { NuxtHTTPInstance } from '@nuxt/http';

interface GetUsersResponse {
  users: User[];
}

interface CreateUserResponse {
  user: User;
}

interface UpdateUserResponse {
  user: User;
}

interface DeleteUserResponse {}

export const list = (
  http: NuxtHTTPInstance
): Promise<GetUsersResponse> =>
  http.$get<GetUsersResponse>('/api/user/v1/list', {});

export const create = (
  http: NuxtHTTPInstance,
  { name, email }: { name: string; email: string }
): Promise<CreateUserResponse> =>
  http.$put<CreateUserResponse>('/api/user/v1/create', { name, email });

export const update = (
  http: NuxtHTTPInstance,
  { id, name, email }: { id: number; name: string; email: string }
): Promise<UpdateUserResponse> =>
  http.$patch<UpdateUserResponse>('/api/user/v1/update', { id, name, email });

export const destroy = (
  http: NuxtHTTPInstance,
  { id }: { id: number }
): Promise<DeleteUserResponse> =>
  http.$delete<DeleteUserResponse>('/api/user/v1/delete', { searchParams: { id } });
