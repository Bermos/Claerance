import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { Observable } from 'rxjs';
import { User } from "./user.struct";

@Injectable({
  providedIn: 'root'
})
export class UserService {

  constructor(private http: HttpClient) { }

  /**
   * Retrieves the user details for the requested user.
   * @param id of the requested user
   */
  public getUser(id: number): Observable<User> {
    return this.http.get<User>(`api/user/${id}`);
  }

  /**
   * Retrieves details of all users.
   */
  public getAllUsers(): Observable<User[]> {
    return this.http.get<User[]>('api/user/list')
  }

  /**
   * Updates the user with the new details.
   * @param user to be updated with updated attributes
   */
  public updateUser(user: User): Observable<any> {
    return this.http.put(`api/user/${user.id}`, user)
  }

  /**
   * Sends a deletion request to the server
   * @param id of the user to be deleted
   */
  public deleteUser(id: number): Observable<any> {
    return this.http.delete(`api/user/${id}`)
  }
}
