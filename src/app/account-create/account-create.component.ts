import { Component } from '@angular/core';
import { MatCard } from '@angular/material/card';
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

  constructor(private router: Router, private http: HttpClient){

  }



  onCreate(data: {username: string, password: string}){
    //console.log(data);
    // this sends the post request to add a username and password
    this.http.post('http://localhost:8080/users', data).subscribe((res) => {
      console.log(res);
    });
  }


}
