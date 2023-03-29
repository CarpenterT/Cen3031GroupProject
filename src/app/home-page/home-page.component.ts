import { Component, OnInit } from '@angular/core';
import { ServerService } from '../services/server.service';
import { UserService } from '../services/user.service';
import { Server } from '../models/server.model';
import { User } from '../models/user.model';

@Component({
  selector: 'app-home=page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {
  user: User = {} as User;
  servers: Server[] = [];
  newServerName: string = '';

  constructor(private serverService: ServerService, private userService: UserService) { }

  ngOnInit(): void {
    this.userService.getUser(localStorage.getItem('currentUser') || '').subscribe((user: User) => {
      this.user = user;
      this.serverService.getServers().subscribe((servers: Server[]) => {
        this.servers = servers.filter(server => user.servers.includes(server.id));
      });
    });
  }

  onCreateServer() {
    this.serverService.createServer(this.newServerName, this.user.username);
  }
}
