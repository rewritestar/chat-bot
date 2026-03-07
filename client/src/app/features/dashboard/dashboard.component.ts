import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Component, inject } from '@angular/core';
import { AuthService } from '../../core/services/auth.service';
import { environment } from '../../environments/environments';

@Component({
  selector: 'dashboard',
  templateUrl: 'dashboard.html',
})
export class Dashboard {
  http = inject(HttpClient);
  constructor() {
    this.onClick();
  }

  onClick() {
    this.http.get(`${environment.apiUrl}/test`).subscribe();
  }
}
