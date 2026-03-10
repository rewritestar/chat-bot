import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environments';

@Injectable({
  providedIn: 'root',
})
export class RoomApiService {
  apiUrl: string;
  constructor(private http: HttpClient) {
    this.apiUrl = environment.apiUrl;
  }

  getRoom() {
    return this.http.get(`${this.apiUrl}/rooms`);
  }

  saveChat(reqData: any) {
    return this.http.post(`${this.apiUrl}/chat`, reqData);
  }
}
