import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'dashboard',
  template: 'dashboard',
})
export class Dashboard {
  constructor(private router: Router) {
    this.router.navigate(['room']);
  }
}
