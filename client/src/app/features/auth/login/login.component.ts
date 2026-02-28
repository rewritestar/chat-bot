import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'login',
  templateUrl: 'login.html',
  imports: [ReactiveFormsModule],
})
export class Login {
  form = new FormGroup({
    email: new FormControl(''),
    password: new FormControl(''),
  });

  constructor(private router: Router) {}

  onSubmit() {
    this.form.getRawValue();
  }

  goToSignin() {
    this.router.navigate(['signin']);
  }
}
