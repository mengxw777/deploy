import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/server',
    method: 'get',
    params
  })
}

export function show(id) {
  return request({
    url: '/server/' + id,
    method: 'get'
  })
}

export function store(params) {
  return request({
    url: '/server/store',
    method: 'post',
    params
  })
}

export function update(id, params) {
  return request({
    url: '/server/' + id,
    method: 'put',
    params
  })
}

export function destroy(id) {
  return request({
    url: '/server/' + id,
    method: 'delete',
  })
}



