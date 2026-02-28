import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';

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

  onSubmit() {
    console.log(this.form.getRawValue());
  }
}
