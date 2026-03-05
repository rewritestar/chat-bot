import { HttpClient } from '@angular/common/http';
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

  constructor(
    private router: Router,
    private http: HttpClient,
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
    this.http.post('http://localhost:4200/api/login', req).subscribe((res) => {
      console.log(res);
    });
  }

  goToSignin() {
    this.router.navigate(['signin']);
  }
}
