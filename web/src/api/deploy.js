import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/deploy',
    method: 'get',
    params
  })
}

export function show(id) {
  return request({
    url: '/deploy/' + id,
    method: 'get'
  })
}

export function build(id) {
  return request({
    url: '/deploy/build/' + id,
    method: 'post'
  })
}

export function deploy(id) {
  return request({
    url: '/deploy/deploy/' + id,
    method: 'post'
  })
}

export function store(params) {
  return request({
    url: '/deploy/store',
    method: 'post',
    params
  })
}

export function destroy(id) {
  return request({
    url: '/deploy/' + id,
    method: 'delete'
  })
}


