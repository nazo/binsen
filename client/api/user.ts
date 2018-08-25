export const list = (axios): any => axios.$get('/api/user/v1/list', { });
export const get = (axios, { id }): any => axios.$get('/api/user/v1/get', { params: { id } });
export const create = (axios, { name, email }): any => axios.$put('/api/user/v1/create', { name, email });
export const update = (axios, { id, name, email }): any => axios.$patch('/api/user/v1/update', { id, name, email });
export const destroy = (axios, { id }): any => axios.$delete('/api/user/v1/delete', { params: { id } });
