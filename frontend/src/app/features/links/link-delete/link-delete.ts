import { Component, inject, signal } from '@angular/core';
import { ActivatedRoute, RouterLink } from '@angular/router';
import { Location } from '@angular/common';
import { LinkService } from '../link.service';
import { LinkResponse } from '../link';
import { toSignal } from '@angular/core/rxjs-interop';
import { catchError, map, Observable, of, switchMap } from 'rxjs';
import { AlertMessage } from '../../../components/alert-message/alert-message';
import { LucideTrash2 } from '@lucide/angular';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-link-delete',
  imports: [AlertMessage, RouterLink, LucideTrash2],
  templateUrl: './link-delete.html',
  styleUrl: './link-delete.scss',
})
export class LinkDelete {
  private activatedRoute = inject(ActivatedRoute);
  private linkService = inject(LinkService);

  protected readonly errorMessage = signal<string>('');

  // 1. Parameter aus der URL als Signal
  protected readonly linkId = toSignal(
    this.activatedRoute.paramMap.pipe(map((params) => params.get('id') ?? '')),
  );

  // 2. Sobald sich die ID ändert, wird automatisch getLink() aufgerufen
  protected readonly link = toSignal(
    this.activatedRoute.paramMap.pipe(
      map((params) => params.get('id') ?? ''),
      switchMap((id) => this.loadLink(id)),
    ),
  );

  private loadLink(id: string): Observable<LinkResponse | null> {
    return this.linkService
      .getLink(id)
      .pipe(catchError((error: HttpErrorResponse) => this.handleLoadError(error, id)));
  }

  private handleLoadError(error: HttpErrorResponse, id: string): Observable<null> {
    if (error.status === 404) {
      this.errorMessage.set(`The shortlink "${id}" could not be found.`);
    } else {
      this.errorMessage.set('The shortlink could not be loaded.');
    }

    return of(null);
  }

  protected cancelClicked() {
    this._location.back()
  }

  constructor(private _location: Location) {}
}
