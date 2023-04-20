import { Component } from '@angular/core';
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

  constructor(private router: Router, private http: HttpClient) { }
  
  onCreate(data: {username: string, password: string}){
    // This function first checks if the username exists in the DB. If it does,
    // then we check if the password they provided matches the one in the DB.
    // You can check console to see status after submitting.
    // For more, see CheckPass() in user.go.
    
    this.http.get('http://localhost:8080/users/user/' + this.username).subscribe((res : any) => {
      console.log(res);
      if(res == "User found."){
        // If user was found, send another request with the password
        this.http.post('http://localhost:8080/users/user', data).subscribe((response) => {
          if(response == "Password validated."){
            //Correct combo
            console.log("Correct username and password!")
            alert("Login successful!")
            
            //create cookie and then navigate home
            localStorage.setItem('currentUser', data.username);
            this.router.navigate(['/home']);
          }else if(response == "Invalid."){
            //DB threw ErrRecordNotFound
            console.log("Incorrect password!")
            alert("Login failed: wrong username or password!")
          }else{
            //hopefully never reach here.
            console.log("Unknown Error!")
          }
        })
      }else{
        // If user was not found
        console.log("User does not exist!")
        alert("Login failed: username does not exist!")
      }
    })
    
/*
    this.http.post('http://localhost:8080/users', data).subscribe((res) => {
      console.log(res);
    });
    */
  }


}
