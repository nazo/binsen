import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';
import { Workspace } from './types/workspace';

interface GetWorkspacesResponse {
  workspaces: Array<Workspace>;
}

interface CreateWorkspacesResponse {
  workspaces: Workspace;
}

export const list = (axios: NuxtAxiosInstance): AxiosPromise<GetWorkspacesResponse> => axios.$get<GetWorkspacesResponse>('/api/workspace/v1/list', { });
export const create = (axios: NuxtAxiosInstance, { name }: { name: string }): AxiosPromise<CreateWorkspacesResponse> => axios.$put<CreateWorkspacesResponse>('/api/workspace/v1/create', { name });
