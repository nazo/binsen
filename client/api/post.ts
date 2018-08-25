export const list = (axios: any, workspace_id: number, page: number): any => axios.$get('/api/post/v1/list', { params: { workspace_id, page } });
export const get = (axios: any, id: number): any => axios.$get('/api/post/v1/get', { params: { id } });
export const create = (axios: any, workspace_id: number, title: string, body: string): any => axios.$put('/api/post/v1/create', { workspace_id, title, body });
export const update = (axios: any, id: number, title: string, body: string): any => axios.$patch('/api/post/v1/update', { id, title, body });
