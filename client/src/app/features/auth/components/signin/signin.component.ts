import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { Router } from '@angular/router';

import { ToastrService } from 'ngx-toastr';

import { AuthApiService } from '../../services/auth-api.service';
import { AuthService } from '../../../../core/services/auth.service';

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
    private authService: AuthService,
    private router: Router,
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

    this.authApi.signin(req).subscribe(() => {
      this.toastr.success('성공적으로 가입되었습니다.');

      this.authApi.login(req).subscribe((res: any) => {
        this.authService.setToken(res);
        this.router.navigate(['']);
      });
    });
  }
}
