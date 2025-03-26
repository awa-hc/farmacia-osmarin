import { Component } from '@angular/core';
import { routes } from '../../app.routes';
import { RouterOutlet } from '@angular/router';
import { DrawerModule } from 'primeng/drawer';
import { ButtonModule } from 'primeng/button';
import { AvatarModule } from 'primeng/avatar';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [
    RouterOutlet,
    DrawerModule,
    ButtonModule,
    AvatarModule,
    CommonModule,
    RouterModule,
  ],
  templateUrl: './dashboard.component.html',
  styleUrl: './dashboard.component.css',
})
export class DashboardComponent {
  visible: boolean = false;

  // Filtrar las rutas hijas del DashboardComponent
  menuRoutes =
    routes
      .find((route) => route.path === 'dashboard')
      ?.children?.filter(
        (childRoute) =>
          !childRoute.path?.includes(':id') && childRoute.path !== ''
      )
      .map((childRoute) => {
        const path = childRoute.path?.split('/')[0]; // Obtener la primera parte de la ruta
        return {
          path: path,
          label: this.getLabelFromPath(path),
          icon: this.getIconFromPath(path),
        };
      }) || [];

  getLabelFromPath(path: string | undefined): string {
    switch (path) {
      case 'products':
        return 'Productos';
      case 'suppliers':
        return 'Proveedores';
      case 'categories':
        return 'Categorías';
      case 'purchases':
        return 'Compras';
      case 'sales':
        return 'Ventas';
      default:
        return '';
    }
  }

  // Asignar íconos basados en la ruta
  getIconFromPath(path: string | undefined): string {
    switch (path) {
      case 'products':
        return 'pi pi-box';
      case 'suppliers':
        return 'pi pi-users';
      case 'categories':
        return 'pi pi-tags';
      case 'purchases':
        return 'pi pi-shopping-cart';
      case 'sales':
        return 'pi pi-dollar';
      default:
        return 'pi pi-home';
    }
  }
}
