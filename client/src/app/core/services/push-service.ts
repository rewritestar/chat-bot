import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../environments/environments';

@Injectable({
  providedIn: 'root',
})
export class PushService {
  constructor(private http: HttpClient) {}
  async enablePush() {
    const permission = await Notification.requestPermission();
    if (permission !== 'granted') return;

    const reg = await navigator.serviceWorker.ready;

    const sub = await reg.pushManager.subscribe({
      userVisibleOnly: true,
      applicationServerKey: environment.pushKey,
    });

    const reqData = {
      endpoint: sub.endpoint,
      p256dh: this.arrayBufferToBase64(sub.getKey('p256dh')),
      auth: this.arrayBufferToBase64(sub.getKey('auth')),
    };

    await this.http.post(`${environment.apiUrl}/push`, reqData).subscribe();
  }

  private arrayBufferToBase64(buffer: ArrayBuffer | null) {
    if (!buffer) return '';
    return btoa(String.fromCharCode(...new Uint8Array(buffer)));
  }
}
