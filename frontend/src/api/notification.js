import request from '@/utils/request'

export function getNotifications(query = {}) {
  return request({
    url: '/notifications',
    method: 'get',
    params: query,
  })
}

export function deleteNotification(id) {
  return request({
    url: `/notifications/${id}`,
    method: 'delete',
  })
}

export function clearNotifications() {
  return request({
    url: '/notifications/clear',
    method: 'post',
  })
}

export function getNotificationModules() {
  return request({
    url: '/notifications/modules',
    method: 'get',
  })
}
