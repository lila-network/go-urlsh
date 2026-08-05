import { Component, inject, signal } from '@angular/core';
import { RouterLink } from '@angular/router';
import { Link, LinksResponse } from '../link';
import { LinkService } from '../link.service';
import { AlertMessage } from '../../../components/alert-message/alert-message';
import {
  LucideSquarePen,
  LucideTrash2,
  LucideInfo,
  LucideRefreshCw,
  LucidePlus,
} from '@lucide/angular';

@Component({
  selector: 'link-overview',
  imports: [
    RouterLink,
    LucideSquarePen,
    LucideTrash2,
    LucideInfo,
    AlertMessage,
    LucideRefreshCw,
    LucidePlus,
  ],
  templateUrl: './link-overview.html',
  styleUrl: './link-overview.scss',
})
export class LinkOverview {
  private readonly linkService = inject(LinkService);

  protected readonly errorMessage = signal<string>('');
  protected readonly successMessage = signal<string>('');
  protected readonly loading = signal<boolean>(false);
  protected readonly links = signal<Link[]>([]);

  protected setError() {
    this.errorMessage.set(
      'Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet. Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea takimata sanctus est Lorem ipsum dolor sit amet.',
    );
  }

  protected setLoading() {
    this.loading.set(true);
  }

  protected setEmpty() {
    this.setClear();
    this.links.set([]);
  }

  protected setContent() {
    this.setClear();

    this.links.set([
      {
        id: 'TestLink1',
        url: 'http://localhost:8080',
        description: '',
        created: '1920-13-24',
        modified: '2010-13-24',
      },
      {
        id: 'TestLink2',
        url: 'http://bürpelhost:8080',
        description: 'asdsad',
        created: '1920-13-24',
        modified: '2010-13-24',
      },
    ]);
  }

  protected setClear() {
    this.loading.set(false);
    this.errorMessage.set('');
  }

  protected fetchLinks() {
    this.setLoading();
    let linksObservable = this.linkService.getLinks();
    linksObservable.subscribe({
      next: (response: LinksResponse) => {
        this.links.set(response.data);
      },
      error: (err) => {
        // RUNS ONLY ON FAILURE (e.g. HTTP 404, 500, or CORS failure)
        // The 'next' and 'complete' blocks will NEVER run if 'error' triggers
        this.loading.set(false);
        this.errorMessage.set(`Failed to load links from local API: ${err.message}`);
      },
      complete: () => {
        this.loading.set(false);
      },
    });
  }

  protected fetchLinksOnce() {
    let links = this.links()

    if (links.length != 0) {
      return
    }

    this.fetchLinks()
  }

  ngOnInit() {
    this.fetchLinksOnce();
  }

  constructor() {
    const message = history.state.successMessage;

    if (typeof message === 'string') {
      this.successMessage.set(message);
    }
  }
}
