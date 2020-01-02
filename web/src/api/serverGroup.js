import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/server_group',
    method: 'get',
    params
  })
}

export function destroy(id) {
  return request({
    url: '/server_group/' + id,
    method: 'delete'
  })
}

export function update(id, params) {
  return request({
    url: '/server_group/' + id,
    method: 'put',
    params
  })
}

export function store(params) {
  return request({
    url: '/server_group/store',
    method: 'post',
    params
  })
}



