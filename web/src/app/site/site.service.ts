import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Site } from '../sites/site.struct';

@Injectable({
  providedIn: 'root'
})
export class SiteService {
  apiPath = 'api/site';

  constructor(private http: HttpClient) { }

  public getAllSites(): Observable<Site[]> {
    return this.http.get<Site[]>(`${this.apiPath}/list`);
  }
}
