import { Injectable } from '@angular/core';
import moment from 'moment';
import { WsService } from './ws.service';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  public isLoggedIn(): boolean {
    const exp = this.getExpiration();
    return !!exp && moment().isBefore(exp);
  }

  public logout() {
    localStorage.removeItem('authSession');
  }

  getAuthSession() {
    const authSession = localStorage.getItem('authSession');
    if (!authSession) {
      return null;
    }
    return JSON.parse(authSession);
  }

  getToken() {
    const authSession = this.getAuthSession();
    if (!authSession) {
      return null;
    }
    return authSession.token;
  }

  getExpiration() {
    const authSession = this.getAuthSession();
    if (!authSession) {
      return null;
    }
    const exp = authSession.exp;
    return moment(exp);
  }

  getId() {
    const authSession = this.getAuthSession();
    if (!authSession) {
      return null;
    }
    const id = authSession.workerId;
    return id;
  }

  setToken(token: any) {
    const authSession = {
      token: token.token,
      exp: token.exp,
      workerId: token.workerId,
    };
    localStorage.setItem('authSession', JSON.stringify(authSession));
  }
}
