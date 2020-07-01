import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  private auth: boolean;

  constructor(private http: HttpClient) { }

  public isAuthenticated(): boolean {
    this.http.head('api/v1/auth', { observe: 'response' }).subscribe(
      () => {
        this.auth = true;
        console.log(true);
      },
      () => {
        this.auth = false;
        console.log(false);
      }
    );
    return this.auth;
  }
}
