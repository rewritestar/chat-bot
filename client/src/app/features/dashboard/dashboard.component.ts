import { Component } from '@angular/core';
import { AuthService } from '../../core/services/auth.service';
import { Router } from '@angular/router';

@Component({
  selector: 'dashboard',
  templateUrl: 'dashboard.html',
})
export class Dashboard {
  constructor(
    private authService: AuthService,
    private router: Router,
  ) {
    console.log('ex1');
    if (!this.authService.isLogin()) {
      console.log('ex2');

      this.router.navigate(['login']);
    }
  }
}
