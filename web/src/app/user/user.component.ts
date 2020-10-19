import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

@Component({
  selector: 'app-user',
  templateUrl: './user.component.html',
  styleUrls: ['./user.component.css']
})
export class UserComponent implements OnInit {
  userId: number;
  username: string;

  constructor(private route: ActivatedRoute) { }

  ngOnInit() {
    this.userId = +this.route.snapshot.paramMap.get('id');
  }

}
