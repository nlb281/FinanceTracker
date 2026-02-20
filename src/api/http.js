const BASE_URL = '/api'

export async function request(path, options = {}) {
  const headers = { ...(options.headers || {}) }

  if (options.body && !headers['Content-Type']) {
    headers['Content-Type'] = 'application/json'
  }

  const res = await fetch(`${BASE_URL}${path}`, {
    ...options,
    headers,
  })

  let data = null
  const contentType = res.headers.get('content-type') || ''

  if (res.status !== 204) {
    if (contentType.includes('application/json')) {
      data = await res.json()
    } else {
      const text = await res.text()
      data = text ? { error: text } : null
    }
  }

  if (!res.ok || data?.status === 'error') {
    const err = new Error(data?.error || data?.message || `HTTP ${res.status}`)
    err.status = res.status
    err.data = data
    throw err
  }

  return data
}
