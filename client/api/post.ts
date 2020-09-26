import { Post } from './types/post';
import { NuxtAxiosInstance } from '@nuxtjs/axios';

interface GetPostsResponse {
  posts: Array<Post>;
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
  axios: NuxtAxiosInstance,
  { workspace_id, page }: { workspace_id: number; page: number }
): Promise<GetPostsResponse> =>
  axios.$get<GetPostsResponse>('/api/post/v1/list', {
    params: { workspace_id, page },
  });

export const get = (
  axios: NuxtAxiosInstance,
  { id }: { id: number }
): Promise<GetPostResponse> =>
  axios.$get<GetPostResponse>('/api/post/v1/get', { params: { id } });

export const create = (
  axios: NuxtAxiosInstance,
  {
    workspace_id,
    title,
    body,
  }: { workspace_id: number; title: string; body: string }
): Promise<CreatePostsResponse> =>
  axios.$put<CreatePostsResponse>('/api/post/v1/create', {
    workspace_id,
    title,
    body,
  });

export const update = (
  axios: NuxtAxiosInstance,
  { id, title, body }: { id: number; title: string; body: string }
): Promise<UpdatePostsResponse> =>
  axios.$patch<UpdatePostsResponse>('/api/post/v1/update', { id, title, body });
