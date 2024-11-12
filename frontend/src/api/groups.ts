// src/api/auth.ts
import apiClient from './axios'
import { Group } from '@/types'
import { GroupsResponse, SuccessResponse } from './responses'

export interface UserGroupRequest {
    userID: number;
  }

export const getAllUserGroups = async (): Promise<Group[]> => {
  try {
    const { data } = await apiClient.get<GroupsResponse>('/v1/groups', {
      withCredentials: true // Включаем отправку cookie
    })
    return data.groups
  } catch (error) {
    return Promise.reject(new Error(
            `Ошибка получения групп: ${error instanceof Error ? error.message : error}`
    ))
  }
}

export const createGroup = async (groupName: string): Promise<Group> => {
  try {
    const { data } = await apiClient.post<Group>('/v1/groups', { groupName }, {
      withCredentials: true // Включаем отправку cookie
    })
    return data
  } catch (error) {
    return Promise.reject(new Error(
            `Ошибка создания группы: ${error instanceof Error ? error.message : error}`
    ))
  }
}

export const deleteGroup = async (groupID: number): Promise<SuccessResponse> => {
  try {
    const response = await apiClient.delete(
            `/v1/groups/${groupID}`, {
              withCredentials: true // Включаем отправку cookie
            })
    return response.data as SuccessResponse
  } catch (error) {
    // Преобразуем ошибку в объект Error, чтобы удовлетворить SonarLint
    return Promise.reject(new Error(
            `Ошибка при удалении группы: ${error instanceof Error ? error.message : error}`
    ))
  }
}

export const addUserToGroup = async (
  groupID: string,
  addUserToGroupRequest: UserGroupRequest
): Promise<SuccessResponse> => {
  try {
    // Отправляем POST-запрос на добавление пользователя в группу
    const response = await apiClient.post(`/v1/groups/${groupID}/users`, addUserToGroupRequest)

    // Если запрос успешен (статус 200), возвращаем данные ответа
    return response.data as SuccessResponse
  } catch (error) {
    return Promise.reject(new Error(
            `Ошибка при добавлении пользователя в группу: ${error instanceof Error ? error.message : error}`
    ))
  }
}

export const removeUserFromGroup = async (
  groupID: string,
  userGroupRequest: UserGroupRequest
): Promise<SuccessResponse> => {
  try {
    const response = await apiClient.delete(`/v1/groups/${groupID}/users`, {
      data: userGroupRequest,
      withCredentials: true // Включаем отправку cookie
    })
    return response.data as SuccessResponse
  } catch (error) {
    return Promise.reject(new Error(
            `Ошибка при удалении пользователя из группы: ${error instanceof Error ? error.message : error}`
    ))
  }
}
