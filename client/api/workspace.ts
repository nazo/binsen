import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';
import { Workspace } from './types/workspace';

interface getWorkspacesResponse {
  workspaces: Array<Workspace>;
}

interface createWorkspacesResponse {
  workspaces: Workspace;
}

export const list = (axios: NuxtAxiosInstance): AxiosPromise<getWorkspacesResponse> => axios.$get<getWorkspacesResponse>('/api/workspace/v1/list', { });
export const create = (axios: NuxtAxiosInstance, { name }: { name: string }): AxiosPromise<createWorkspacesResponse> => axios.$put<createWorkspacesResponse>('/api/workspace/v1/create', { name });
