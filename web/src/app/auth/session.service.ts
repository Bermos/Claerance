import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  constructor(private http: HttpClient) { }

  public isAuthenticated(): Observable<any> {
    return this.http.get('api/session/', { observe: 'response' });
  }

  public logout(): Observable<any> {
    return this.http.delete('api/session/');
  }
}
