import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from '../models/user.model';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private baseUrl = 'http://localhost:8080'; // change this to server's URL

  constructor(private http: HttpClient) { }

  getUsers(): Observable<User[]> {
    return this.http.get<User[]>(`${this.baseUrl}/users`);
  }

  getUser(username: string): Observable<User> {
    return this.http.get<User>(`${this.baseUrl}/users/${username}`);
  }

  createUser(username: string, password: string): Observable<User> {
    return this.http.post<User>(`${this.baseUrl}/users`, { username, password });
  }

  updateUser(username: string, password: string): Observable<User> {
    return this.http.put<User>(`${this.baseUrl}/users/${username}`, { password });
  }

  deleteUser(username: string): Observable<any> {
    return this.http.delete<any>(`${this.baseUrl}/users/${username}`);
  }
}
