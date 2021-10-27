import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SiteService } from '../site.service';

@Component({
  selector: 'app-site-create',
  templateUrl: './site-create.component.html',
  styleUrls: ['./site-create.component.css']
})
export class SiteCreateComponent implements OnInit {
  name: string;
  url: string;

  constructor(private router: Router, private ss: SiteService) { }

  ngOnInit(): void {
  }

  createSite() {
    this.ss.createSite(this.name, this.url).subscribe(
      () => this.router.navigate(['/sites']),
    );
  }
}
