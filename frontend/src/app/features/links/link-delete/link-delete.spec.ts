import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LinkDelete } from './link-delete';

describe('LinkDelete', () => {
  let component: LinkDelete;
  let fixture: ComponentFixture<LinkDelete>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LinkDelete],
    }).compileComponents();

    fixture = TestBed.createComponent(LinkDelete);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
