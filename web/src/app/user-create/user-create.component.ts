import { Component, OnInit } from '@angular/core';
import { UserService } from '../user/user.service';
import { Router } from '@angular/router';

@Component({
  selector: 'app-user-create',
  templateUrl: './user-create.component.html',
  styleUrls: ['./user-create.component.css']
})
export class UserCreateComponent implements OnInit {
  username: string;
  password: string;
  passwordConfirm: string;

  constructor(private router: Router, private us: UserService) { }

  ngOnInit(): void {
  }

  createUser() {
    if (this.password !== this.passwordConfirm) {
      alert('Passwords don\'t match.');
      return;
    }

    this.us.createUser(this.username, this.password).subscribe(
      () => this.router.navigate(['/users']),
      error => console.log('Could not create user.', error)
    );
  }
}
