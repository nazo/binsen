import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';
import { Group } from './types/group';

interface getGroupsResponse {
  groups: Array<Group>;
}

interface createGroupResponse {
  group: Group;
}

interface updateGroupResponse {
  group: Group;
}

interface deleteGroupResponse {
}

export const list = (axios: NuxtAxiosInstance): AxiosPromise<getGroupsResponse> => axios.$get<getGroupsResponse>('/api/group/v1/list', { });
export const create = (axios: NuxtAxiosInstance, { name }: { name: string }): AxiosPromise<createGroupResponse> => axios.$put<createGroupResponse>('/api/group/v1/create', { name });
export const update = (axios: NuxtAxiosInstance, { id, name }: { id: number, name: string }): AxiosPromise<updateGroupResponse> => axios.$patch<updateGroupResponse>('/api/group/v1/update', { id, name });
export const destroy = (axios: NuxtAxiosInstance, { id }: { id: number }): AxiosPromise<deleteGroupResponse> => axios.$delete<deleteGroupResponse>('/api/group/v1/delete', { params: { id } });