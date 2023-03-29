import { Component } from '@angular/core';
import { MatCard } from '@angular/material/card';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css'],
  
})
export class LoginPageComponent {
  
  username: string = '';
  password: string = '';

  constructor(private router: Router, private http: HttpClient){

  }



  onCreate(data: {username: string, password: string}){
    //console.log(data);
    // We need a url to send the POST to. Something through RESTapi or whatever.
    this.http.post('http://localhost:8080/users', data).subscribe((res) => {
      console.log(res);
    });
  }


}
