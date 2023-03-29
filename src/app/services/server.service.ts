import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Server } from '../models/server.model';

@Injectable({
  providedIn: 'root'
})
export class ServerService {
  private baseUrl = 'http://localhost:8080'; // change this to server's URL

  constructor(private http: HttpClient) { }

  getServers(): Observable<Server[]> {
    return this.http.get<Server[]>(`${this.baseUrl}/servers`);
  }

  getServer(serverId: string): Observable<Server> {
    return this.http.get<Server>(`${this.baseUrl}/servers/${serverId}`);
  }

  createServer(name: string, admin: string): Observable<Server> {
    return this.http.post<Server>(`${this.baseUrl}/servers`, { name, admin });
  }

  addMember(serverId: string, username: string): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/servers/${serverId}/addMember`, { username });
  }

  removeMember(serverId: string, username: string): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/servers/${serverId}/removeMember`, { username });
  }

  sendMessage(serverId: string, username: string, message: string): Observable<any> {
    return this.http.post<any>(`${this.baseUrl}/servers/${serverId}/sendMessage`, { username, message });
  }
}
