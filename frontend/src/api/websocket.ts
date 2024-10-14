// websocket.ts
class WebSocketClient {
    private socket: WebSocket | null = null;

    // Подключение к WebSocket
    connect (url: string): Promise<void> {
      return new Promise((resolve, reject) => {
        this.socket = new WebSocket(url)

        this.socket.onopen = () => {
          console.log('WebSocket соединение установлено')
          resolve()
        }

        this.socket.onerror = (event) => {
          const errorMessage = event instanceof ErrorEvent ? event.message : 'Ошибка WebSocket'
          reject(new Error(errorMessage))
        }

        this.socket.onclose = () => {
          console.log('WebSocket соединение закрыто')
        }
      })
    }

    // Отправка сообщения и получение ответа с универсальным типом (дженерик T)
    send<T> (data: any): Promise<T> {
      return new Promise((resolve, reject) => {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
          this.socket.send(JSON.stringify(data))

          // Ожидание ответа
          this.socket.onmessage = (event) => {
            try {
              const response: T = JSON.parse(event.data)
              resolve(response) // Возвращаем ответ любого типа
            } catch (error) {
              reject(new Error('Ошибка при разборе сообщения: ' + error))
            }
          }
        } else {
          reject(new Error('WebSocket не подключен или закрыт'))
        }
      })
    }

    // Закрытие WebSocket соединения
    close (): void {
      if (this.socket) {
        this.socket.close()
        this.socket = null // Очищаем ссылку на WebSocket
      }
    }

    // Проверка статуса соединения
    isConnected (): boolean {
      return this.socket !== null && this.socket.readyState === WebSocket.OPEN
    }
}

export const webSocketClient = new WebSocketClient()
