import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Service, inject, signal } from '@angular/core';
import { Link, LinksErrorResponse, LinksResponse, LinkResponse, LinkCreationRequest } from './link';

@Service()
export class LinkService {
  private readonly http = inject(HttpClient);
  private readonly _links = signal<Link[]>([]);
  public readonly links = this._links.asReadonly();

  getLinks(): Observable<LinksResponse> {
    return this.http.get<LinksResponse>('http://localhost:3000/api/v1/links');
  }

  getLink(id: string): Observable<LinkResponse> {
    return this.http.get<LinkResponse>('http://localhost:3000/api/v1/links/' + id);
  }

  newLink(link: LinkCreationRequest): Observable<LinkResponse | LinksErrorResponse> {
    return this.http.post<LinkResponse | LinksErrorResponse>(
      'http://localhost:3000/api/v1/links',
      link,
      {
        headers: {
          'Content-Type': 'application/json',
          Accept: 'application/json',
        },
      },
    );
  }

  deleteLink(id: string) {

  }
}
