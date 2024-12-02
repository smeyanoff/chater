class WebSocketClient {
  private socket: WebSocket | null = null;
  private readonly messageListeners: ((message: unknown) => void)[] = [];

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
        const message = JSON.parse(event.data) as unknown
        this.messageListeners.forEach(listener => listener(message)) // Вызываем все подписанные колбэки
      }
    })
  }

  // Добавляем возможность подписки на получение сообщений
  onMessage (listener: (message: unknown) => void): void {
    this.clearMessageListeners() // Очищаем все предыдущие подписки
    this.messageListeners.push(listener)
  }

  // Очистка всех подписчиков сообщений
  private clearMessageListeners (): void {
    this.messageListeners.length = 0
  }

  // Отправка сообщения через WebSocket
  send (data: Record<string, unknown>): void {
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
