import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  constructor(private http: HttpClient) { }

 /* private static handleError(error: HttpErrorResponse) {
    console.log(error);
    return throwError('Error while fetching session');
  }*/

  public isAuthenticated(): Observable<any> {
    return this.http.get('api/session/', { observe: 'response' }); /*.pipe(
      catchError(SessionService.handleError)
    );*/
  }

  public logout(): Observable<any> {
    return this.http.delete('api/session/');
  }
}
