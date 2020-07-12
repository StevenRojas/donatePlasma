// Angular
import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {HttpClientModule} from '@angular/common/http';
// Partials
import {PartialsModule} from '../partials/partials.module';
// Pages
import {CoreModule} from '../../core/core.module';
import {DonorComponent} from './donate/donor/donor.component';
import {HomeComponent} from './donate/home/home.component';
import {ReceptorComponent} from './donate/receptor/receptor.component';
import {RouterModule} from '@angular/router';
import {ReceptorsComponent} from './donate/receptors/receptors.component';
import {HeaderPlasmaComponent} from './donate/header-plasma/header-plasma.component';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatRadioModule} from '@angular/material/radio';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatSelectModule} from '@angular/material/select';
import {MatInputModule} from '@angular/material/input';
import {SweetAlert2Module} from '@sweetalert2/ngx-sweetalert2';
import { FaqComponent } from './donate/faq/faq.component';
import { FooterPlasmaComponent } from './donate/footer-plasma/footer-plasma.component';
import {ImageDirective} from "../../core/_base/layout";
import {ShareModule} from "ngx-sharebuttons";
import {MatButtonModule} from "@angular/material/button";
import {ShareButtonModule} from "ngx-sharebuttons/button";
import {ShareButtonsModule} from "ngx-sharebuttons/buttons";
import {ShareIconsModule} from "ngx-sharebuttons/icons";
import {MatPaginatorModule} from "@angular/material/paginator";
import {FilePondModule} from "ngx-filepond";
import { DonorsComponent } from './donate/donors/donors.component';
import {MatTabsModule} from "@angular/material/tabs";

@NgModule({
    declarations: [DonorComponent, HomeComponent, ReceptorComponent, ReceptorsComponent, HeaderPlasmaComponent, FaqComponent, FooterPlasmaComponent, ImageDirective, DonorsComponent],
  exports: [],
  imports: [
    CommonModule,
    HttpClientModule,
    FormsModule,
    CoreModule,
    PartialsModule,
    RouterModule,
    MatCheckboxModule,
    MatRadioModule,
    MatFormFieldModule,
    MatSelectModule,
    ReactiveFormsModule,
    MatInputModule,
    SweetAlert2Module,
    ShareModule,
    MatButtonModule,
    ShareButtonModule,
    ShareButtonsModule,
    ShareIconsModule,
    MatPaginatorModule,
    FilePondModule,
    MatTabsModule
  ],
  providers: []
})
export class PagesModule {
}
