import { Component, inject, signal } from '@angular/core';
import { RouterLink, ActivatedRoute } from '@angular/router';

@Component({
  selector: 'link-details',
  imports: [RouterLink],
  templateUrl: './link-details.html',
  styleUrl: './link-details.scss',
})
export class LinkDetails {
  linkId = signal('');
  private activatedRoute = inject(ActivatedRoute);

  constructor() {
    // Access route parameters
    this.activatedRoute.params.subscribe((params) => {
      this.linkId.set(params['id']);
    });
  }
}
