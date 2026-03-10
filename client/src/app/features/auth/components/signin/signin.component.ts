import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { AuthApiService } from '../../services/auth-api.service';
import { Router } from '@angular/router';

@Component({
  selector: 'signin',
  templateUrl: 'signin.html',
  imports: [ReactiveFormsModule],
})
export class Signin {
  form = new FormGroup({
    email: new FormControl(''),
    password: new FormControl(''),
  });

  constructor(
    private authApi: AuthApiService,
    private router: Router,
  ) {}

  onSubmit() {
    if (this.form.invalid) {
      console.log(this.form.errors);
      return;
    }

    const req = {
      email: this.form.value.email?.trim(),
      password: this.form.value.password?.trim(),
    };

    this.authApi.signin(req).subscribe(() => {
      this.router.navigate(['login']);
    });
  }
}
