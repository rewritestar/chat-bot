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

  indexRoom() {
    return this.http.get(`${this.apiUrl}/rooms`);
  }

  getRoom(id: number) {
    return this.http.get(`${this.apiUrl}/rooms/${id}`);
  }

  saveRoom(data: any) {
    return this.http.post(`${this.apiUrl}/rooms`, data);
  }

  updateRoom(id: number, data: any) {
    return this.http.put(`${this.apiUrl}/rooms/${id}`, data);
  }
}
