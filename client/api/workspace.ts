import { NuxtHTTPInstance } from '@nuxt/http';
import { Workspace } from './types/workspace';

interface GetWorkspacesResponse {
  workspaces: Workspace[];
}

interface CreateWorkspacesResponse {
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
