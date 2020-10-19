import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from "./user.service";
import { User } from "./user.struct";

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  userId: number;
  user: User;

  constructor(private route: ActivatedRoute, private router: Router, private us: UserService) { }

  ngOnInit() {
    this.userId = +this.route.snapshot.paramMap.get('id');
    this.us.getUser(this.userId).subscribe(
      user => this.user = user,
      () => this.router.navigate(['/dashboard'])
    )
  }

  editUser() {
    console.log("Implement me edit mode")
  }

  deleteUser() {
    this.us.deleteUser(this.user.id).subscribe(
      () => this.router.navigate(['/dashboard']),
      err => console.log("Could not delete user.", err)
    )
  }
}
