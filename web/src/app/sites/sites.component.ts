import { AfterViewInit, Component, OnInit, ViewChild } from '@angular/core';
import { MatTableDataSource } from '@angular/material/table';
import { MatPaginator } from '@angular/material/paginator';
import { SiteService } from '../site/site.service';
import { Site } from '../site/site.struct';

@Component({
  selector: 'app-sites',
  templateUrl: './sites.component.html',
  styleUrls: ['./sites.component.css']
})
export class SitesComponent implements OnInit, AfterViewInit {
  displayedColumns: string[] = ['name', 'url', 'createdAt'];
  sitesDataSource = new MatTableDataSource<Site>();
  @ViewChild(MatPaginator) paginator: MatPaginator;

  ngAfterViewInit() {
    this.sitesDataSource.paginator = this.paginator;
  }

  constructor(private ss: SiteService) { }

  ngOnInit(): void {
    this.ss.getAllSites().subscribe({
      next: (sites) => this.sitesDataSource.data = sites
    });
  }

  applyFilter(event: KeyboardEvent) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.sitesDataSource.filter = filterValue.trim().toLowerCase();
  }
}
