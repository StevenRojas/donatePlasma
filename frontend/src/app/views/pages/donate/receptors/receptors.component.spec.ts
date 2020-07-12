import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { ReceptorsComponent } from './receptors.component';

describe('ReceptorsComponent', () => {
  let component: ReceptorsComponent;
  let fixture: ComponentFixture<ReceptorsComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ ReceptorsComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ReceptorsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
