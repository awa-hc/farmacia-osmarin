import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DashboardProductComponent } from './dashboard-product.component';

describe('DashboardProductComponent', () => {
  let component: DashboardProductComponent;
  let fixture: ComponentFixture<DashboardProductComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DashboardProductComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DashboardProductComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
