import { Component } from '@angular/core';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
  username: string = '';
  password: string = '';

  submit() {
    window.alert('Username is: "' + this.username + '". Password is: "' + this.password + '".')
    this.clear()
  }
  clear() {
    this.username = '';
    this.password = '';
  }
}
