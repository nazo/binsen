export const auth = (axios): any => axios.$get('/api/auth/v1/google/auth');
export const callback = (axios, code, state): any => axios.$get('/api/auth/v1/google/callback', { params: { code, state } });
