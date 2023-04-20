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

    constructor(private router: Router, private http: HttpClient) { }
    
    // Request a list of all messages
    ngOnInit() {
      this.updateMessages();
    }

    updateMessages(){
      this.http.get('http://localhost:8080/chat').subscribe((res: any) => {
        console.log(res);
        let i = 0;
        for (var entry of res){
          this.messages[i] = entry.username + ": " + entry.msg;
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
}
