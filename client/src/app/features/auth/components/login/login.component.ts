import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { AuthApiService } from '../../services/auth-api.service';

@Component({
  selector: 'login',
  templateUrl: 'login.html',
  imports: [ReactiveFormsModule],
})
export class Login {
  form = new FormGroup({
    email: new FormControl('', [Validators.required]),
    password: new FormControl('', [Validators.required]),
  });

  constructor(
    private router: Router,
    private authApi: AuthApiService,
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

    this.authApi.login(req).subscribe((res: any) => {
      this.setToken(res);
      this.router.navigate(['']);
    });
  }

  setToken(token: any) {
    const authSession = {
      token: token.token,
      exp: token.exp,
      workerId: token.workerId,
    };
    localStorage.setItem('authSession', JSON.stringify(authSession));
  }

  goToSignin() {
    this.router.navigate(['signin']);
  }
}
