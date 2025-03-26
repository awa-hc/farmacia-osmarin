import { Routes } from '@angular/router';
import { HomeComponent } from './pages/home/home.component';
import { DashboardComponent } from './pages/dashboard/dashboard.component';
import { DashboardSalesComponent } from './dashboard/dashboard-sales/dashboard-sales.component';
import { DashboardPurchasesComponent } from './dashboard/dashboard-purchases/dashboard-purchases.component';
import { DashboardCategoriesComponent } from './dashboard/dashboard-categories/dashboard-categories.component';
import { DashboardSuppliersComponent } from './dashboard/dashboard-suppliers/dashboard-suppliers.component';
import { DashboardProductComponent } from './dashboard/dashboard-product/dashboard-product.component';
export const routes: Routes = [
  {
    path: '',
    component: HomeComponent,
  },

  {
    path: 'dashboard',
    component: DashboardComponent,
    children: [
      { path: 'products', component: DashboardProductComponent }, // Solo una página para productos
      { path: 'suppliers', component: DashboardSuppliersComponent }, // Solo una página para proveedores
      { path: 'categories', component: DashboardCategoriesComponent }, // Solo una página para categorías
      { path: 'purchases', component: DashboardPurchasesComponent }, // Solo una página para compras
      { path: 'sales', component: DashboardSalesComponent }, // Solo una página para ventas

      {
        path: '',
        redirectTo: 'products',
        pathMatch: 'full',
      },
    ],
  },

  { path: '**', redirectTo: '/dashboard', pathMatch: 'full' },
];
