import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../environments/environments';

@Injectable({
  providedIn: 'root',
})
export class AuthApiService {
  apiUrl: string;
  constructor(private http: HttpClient) {
    this.apiUrl = environment.apiUrl;
  }

  login(reqData: any) {
    return this.http.post(`${this.apiUrl}/login`, reqData);
  }

  signin(reqData: any) {
    return this.http.post(`${this.apiUrl}/signin`, reqData);
  }
}
