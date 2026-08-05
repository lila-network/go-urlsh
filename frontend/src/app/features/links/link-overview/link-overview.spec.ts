import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LinkOverview } from './link-overview';

describe('LinkOverview', () => {
  let component: LinkOverview;
  let fixture: ComponentFixture<LinkOverview>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LinkOverview],
    }).compileComponents();

    fixture = TestBed.createComponent(LinkOverview);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
