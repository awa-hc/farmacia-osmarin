import { ComponentFixture, TestBed } from '@angular/core/testing';

import { DashboardPurchasesComponent } from './dashboard-purchases.component';

describe('DashboardPurchasesComponent', () => {
  let component: DashboardPurchasesComponent;
  let fixture: ComponentFixture<DashboardPurchasesComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [DashboardPurchasesComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(DashboardPurchasesComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
