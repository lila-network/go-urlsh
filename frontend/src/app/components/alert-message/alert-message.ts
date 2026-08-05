import { Component, input } from '@angular/core';
import { LucideCircleAlert, LucideCircleCheck, LucideInfo } from '@lucide/angular';

@Component({
  selector: 'alert-message',
  imports: [LucideCircleAlert, LucideCircleCheck, LucideInfo],
  templateUrl: './alert-message.html',
  styleUrl: './alert-message.scss',
})
export class AlertMessage {
  public readonly message = input.required<string>();
  public readonly type = input<'error' | 'warning' | 'success' | 'info'>();
}
