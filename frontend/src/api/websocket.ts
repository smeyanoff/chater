// websocket.ts
class WebSocketClient {
  private socket: WebSocket | null = null;
  private messageListeners: ((message: any) => void)[] = [];

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

      // Подписка на входящие сообщения
      this.socket.onmessage = (event) => {
        const message = JSON.parse(event.data)
        this.messageListeners.forEach(listener => listener(message)) // Вызываем все подписанные колбэки
      }
    })
  }

  // Добавляем возможность подписки на получение сообщений
  onMessage (listener: (message: any) => void): void {
    this.messageListeners.push(listener)
  }

  // Отправка сообщения через WebSocket
  send (data: any): void {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      this.socket.send(JSON.stringify(data))
    } else {
      console.error('WebSocket не подключен')
    }
  }

  // Закрытие соединения
  close (): void {
    if (this.socket) {
      this.socket.close()
      this.socket = null
    }
  }

  // Проверка статуса соединения
  isConnected (): boolean {
    return this.socket !== null && this.socket.readyState === WebSocket.OPEN
  }
}

export const webSocketClient = new WebSocketClient()
