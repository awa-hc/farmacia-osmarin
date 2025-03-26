import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DashboardSuppliersComponent } from './dashboard-suppliers.component';

describe('DashboardSuppliersComponent', () => {
  let component: DashboardSuppliersComponent;
  let fixture: ComponentFixture<DashboardSuppliersComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DashboardSuppliersComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DashboardSuppliersComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
