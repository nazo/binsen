import { NuxtHTTPInstance } from '@nuxt/http';
import { Workspace } from './types/workspace';

interface GetWorkspacesResponse {
  workspaces: Workspace[];
}

interface CreateWorkspacesResponse {
  workspaces: Workspace;
}

interface UpdateWorkspacesResponse {
  workspaces: Workspace;
}

export const list = (
  http: NuxtHTTPInstance
): Promise<GetWorkspacesResponse> =>
  http.$get<GetWorkspacesResponse>('/api/workspace/v1/list', {});

export const create = (
  http: NuxtHTTPInstance,
  { name }: { name: string }
): Promise<CreateWorkspacesResponse> =>
  http.$put<CreateWorkspacesResponse>('/api/workspace/v1/create', { name });

export const update = (
  http: NuxtHTTPInstance,
  { id, name }: { id: number, name: string }
): Promise<UpdateWorkspacesResponse> =>
  http.$patch<UpdateWorkspacesResponse>('/api/workspace/v1/update', { id, name });
