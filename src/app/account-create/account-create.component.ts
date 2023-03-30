import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login-page',
  templateUrl: './account-create.component.html',
  styleUrls: ['./account-create.component.css'],
  
})
export class AccountCreateComponent {
  
  username: string = '';
  password: string = '';

  constructor(private router: Router, private http: HttpClient){ }
  
  onCreate(data: {username: string, password: string}){
    //console.log(data);
    // this sends the post request to add a username and password
    //TODO: Need to check if user already exists in DB. Can use GET to check. See main.go and user.go.
    //      If user already exists, need to display that it failed.
    //      If user doesn't exist, need to add using post, then check if it succeeded, then display success.
    this.http.post('http://localhost:8080/users', data).subscribe((res) => {
      console.log(res);
    });
    localStorage.setItem('currentUser', data.username);
    this.router.navigate(['/home']);
  }

}
