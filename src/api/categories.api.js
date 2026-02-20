import { request } from './http'

export const getCategories = () => request('/categories')
export const createCategories = (payload) =>
  request('/categories', { method: 'POST', body: JSON.stringify(payload) })
export const deleteCategory = (id) => request(`/categories/${id}`, { method: 'DELETE' })
