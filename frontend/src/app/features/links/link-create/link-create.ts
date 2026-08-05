import { Component, signal, inject } from '@angular/core';
import { Router, RouterLink } from '@angular/router';
import { FormsModule, NgForm } from '@angular/forms';
import { LinkCreationRequest } from '../link';
import { LinkService } from '../link.service';
import { AlertMessage } from '../../../components/alert-message/alert-message';

@Component({
  selector: 'link-create',
  imports: [AlertMessage, RouterLink, FormsModule],
  templateUrl: './link-create.html',
  styleUrl: './link-create.scss',
})
export class LinkCreate {
  protected readonly isLoading = signal<boolean>(false);
  protected readonly errorMessage = signal<string>('test error');
  protected readonly linkService = inject(LinkService);
  private readonly router = inject(Router);

  formSlug = '';
  formTarget = '';
  formDescription = '';

  createLink(form: NgForm) {
    if (form.invalid) return;

    this.isLoading.set(true);

    let newLink: LinkCreationRequest = {
      id: this.formSlug,
      url: this.formTarget,
      description: this.formDescription,
    };

    let observable = this.linkService.newLink(newLink);

    observable.subscribe({
      next: () => {
        this.router.navigate(['/links'], {
          state: {
            successMessage: 'New Shortlink was successfully created!'
          }
        });

      },
      error: (err) => {
        // RUNS ONLY ON FAILURE (e.g. HTTP 404, 500, or CORS failure)
        // The 'next' and 'complete' blocks will NEVER run if 'error' triggers
        this.isLoading.set(false);
        this.errorMessage.set(`Failed to create link: ${err.message}`);
      },
      complete: () => {
        this.isLoading.set(false);
      },
    });
  }
}
