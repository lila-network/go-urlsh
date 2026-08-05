import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LinkDetails } from './link-details';

describe('LinkDetails', () => {
  let component: LinkDetails;
  let fixture: ComponentFixture<LinkDetails>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LinkDetails],
    }).compileComponents();

    fixture = TestBed.createComponent(LinkDetails);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
