import {Component, inject } from '@angular/core';
import {LucideSun, LucideMoonStar} from '@lucide/angular';
import { ThemeService } from '../../theme.service';


@Component({
  selector: 'theme-toggle-button',
  imports: [
    LucideSun,
    LucideMoonStar
  ],
  templateUrl: './theme-toggle-button.html',
  styleUrl: './theme-toggle-button.scss',
})
export class ThemeToggleButton {
  protected readonly themeService = inject(ThemeService)

  toggleTheme() {
    this.themeService.toggleTheme()
  }
}
