import { Post } from './types/post';
import { NuxtHTTPInstance } from '@nuxt/http';

interface GetPostsResponse {
  posts: Post[];
}

interface GetPostResponse {
  post: Post;
}

interface CreatePostsResponse {
  post: Post;
}

interface UpdatePostsResponse {
  post: Post;
}

export const list = (
  http: NuxtHTTPInstance,
  { workspace_id, page }: { workspace_id: number; page: number }
): Promise<GetPostsResponse> =>
  http.$get<GetPostsResponse>('/api/post/v1/list', {
    searchParams: { workspace_id, page },
  });

export const get = (
  http: NuxtHTTPInstance,
  { id }: { id: number }
): Promise<GetPostResponse> =>
  http.$get<GetPostResponse>('/api/post/v1/get', { searchParams: { id } });

export const create = (
  http: NuxtHTTPInstance,
  {
    workspace_id,
    title,
    body,
  }: { workspace_id: number; title: string; body: string }
): Promise<CreatePostsResponse> =>
  http.$put<CreatePostsResponse>('/api/post/v1/create', {
    workspace_id,
    title,
    body,
  });

export const update = (
  http: NuxtHTTPInstance,
  { id, title, body }: { id: number; title: string; body: string }
): Promise<UpdatePostsResponse> =>
  http.$patch<UpdatePostsResponse>('/api/post/v1/update', { id, title, body });
