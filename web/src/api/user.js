import request from '@/utils/request'

export function login(data) {
  return request({
    url: '/auth/login',
    method: 'post',
    data
  })
}

export function getInfo(token) {
  return request({
    url: '/auth/info',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/auth/logout',
    method: 'post'
  })
}

export function getList() {
  return request({
    url: '/user',
    method: 'get'
  })
}

export function update(userID, data) {
  return request({
    url: '/user/' + userID,
    method: 'put',
    data
  })
}

export function register(data) {
  return request({
    url: '/auth/register',
    method: 'post',
    data
  })
}

export function destroy(userID) {
  return request({
    url: '/user/' + userID,
    method: 'delete'
  })
}
