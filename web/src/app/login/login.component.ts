import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { SessionService } from '../session/session.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  private sessionsUrl = 'api/session/';
  username: string;
  password: string;

  constructor(private http: HttpClient, private router: Router, private sess: SessionService) { }

  ngOnInit(): void {
    // TODO: fix automatic redirection if already logged in
    this.sess.isAuthenticated().subscribe(
      session => {
        if (session.isValid) {
          this.router.navigate(['/dashboard']);
        }
      }
    );
  }

  login() {
    if (!this.username || !this.password) {
      console.log('Username or password missing');
      return;
    }

    this.http.post(this.sessionsUrl, {
      username: this.username,
      password: this.password
    }).subscribe(() => {
      this.sess.update();
      this.router.navigate(['/dashboard']);
    });
  }
}
