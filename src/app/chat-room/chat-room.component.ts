import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-chat-room',
  templateUrl: './chat-room.component.html',
  styleUrls: ['./chat-room.component.css']
})
export class ChatRoomComponent {
    messages: string[] = [];
    message: string = '';
    users: string[] = [];
    oneUser: string = '';
    whenCreateds: string[] = [];
    whenCreated: string = '';

    constructor(private router: Router, private http: HttpClient) { }
    
    // Request a list of all messages
    ngOnInit() {
      this.updateMessages();
    }

    updateMessages(){
      this.http.get('http://localhost:8080/chat').subscribe((res: any) => {
        console.log(res);
        let i = 0;
        // This for loop uses the array from the GET to load the arrays below.
        // These are then drawn on the page.
        for (var entry of res){
          this.whenCreateds[i] = entry.CreatedAt.substring(0,19);
          this.messages[i] = entry.msg;
          this.users[i] = entry.username;
          i++;
        }
      })
    }

    sendMessage() {
      // Form the struct for the post request.
      let data: {msg: string; username: string} = {
        msg: this.message,
        username: localStorage.getItem('currentUser') || ''
      }

      // Send the post, containing the username and message
      this.http.post('http://localhost:8080/chat', data).subscribe((res : any) => {
        if(res == "Message sent."){
            console.log("Message: " + this.message + ", was sent!");
            this.message = '';
          }else{
            console.log("Failed to send message!");
            alert("Message failed to send.")
          }
          // Refresh the messages
          this.updateMessages();
      })
    }

    editMessage(when: string, who: string, what: string){
      // this is awkward
      let data: {when1: string, who1: string, what1: string} = {
        when1: when,
        who1: who,
        what1: what
      }
      console.log(data);
      // First get the message from the server using the time, username,
      // and message.
      //this.http.get('http://localhost:8080/chat/find/' + data).subscribe((res: any) => {


      //})
      
    }
}
