import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  private sessionsUrl = 'api/v1/session';
  username: string;
  password: string;

  constructor(private http: HttpClient, private router: Router) { }

  ngOnInit(): void {
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
