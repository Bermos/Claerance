import { Component, OnInit } from '@angular/core';
import { SessionService } from './auth/session.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'Clærance';
  loggedIn = false;
  username: string;
  userId: number;

  constructor(private sess: SessionService) {
  }

  ngOnInit(): void {
    this.sess.isAuthenticated().subscribe(
      res => {
        this.loggedIn = true;
        this.username = res.body['username'];
        this.userId = res.body['user_id'];
      },
      res => {
        console.log(res.status);
      }
    );
  }

  logout() {
    this.sess.logout().subscribe();
  }
}
