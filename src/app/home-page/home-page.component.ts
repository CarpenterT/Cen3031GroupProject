import { Component, OnInit } from '@angular/core';
import { ServerService } from '../services/server.service';
import { UserService } from '../services/user.service';
import { Server } from '../models/server.model';
import { User } from '../models/user.model';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';

@Component({
  selector: 'app-home=page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {
  //user: User = {} as User;
  username: string = '';
  //servers: Server[] = [];
  serverIDs: number[] = [];
  newServerName: string = '';
  serverName: Record<number, string> = {};

  constructor(private router: Router, private http: HttpClient, private serverService: ServerService, private userService: UserService) { }

  ngOnInit(): void {
    if (!localStorage.getItem('currentUser') || localStorage.getItem('currentUser') == '') {
      this.router.navigate(['/']);
    } else {
      this.username = localStorage.getItem('currentUser') || '';
      this.http.get('http://localhost:8080/users/user/' + this.username + '/servers').subscribe((serversString: any) => {
        this.serverIDs = serversString.split(',').map(Number);
      })
      for (var serverID in this.serverIDs) {
        this.http.get('http://localhost:8080/servers/' + serverID).subscribe((serverName: any) => {
          this.serverName[serverID] = serverName;
        })
      }
    }
    // TODO: change so that user is retrieved as an object. Need list of servers that user is a member of.
    /*
    this.userService.getUser(localStorage.getItem('currentUser') || '').subscribe((user: User) => {
      this.user = user;
      this.serverService.getServers().subscribe((servers: Server[]) => {
        this.servers = servers.filter(server => user.servers.includes(server.id));
      });
    });
    */
  }

  onCreateServer() {
    //this.serverService.createServer(this.newServerName, this.username);
    this.http.post('http://localhost:8080/server', { name: this.newServerName, admin: this.username }).subscribe((res) => {
      console.log(res);
    })
    this.newServerName = '';
    this.ngOnInit();
  }

  onLogout() {
    localStorage.setItem('currentUser', '');
    this.router.navigate(['/']);
  }
}
