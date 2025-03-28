import { Component, OnInit } from '@angular/core';
import {
  AbstractControl,
  FormArray,
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { DialogModule } from 'primeng/dialog';
import { CommonModule, CurrencyPipe } from '@angular/common';
import { TableModule } from 'primeng/table';
import { DashboardService } from '../../pages/dashboard/dashboard.service';

@Component({
  selector: 'app-dashboard-sale',
  imports: [
    TableModule,
    ReactiveFormsModule,
    DialogModule,
    CurrencyPipe,
    CommonModule,
  ],
  templateUrl: './dashboard-sales.component.html',
  styleUrls: ['./dashboard-sales.component.css'],
})
export class DashboardSalesComponent implements OnInit {
  sales: any[] = [];
  products: any[] = [];
  displayDialog: boolean = false;
  isEditMode: boolean = false;
  saleForm: FormGroup;

  constructor(
    private dashboardService: DashboardService,
    private fb: FormBuilder
  ) {
    // Definir todos los campos del formulario
    this.saleForm = this.fb.group({
      ID: [null],
      customerName: ['', Validators.required], // Nombre del cliente
      totalAmount: [0, Validators.required], // Monto total de la venta
      status: ['', Validators.required], // Estado de la venta
      details: this.fb.array([]), // Detalles de la venta
    });
  }

  ngOnInit(): void {
    this.loadSales();
    this.loadProducts();
  }

  loadSales(): void {
    this.dashboardService.getAll<any>('sales').subscribe((response: any) => {
      this.sales = response.data; // Ajusta según la estructura de tu respuesta del backend
    });
  }

  loadProducts(): void {
    this.dashboardService.getAll<any>('products').subscribe((response: any) => {
      console.log(response.data);
      this.products = response.data; // Cargar lista de productos disponibles
    });
  }

  openCreateDialog(): void {
    this.isEditMode = false;
    this.saleForm.reset(); // Limpiar el formulario
    this.displayDialog = true;
  }

  openEditDialog(sale: any): void {
    this.isEditMode = true;
    this.saleForm.patchValue({
      ID: sale.ID,
      customerName: sale.customer_name,
      totalAmount: Number(sale.total_amount), // Convertir a número
      status: sale.status,
    });

    const details = sale.details.map((detail: any) =>
      this.fb.group({
        productID: detail.product_id,
        quantity: Number(detail.quantity), // Convertir a número
        unitPrice: Number(detail.unit_price), // Convertir a número
        subtotal: Number(detail.subtotal), // Convertir a número
      })
    );
    this.details.clear();
    details.forEach((detail: any) => this.details.push(detail));
    this.displayDialog = true;
  }

  closeDialog(): void {
    this.displayDialog = false;
  }

  onSubmit(): void {
    if (this.saleForm.valid) {
      const saleData = this.transformFormData(this.saleForm.value);
      if (this.isEditMode) {
        this.dashboardService
          .update('sales', saleData.ID, saleData)
          .subscribe(() => {
            this.loadSales();
            this.closeDialog();
          });
      } else {
        console.log(saleData);
        this.dashboardService.create('sales', saleData).subscribe(() => {
          this.loadSales();
          this.closeDialog();
        });
      }
    }
  }

  deleteSale(ID: number): void {
    this.dashboardService.delete('sales', ID).subscribe(() => {
      this.loadSales();
    });
  }

  // Transformar los datos del formulario para que coincidan con el backend
  transformFormData(formData: any): any {
    const now = new Date().toISOString();
    return {
      ID: formData.ID,
      customer_name: formData.customerName,
      sale_date: now,
      total_amount: Number(formData.totalAmount),
      status: formData.status,
      details: formData.details.map((detail: any) => ({
        product_id: Number(detail.productID),
        quantity: Number(detail.quantity),
        unit_price: Number(detail.unitPrice),
        subtotal: Number(detail.subtotal),
      })),
    };
  }

  // Métodos para manejar detalles de la venta
  get details(): FormArray {
    return this.saleForm.get('details') as FormArray;
  }

  addDetail(): void {
    const detailGroup = this.fb.group({
      productID: ['', Validators.required],
      quantity: [1, [Validators.required, Validators.min(1)]],
      unitPrice: [0, Validators.required],
      subtotal: [0, Validators.required],
    });

    // Observar cambios en cantidad
    detailGroup.get('quantity')?.valueChanges.subscribe(() => {
      this.calculateSubtotal(detailGroup);
    });

    // Observar cambios en precio unitario
    detailGroup.get('unitPrice')?.valueChanges.subscribe(() => {
      this.calculateSubtotal(detailGroup);
    });

    this.details.push(detailGroup);
  }

  calculateSubtotal(detailGroup: FormGroup): void {
    const quantity = Number(detailGroup.get('quantity')?.value) || 0;
    const unitPrice = Number(detailGroup.get('unitPrice')?.value) || 0;
    const subtotal = quantity * unitPrice;

    // Actualizar el campo subtotal usando patchValue
    detailGroup.patchValue({ subtotal }, { emitEvent: false });

    // Recalcular el total de la venta
    this.calculateTotalAmount();
  }

  calculateTotalAmount(): void {
    const total = this.details.controls.reduce(
      (sum: number, control: AbstractControl) => {
        if (control instanceof FormGroup) {
          return sum + (Number(control.get('subtotal')?.value) || 0);
        }
        return sum;
      },
      0
    );
    this.saleForm.get('totalAmount')?.setValue(total);
  }

  removeDetail(index: number): void {
    this.details.removeAt(index);
    this.calculateTotalAmount();
  }

  // En el método onProductChange
  onProductChange(detailControl: AbstractControl): void {
    const detailGroup = detailControl as FormGroup;
    const productID = Number(detailGroup.get('productID')?.value);
    const product = this.products.find((p) => p.ID === productID);

    if (product && product.selling_price !== undefined) {
      detailGroup
        .get('unitPrice')
        ?.setValue(Number(product.selling_price), { emitEvent: true });
      this.calculateSubtotal(detailGroup);
    }
  }

  // En el método getProductByID (si lo sigues usando)
  getProductByID(productID: number): any {
    return this.products.find((product) => product.ID === Number(productID));
  }
}
