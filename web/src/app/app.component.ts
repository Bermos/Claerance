import {Component, OnInit} from '@angular/core';
import {SessionService} from './auth/session.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Clærance';
  loggedIn = false;
  username: string;

  constructor(private sess: SessionService) { }

  ngOnInit(): void {
    this.sess.isAuthenticated().subscribe(
      res => {
        this.loggedIn = true;
        this.username = res.body['username'];
      }
    );
  }

  logout() {
    this.sess.logout().subscribe();
  }
}
