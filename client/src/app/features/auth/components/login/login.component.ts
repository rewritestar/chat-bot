import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { NgClass } from '@angular/common';

import { ToastrService } from 'ngx-toastr';

import { AuthApiService } from '../../services/auth-api.service';
import { AuthService } from '../../../../core/services/auth.service';

@Component({
  selector: 'login',
  templateUrl: 'login.html',
  imports: [ReactiveFormsModule, NgClass],
})
export class Login {
  form = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(8)]),
  });

  isEmailValid: boolean = true;
  isPasswordValid: boolean = true;

  constructor(
    private router: Router,
    private authApi: AuthApiService,
    private authService: AuthService,
    private toastr: ToastrService,
  ) {
    this.form.get('email')?.statusChanges.subscribe((value) => {
      this.isEmailValid = value === 'VALID' || this.isEmailValid;
    });
    this.form.get('password')?.statusChanges.subscribe((value) => {
      this.isPasswordValid = value === 'VALID' || this.isPasswordValid;
    });
  }

  onSubmit() {
    if (this.form.invalid) {
      console.log(this.form.get('email')?.errors, this.form.get('password')?.errors);
      this.isEmailValid = this.form.get('email')?.valid || false;
      this.isPasswordValid = this.form.get('password')?.valid || false;
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
