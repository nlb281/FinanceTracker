import { request } from './http'

export const getTransactions = () => request('/transactions')
export const createTransaction = (payload) =>
  request('/transactions', { method: 'POST', body: JSON.stringify(payload) })
export const deleteTransaction = (id) => request(`/transactions/${id}`, { method: 'DELETE' })
