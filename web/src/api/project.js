import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/project',
    method: 'get',
    params
  })
}

export function store(params) {
  return request({
    url: '/project/store',
    method: 'post',
    params
  })
}

export function show(id) {
  return request({
    url: '/project/' + id,
    method: 'get'
  })
}

export function update(id, params) {
  return request({
    url: '/project/' + id,
    method: 'put',
    params
  })
}

export function destroy(id) {
  return request({
    url: '/api/project/' + id,
    method: 'delete'
  })
}
