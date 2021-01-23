import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { User } from '../user/user.struct';
import { MatPaginator } from '@angular/material/paginator';
import { UserService } from '../user/user.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit, AfterViewInit {
  displayedColumns: string[] = ['username', 'createdAt'];
  usersDataSource = new MatTableDataSource<User>();
  @ViewChild(MatPaginator) paginator: MatPaginator;

  ngAfterViewInit() {
    this.usersDataSource.paginator = this.paginator;
  }

  constructor(private us: UserService) { }

  ngOnInit(): void {
    this.us.getAllUsers().subscribe(
      users => this.usersDataSource.data = users
    );
  }

  applyFilter(event: KeyboardEvent) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.usersDataSource.filter = filterValue.trim().toLowerCase();
  }
}
