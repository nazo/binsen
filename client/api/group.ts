import { NuxtHTTPInstance } from '@nuxt/http';
import { Group } from './types/group';

interface GetGroupsResponse {
  groups: Group[];
}

interface CreateGroupResponse {
  group: Group;
}

interface UpdateGroupResponse {
  group: Group;
}

interface DeleteGroupResponse {}

export const list = (http: NuxtHTTPInstance): Promise<GetGroupsResponse> => http.$get<GetGroupsResponse>('/api/group/v1/list', {});

export const create = (http: NuxtHTTPInstance, { name }: { name: string }): Promise<CreateGroupResponse> => http.$put<CreateGroupResponse>('/api/group/v1/create', { name });

export const update = (http: NuxtHTTPInstance, { id, name }: { id: number; name: string }): Promise<UpdateGroupResponse> => http.$patch<UpdateGroupResponse>('/api/group/v1/update', { id, name });

export const destroy = (http: NuxtHTTPInstance, { id }: { id: number }): Promise<DeleteGroupResponse> =>
  http.$delete<DeleteGroupResponse>('/api/group/v1/delete', {
    searchParams: { id }
  });
