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
  private wsUrl = '';
  private token = '';

  constructor(private authService: AuthService) {
    this.wsUrl = environment.wsUrl;
    this.token = this.authService.getToken();
  }

  connect() {
    if (this.socket && this.socket.readyState === this.socket.OPEN) {
      return;
    }
    this.socket = new WebSocket(this.wsUrl);

    this.socket.onopen = () => this.onOpen();

    this.socket.onmessage = (event) => {
      const data: MessageType = JSON.parse(event.data);
      this.messageSubject.next(data);
    };

    this.socket.onclose = () => {
      console.log('WebSocket closed');
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

  close() {
    this.socket.close();
  }

  private onOpen() {
    const req = {
      type: environment.messageType.auth,
      data: {
        token: this.token,
      },
    };
    if (this.socket.readyState === this.socket.OPEN) {
      this.socket.send(JSON.stringify(req));
      console.log('WebSocket connected');
    }
  }
}

type MessageType = {
  roomId: number;
  content: string;
};
