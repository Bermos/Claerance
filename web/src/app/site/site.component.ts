import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { SiteService } from './site.service';
import { Site } from './site.struct';

@Component({
  selector: 'app-site',
  templateUrl: './site.component.html',
  styleUrls: ['./site.component.css']
})
export class SiteComponent implements OnInit {
  site: Site;

  constructor(private route: ActivatedRoute, private router: Router, private ss: SiteService) { }

  ngOnInit(): void {
    const siteId = +this.route.snapshot.paramMap.get('id');
    this.ss.getSite(siteId).subscribe({
      next: (site) => this.site = site,
      error: () => this.router.navigate(['/dashboard'])
    });
  }

  editSite() {

  }

  deleteSite() {
    this.ss.deleteSite(this.site.ID).subscribe(
      () => this.router.navigate(['/dashboard']),
    );
  }
}
