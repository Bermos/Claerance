import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { Session } from './session.struct';

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  private static session: Subject<Session> = new Subject<Session>();
  public static nullSession: Session = new Session(null, false, null);
  private apiPath = 'api/session';

  constructor(private http: HttpClient) { }

  public isAuthenticated(): Observable<Session> {
    this.http.get<Session>(`${this.apiPath}/`).subscribe({
      next: (sess) => { SessionService.session.next(sess); },
      error: () => SessionService.session.next(SessionService.nullSession)
    });
    return SessionService.session;
  }

  public logout(): Observable<Session> {
    this.http.delete<Session>(`${this.apiPath}/`).subscribe({
      next: () => SessionService.session.next(SessionService.nullSession),
      error: () => SessionService.session.error('Something went wrong')
    });
    return SessionService.session;
  }

  public update() {
    this.isAuthenticated();
  }
}
