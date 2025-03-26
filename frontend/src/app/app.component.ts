import { Component } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { TableModule } from 'primeng/table';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { CardModule } from 'primeng/card';
import { DropdownModule } from 'primeng/dropdown';

@Component({
  selector: 'app-root',
  standalone: true,
  providers: [HttpClient],
  imports: [
    CommonModule,
    FormsModule,
    TableModule,
    InputTextModule,
    ButtonModule,
    CardModule,
    DropdownModule,
  ],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent {
  productos: any[] = [];
  nuevoProducto = {
    producto: '',
    presentacion: '',
    stock: 0,
    precio: 0,
    vencimiento: '',
  };
  venta = { producto_id: 0, cantidad: 0 };

  constructor(private http: HttpClient) {}

  ngOnInit() {
    this.cargarProductos();
  }

  cargarProductos() {
    this.http
      .get<any[]>('https://farmacia-osmarin.fly.dev/productos')
      .subscribe((data) => {
        this.productos = data;
      });
  }

  agregarProducto() {
    this.http
      .post('https://farmacia-osmarin.fly.dev/productos', this.nuevoProducto)
      .subscribe(() => {
        this.cargarProductos();
      });
  }

  venderProducto() {
    this.http.post('https://farmacia-osmarin.fly.dev/vender', this.venta).subscribe(() => {
      this.cargarProductos();
    });
  }
}
