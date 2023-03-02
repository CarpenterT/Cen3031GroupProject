import { Component } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent {
  username: string = '';
  password: string = '';
  constructor(private router: Router){

  }
  submit() {
    window.alert('Username is: "' + this.username + '". Password is: "' + this.password + '".')
    this.clear()
    this.router.navigate(['/login'], { queryParams: { username: this.username, password: this.password } });
  }
  clear() {
    //this.username = '';
    //this.password = '';
  }
}
