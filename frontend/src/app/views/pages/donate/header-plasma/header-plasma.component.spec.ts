import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { HeaderPlasmaComponent } from './header-plasma.component';

describe('HeaderPlasmaComponent', () => {
  let component: HeaderPlasmaComponent;
  let fixture: ComponentFixture<HeaderPlasmaComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ HeaderPlasmaComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(HeaderPlasmaComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
