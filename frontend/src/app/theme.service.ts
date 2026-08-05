import { effect, Service, signal } from '@angular/core';

@Service()
export class ThemeService {
  private readonly _theme = signal<'light' | 'dark'>('dark');
  public readonly theme = this._theme.asReadonly();

  public setTheme(theme: 'light' | 'dark') {
    this._theme.set(theme);
  }

  public toggleTheme() {
    if (this.theme() === 'light') {
      this.setTheme('dark')
    } else {
      this.setTheme('light')
    }
  }

  constructor() {
    effect(() => {
      const htmlElement = document.documentElement;

      htmlElement.setAttribute('data-bs-theme', this.theme());
    });
  }
}
