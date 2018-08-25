export const list = (axios): any => axios.$get('/api/group/v1/list', { });
export const get = (axios, { id }): any => axios.$get('/api/group/v1/get', { params: { id } });
export const create = (axios, { name, email }): any => axios.$put('/api/group/v1/create', { name, email });
export const update = (axios, { id, name, email }): any => axios.$patch('/api/group/v1/update', { id, name, email });
export const destroy = (axios, { id }): any => axios.$delete('/api/group/v1/delete', { params: { id } });
