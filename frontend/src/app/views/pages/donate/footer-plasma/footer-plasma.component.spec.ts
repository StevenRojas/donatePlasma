import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { FooterPlasmaComponent } from './footer-plasma.component';

describe('FooterPlasmaComponent', () => {
  let component: FooterPlasmaComponent;
  let fixture: ComponentFixture<FooterPlasmaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ FooterPlasmaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(FooterPlasmaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
