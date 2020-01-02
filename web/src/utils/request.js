import axios from 'axios'
import { MessageBox, Message } from 'element-ui'
import store from '@/store'
import { getToken } from '@/utils/auth'
import qs from 'qs'

const service = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  // withCredentials: true,
  timeout: 5000

  // paramsSerializer: function(params) {
  //   return qs.stringify(params, { arrayFormat: 'brackets' })
  // },

  // headers: {
  //   'Content-Type': 'application/x-www-form-urlencoded'
  // }

})

service.interceptors.request.use((request) => {

  if (store.getters.token) {
    request.headers['Authorization'] = 'Bearer ' + getToken()
  }

  if (request.method === 'post') {
    request.data = qs.stringify(request.data)
    request.headers['Content-Type'] = 'application/x-www-form-urlencoded'
  }

  return request
}, function(error) {
  return Promise.reject(error)
})

service.interceptors.response.use(
  response => {
    const res = response.data

    if (res.code !== 200) {
      Message({
        message: res.message || 'Error',
        type: 'error',
        duration: 5 * 1000
      })

      if (res.code === 401) {
        MessageBox.confirm('You have been logged out, you can cancel to stay on this page, or log in again', 'Confirm logout', {
          confirmButtonText: 'Re-Login',
          cancelButtonText: 'Cancel',
          type: 'warning'
        }).then(() => {
          store.dispatch('api/auth/refresh_token').then(() => {
            location.reload()
          })
        })
      }

      return Promise.reject(new Error(res.message || 'Error'))
    } else {
      return res
    }
  },
  error => {
    console.log('err' + error) // for debug

    Message({
      message: error.message,
      type: 'error',
      duration: 5 * 1000
    })

    return Promise.reject(error)
  }
)

export default service
