import { ComponentFixture, TestBed } from '@angular/core/testing';

import { LinkCreate } from './link-create';

describe('LinkCreate', () => {
  let component: LinkCreate;
  let fixture: ComponentFixture<LinkCreate>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [LinkCreate],
    }).compileComponents();

    fixture = TestBed.createComponent(LinkCreate);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
