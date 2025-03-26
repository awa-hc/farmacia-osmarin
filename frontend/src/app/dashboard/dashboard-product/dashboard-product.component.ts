import { Component, OnInit } from '@angular/core';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { DashboardService } from '../../pages/dashboard/dashboard.service';
import { DialogModule } from 'primeng/dialog';
import { CommonModule, CurrencyPipe } from '@angular/common';
import { TableModule } from 'primeng/table';

@Component({
  selector: 'app-dashboard-product',
  imports: [
    TableModule,
    ReactiveFormsModule,
    DialogModule,
    CurrencyPipe,
    CommonModule,
  ],
  templateUrl: './dashboard-product.component.html',
  styleUrls: ['./dashboard-product.component.css'],
})
export class DashboardProductComponent implements OnInit {
  products: any[] = [];
  displayDialog: boolean = false;
  isEditMode: boolean = false;
  productForm: FormGroup;

  constructor(
    private dashboardService: DashboardService,
    private fb: FormBuilder
  ) {
    // Definir todos los campos del formulario
    this.productForm = this.fb.group({
      ID: [null],
      name: ['', Validators.required], // Nombre
      activeIngredient: ['', Validators.required], // Principio activo
      presentation: ['', Validators.required], // Presentación
      expiryDate: ['', Validators.required], // Fecha de vencimiento
      industry: ['', Validators.required], // Industria farmacéutica
      entryPrice: [0, Validators.required], // Precio de entrada
      sellingPrice: [0, Validators.required], // Precio de venta
      stock: [0, Validators.required], // Stock
      category: ['', Validators.required], // Categoría
    });
  }

  ngOnInit(): void {
    this.loadProducts();
  }

  loadProducts(): void {
    this.dashboardService.getAll<any>('products').subscribe((response: any) => {
      console.log(response);
      this.products = response.data; // Ajusta según la estructura de tu respuesta del backend
    });
  }

  openCreateDialog(): void {
    this.isEditMode = false;
    this.productForm.reset(); // Limpiar el formulario
    this.displayDialog = true;
  }

  openEditDialog(product: any): void {
    this.isEditMode = true;
    // Mapear los nombres de los campos del backend al formulario
    this.productForm.patchValue({
      ID: product.ID,
      name: product.name,
      activeIngredient: product.active_ingredient,
      presentation: product.presentation,
      expiryDate: product.expiry_date
        ? new Date(product.expiry_date).toISOString().split('T')[0]
        : '',
      industry: product.industry,
      entryPrice: product.entry_price,
      sellingPrice: product.selling_price,
      stock: product.stock,
      category: product.category,
    });
    this.displayDialog = true;
  }

  closeDialog(): void {
    this.displayDialog = false;
  }

  onSubmit(): void {
    if (this.productForm.valid) {
      const productData = this.transformFormData(this.productForm.value);
      if (this.isEditMode) {
        this.dashboardService
          .update('products', productData.ID, productData)
          .subscribe(() => {
            this.loadProducts();
            this.closeDialog();
          });
      } else {
        this.dashboardService.create('products', productData).subscribe(() => {
          this.loadProducts();
          this.closeDialog();
        });
      }
    }
  }

  deleteProduct(ID: number): void {
    this.dashboardService.delete('products', ID).subscribe(() => {
      this.loadProducts();
    });
  }

  // Transformar los datos del formulario para que coincidan con el backend
  transformFormData(formData: any): any {
    return {
      ID: formData.ID,
      name: formData.name,
      active_ingredient: formData.activeIngredient,
      presentation: formData.presentation,
      expiry_date: formData.expiryDate
        ? new Date(formData.expiryDate).toISOString()
        : null, // Formato RFC3339
      industry: formData.industry,
      entry_price: formData.entryPrice,
      selling_price: formData.sellingPrice,
      stock: formData.stock,
      category: formData.category,
    };
  }
}
