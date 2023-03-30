import { Component } from '@angular/core';
import { MatCard } from '@angular/material/card';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-account-create',
  templateUrl: './account-create.component.html',
  styleUrls: ['./account-create.component.css'],
  
})
export class AccountCreateComponent {
  
  username: string = '';
  password: string = '';

  constructor(private router: Router, private http: HttpClient){

  }



  onCreate(data: {username: string, password: string}){

    // this sends the post request to add a username and password
    //TODO: Need to check if user already exists in DB. Can use GET to check. See main.go and user.go.
    //      If user already exists, need to display that it failed.
    //      If user doesn't exist, need to add using post, then check if it succeeded, then display success.
    this.http.get('http://localhost:8080/users/user/' + this.username).subscribe((res : any) => {
      console.log(res);
      if(res == "User found."){
        //If username is already taken
        alert("Username is already taken.")
        this.clear()

      } else if(res == "User not found."){
        // If username is not found, create the account
        this.http.post('http://localhost:8080/users', data).subscribe((res) => {
          console.log(res);
          if(res == "User successfully created."){
            alert("Account successfully created!")
            this.clear()
          }else{
            alert("Error, try again.")
          }
        });
      }
    })
    


  }

  clear(){
    this.username = '';
    this.password = '';
  }
}
