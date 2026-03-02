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
    this.http.get('http://localhost:4200/api/test').subscribe((res) => {
      console.log(res);
    });
    this.form.getRawValue();
  }

  goToSignin() {
    this.router.navigate(['signin']);
  }
}
