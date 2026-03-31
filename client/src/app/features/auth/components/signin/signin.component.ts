import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { NgClass } from '@angular/common';

import { ToastrService } from 'ngx-toastr';

import { AuthApiService } from '../../services/auth-api.service';
import { AuthService } from '../../../../core/services/auth.service';

@Component({
  selector: 'signin',
  templateUrl: 'signin.html',
  imports: [ReactiveFormsModule, NgClass],
})
export class Signin {
  form = new FormGroup({
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [Validators.required, Validators.minLength(8)]),
  });

  isEmailValid: boolean = true;
  isPasswordValid: boolean = true;

  constructor(
    private authApi: AuthApiService,
    private authService: AuthService,
    private router: Router,
    private toastr: ToastrService,
  ) {
    this.form.get('email')?.statusChanges.subscribe((value) => {
      this.isEmailValid = value === 'VALID' ? true : this.isEmailValid;
    });
    this.form.get('password')?.statusChanges.subscribe((value) => {
      this.isPasswordValid = value === 'VALID' ? true : this.isPasswordValid;
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

    this.authApi.signin(req).subscribe(() => {
      this.toastr.success('성공적으로 가입되었습니다.');

      this.authApi.login(req).subscribe((res: any) => {
        this.authService.setToken(res);
        this.router.navigate(['']);
      });
    });
  }

  goToLogin() {
    this.router.navigate(['login']);
  }
}
