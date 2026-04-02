import { Injectable } from '@angular/core';
import moment from 'moment';
import { map, Subject } from 'rxjs';
import { environment } from '../../environments/environments';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class WsService {
  private socket!: WebSocket;
  private messageSubject = new Subject<MessageType>();
  private messages$ = this.messageSubject.asObservable();
  private wsUrl = environment.wsUrl;
  private retry = 0;

  constructor(private authService: AuthService) {}

  connect() {
    if (this.socket && this.socket.readyState === this.socket.OPEN) {
      return;
    }

    this.socket = new WebSocket(this.wsUrl);

    this.socket.onopen = () => this.onOpen();

    this.socket.onmessage = (event) => {
      const data: MessageType = JSON.parse(event.data);
      switch (data.type) {
        case environment.messageType.chat:
          this.messageSubject.next(data);
          break;
        default:
      }
    };

    this.socket.onclose = () => {
      console.log('WebSocket closed.');
      this.reconnect();
    };

    this.socket.onerror = (err) => {
      console.error('WebSocket error', err);
    };
  }

  sendMessage(reqData: any) {
    const req = {
      type: environment.messageType.chat,
      data: reqData,
    };
    this.socket.send(JSON.stringify(req));
  }

  getMessage() {
    return this.messages$;
  }

  joinRoom(roomId: number) {
    if (this.socket && this.socket.readyState === this.socket.OPEN) {
      const req = {
        type: environment.messageType.join,
        data: {
          roomId: roomId,
        },
      };
      this.socket.send(JSON.stringify(req));
    }
  }

  leaveRoom() {
    if (this.socket && this.socket.readyState === this.socket.OPEN) {
      const req = {
        type: environment.messageType.leave,
        data: null,
      };
      this.socket.send(JSON.stringify(req));
    }
  }

  close() {
    if (this.socket) {
      this.socket.close();
    }
  }

  private onOpen() {
    const req = {
      type: environment.messageType.auth,
      data: {
        token: this.authService.getToken(),
      },
    };
    this.socket.send(JSON.stringify(req));
    this.retry = 0;
    console.log('WebSocket connected');
  }

  private reconnect() {
    if (!this.authService.isLoggedIn()) return;

    const baseDelay = Math.min(1000 * Math.pow(2, this.retry), 30000);
    const jitter = Math.random() * 1000;
    const delay = baseDelay + jitter;

    this.retry++;
    setTimeout(() => this.connect(), delay);
    console.log('WebSocket reconnecting. retry: ', this.retry);
  }
}

type MessageType = {
  type: string;
  roomId: number;
  content: string;
};
