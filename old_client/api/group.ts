import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';
import { Group } from './types/group';

interface GetGroupsResponse {
  groups: Array<Group>;
}

interface CreateGroupResponse {
  group: Group;
}

interface UpdateGroupResponse {
  group: Group;
}

interface DeleteGroupResponse {}

export const list = (
  axios: NuxtAxiosInstance
): AxiosPromise<GetGroupsResponse> =>
  axios.$get<GetGroupsResponse>('/api/group/v1/list', {});

export const create = (
  axios: NuxtAxiosInstance,
  { name }: { name: string }
): AxiosPromise<CreateGroupResponse> =>
  axios.$put<CreateGroupResponse>('/api/group/v1/create', { name });

export const update = (
  axios: NuxtAxiosInstance,
  { id, name }: { id: number; name: string }
): AxiosPromise<UpdateGroupResponse> =>
  axios.$patch<UpdateGroupResponse>('/api/group/v1/update', { id, name });

export const destroy = (
  axios: NuxtAxiosInstance,
  { id }: { id: number }
): AxiosPromise<DeleteGroupResponse> =>
  axios.$delete<DeleteGroupResponse>('/api/group/v1/delete', {
    params: { id },
  });
