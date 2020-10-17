import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import {SessionService} from '../auth/session.service';

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
    if (this.sess.isAuthenticated()) {
      this.router.navigate(['/dashboard']);
    }
  }

  login() {
    if (!this.username || !this.password) {
      console.log('Username or password missing');
      return;
    }

    this.http.post(this.sessionsUrl, {
      username: this.username,
      password: this.password
    }).subscribe(() => this.router.navigate(['/dashboard']));
  }
}
