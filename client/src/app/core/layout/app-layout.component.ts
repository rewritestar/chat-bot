import { WsService } from './../services/ws.service';
import { Component } from '@angular/core';
import { Router, RouterOutlet } from '@angular/router';

import { AuthService } from './../services/auth.service';

@Component({
  selector: 'app-layout',
  templateUrl: 'app-layout.html',
  imports: [RouterOutlet],
})
export class AppLayout {
  constructor(
    private wsService: WsService,
    private authService: AuthService,
    private router: Router,
  ) {}
  logout() {
    this.wsService.close();
    this.authService.logout();
    this.router.navigate(['login']);
  }
}
