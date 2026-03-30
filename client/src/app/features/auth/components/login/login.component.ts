import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';

import { ToastrService } from 'ngx-toastr';

import { AuthApiService } from '../../services/auth-api.service';
import { AuthService } from '../../../../core/services/auth.service';

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
    private authService: AuthService,
    private toastr: ToastrService,
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
      this.toastr.success('성공적으로 로그인되었습니다.');
      this.authService.setToken(res);
      this.router.navigate(['']);
    });
  }

  goToSignin() {
    this.router.navigate(['signin']);
  }
}
