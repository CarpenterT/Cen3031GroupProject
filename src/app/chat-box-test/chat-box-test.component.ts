import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';

@Component({
  selector: 'app-chat-box-test',
  templateUrl: './chat-box-test.component.html',
  styleUrls: ['./chat-box-test.component.css']
})
export class ChatBoxTestComponent {
  username = 'username';
  message = '';
  messages = [];

  constructor(private http: HttpClient) {
  }
  
  submit(): void {
    this.http.post('http://localhost:8000/api/messages', {
      username: this.username,
      message: this.message
    }).subscribe(() => this.message = '');
  }

}
