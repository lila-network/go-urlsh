import { Component } from '@angular/core';
import { ThemeToggleButton } from './components/theme-toggle-button/theme-toggle-button';
import { RouterOutlet, RouterLink } from '@angular/router';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet, RouterLink, ThemeToggleButton],
  templateUrl: './app.html',
  styleUrl: './app.scss',
})
export class App {
}

