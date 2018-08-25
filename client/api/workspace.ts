export const list = (axios): any => axios.$get('/api/workspace/v1/list', { });
export const get = (axios, { id }): any => axios.$get('/api/workspace/v1/get', { params: { id } });
export const create = (axios, { name, email }): any => axios.$put('/api/workspace/v1/create', { name, email });
export const update = (axios, { id, name, email }): any => axios.$patch('/api/workspace/v1/update', { id, name, email });
export const destroy = (axios, { id }): any => axios.$delete('/api/workspace/v1/delete', { params: { id } });
