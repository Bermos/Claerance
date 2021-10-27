import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Site } from './site.struct';

@Injectable({
  providedIn: 'root'
})
export class SiteService {
  apiPath = 'api/site';

  constructor(private http: HttpClient) { }

  public getAllSites(): Observable<Site[]> {
    return this.http.get<Site[]>(`${this.apiPath}/list`);
  }

  /**
   * Retrieves the site details for the requested site.
   * @param id of the requested site
   */
  public getSite(id: number): Observable<Site> {
    return this.http.get<Site>(`${this.apiPath}/${id}`);
  }

  /**
   * Updates the user with the new details.
   * @param site to be updated with updated attributes
   */
  public updateSite(site: Site): Observable<any> {
    return this.http.put(`${this.apiPath}/${site.ID}`, site);
  }

  /**
   * Sends a creation request to the server
   * @param name for the new user
   * @param url  for the new user
   */
  public createSite(name: string, url: string): Observable<any> {
    return this.http.post(`${this.apiPath}/create`, {
      name,
      url
    });
  }

  /**
   * Sends a deletion request to the server
   * @param id of the user to be deleted
   */
  public deleteSite(id: number): Observable<any> {
    return this.http.delete(`${this.apiPath}/${id}`);
  }
}
