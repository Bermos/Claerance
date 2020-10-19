import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from 'rxjs';
import { User } from "./user.struct";

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  public getUser(id: number): Observable<User> {
    return this.http.get<User>(`api/user/${id}`);
  }

  public getAllUsers(): Observable<User[]> {
    return this.http.get<User[]>('api/user')
  }

  public updateUser(user: User): Observable<any> {
    return this.http.put(`api/user/${user.id}`, user)
  }
}
