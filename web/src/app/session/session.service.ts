import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, Subject } from 'rxjs';
import { Session } from "./session.struct";

@Injectable({
  providedIn: 'root'
})
export class SessionService {
  private static session: Subject<Session> = new Subject<Session>();
  public static null_session: Session = new Session(null, false, null);

  constructor(private http: HttpClient) { }

  public isAuthenticated(): Observable<Session> {
    this.http.get<Session>('api/session/', ).subscribe(
      sess => {SessionService.session.next(sess)},
      () => SessionService.session.next(SessionService.null_session)
    );
    return SessionService.session
  }

  public logout(): Observable<Session> {
    this.http.delete<Session>('api/session/').subscribe(
      () => SessionService.session.next(SessionService.null_session),
      () => SessionService.session.error("Something went wrong")
    );
    return SessionService.session;
  }

  public update() {
    this.isAuthenticated()
  }
}
