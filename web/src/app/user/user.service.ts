import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { User } from './user.struct';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  apiPath = 'api/user';

  constructor(private http: HttpClient) { }

  /**
   * Retrieves the user details for the requested user.
   * @param id of the requested user
   */
  public getUser(id: number): Observable<User> {
    return this.http.get<User>(`${this.apiPath}/${id}`);
  }

  /**
   * Retrieves details of all users.
   */
  public getAllUsers(): Observable<User[]> {
    return this.http.get<User[]>(`${this.apiPath}/list`);
  }

  /**
   * Updates the user with the new details.
   * @param user to be updated with updated attributes
   */
  public updateUser(user: User): Observable<any> {
    return this.http.put(`${this.apiPath}/${user.ID}`, user);
  }

  /**
   * Sends a creation request to the server
   * @param username for the new user
   * @param password for the new user
   */
  public createUser(username: string, password: string): Observable<any> {
    return this.http.post(`${this.apiPath}/create`, {
      username,
      password
    });
  }

  /**
   * Sends a deletion request to the server
   * @param id of the user to be deleted
   */
  public deleteUser(id: number): Observable<any> {
    return this.http.delete(`${this.apiPath}/${id}`);
  }
}
