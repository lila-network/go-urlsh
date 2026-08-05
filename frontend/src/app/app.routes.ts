import {Routes} from '@angular/router';
import {LinkOverview} from './features/links/link-overview/link-overview';
import {LinkDetails} from './features/links/link-details/link-details';
import { LinkCreate } from './features/links/link-create/link-create';
import { LinkDelete } from './features/links/link-delete/link-delete';

export const routes: Routes = [
  {
    path: 'links',
    component: LinkOverview,
  },
  {
    path: 'links/new',
    component: LinkCreate,
  },
  {
    path: 'links/:id/delete',
    component: LinkDelete,
  },
  {
    path: 'links/:id',
    component: LinkDetails,
  },
];
