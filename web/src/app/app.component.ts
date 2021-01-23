import { Component, OnInit } from '@angular/core';
import { SessionService } from './session/session.service';
import { Session } from "./session/session.struct";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Clærance';
  session: Session;

  constructor(private sess: SessionService) {
  }

  ngOnInit(): void {
    this.sess.isAuthenticated().subscribe(
      session => this.session = session,
      () => this.session = SessionService.null_session
    );
  }

  logout() {
    this.sess.logout().subscribe();
  }
}
