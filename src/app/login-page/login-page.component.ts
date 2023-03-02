import { Component } from '@angular/core';
import { MatCard } from '@angular/material/card';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css'],
  
})
export class LoginPageComponent {
  
  username: string = '';
  password: string = '';

  constructor(private router: Router){

  }

  
  submit() {
    //The alert menu does not update properly in Cypress. Works in normal operation though.
    window.alert('Username is: "' + this.username + '". Password is: "' + this.password + '".')
    //The below line causes a huge error in Cypress.
    this.router.navigate(['/login'], { queryParams: { username: this.username, password: this.password } });

    this.clear()
    
  }
  clear() {
    this.username = '';
    this.password = '';
  }
}
