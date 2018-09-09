import { Post } from './types/post';
import { NuxtAxiosInstance, AxiosPromise } from '@nuxtjs/axios';

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

export const list = (axios: NuxtAxiosInstance, { workspace_id, page }: { workspace_id: number, page: number }): AxiosPromise<GetPostsResponse> => axios.$get<GetPostsResponse>('/api/post/v1/list', { params: { workspace_id, page } });
export const get = (axios: NuxtAxiosInstance, { id } : { id: number }): AxiosPromise<GetPostResponse> => axios.$get<GetPostResponse>('/api/post/v1/get', { params: { id } });
export const create = (axios: NuxtAxiosInstance, { workspace_id, title, body }: { workspace_id: number, title: string, body: string }): AxiosPromise<CreatePostsResponse> => axios.$put<CreatePostsResponse>('/api/post/v1/create', { workspace_id, title, body });
export const update = (axios: NuxtAxiosInstance, { id, title, body }: { id: number, title: string, body: string }): AxiosPromise<UpdatePostsResponse> => axios.$patch<UpdatePostsResponse>('/api/post/v1/update', { id, title, body });
