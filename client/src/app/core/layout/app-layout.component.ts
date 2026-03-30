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
  ) {
    this.wsService.connect();
  }

  logout() {
    this.authService.logout();
    this.wsService.close();
    this.router.navigate(['login']);
  }
}
