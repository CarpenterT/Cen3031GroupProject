import { Component } from '@angular/core';

@Component({
  selector: 'app-account-create',
  templateUrl: './account-create.component.html',
  styleUrls: ['./account-create.component.css']
})
export class AccountCreateComponent {
  username: string = '';
  password: string = '';
  submit() {
    window.alert('Username is: "' + this.username + '". Password is: "' + this.password + '".')
    localStorage.setItem('currentUser', this.username);
    this.clear()
  }
  clear() {
    this.username = '';
    this.password = '';
  }
}

