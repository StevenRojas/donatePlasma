// Angular
import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
// Components
import { DonorComponent } from './views/pages/donate/donor/donor.component';
import { HomeComponent } from './views/pages/donate/home/home.component';
import {ReceptorComponent} from './views/pages/donate/receptor/receptor.component';
import {ReceptorsComponent} from './views/pages/donate/receptors/receptors.component';
import {FaqComponent} from "./views/pages/donate/faq/faq.component";
import { DonorsComponent } from './views/pages/donate/donors/donors.component';

const routes: Routes = [
  {path: 'error', loadChildren: () => import('./views/pages/error/error.module').then(m => m.ErrorModule)},
  {path: '', component: HomeComponent, pathMatch: 'full'},
  {path: 'home', component: HomeComponent},
  {path: 'dona', component: DonorComponent},
  {path: 'donadores', component: DonorsComponent},
  {path: 'recibe', component: ReceptorComponent},
  {path: 'receptores', component: ReceptorsComponent},
  {path: 'faq', component: FaqComponent},
  {path: '404', redirectTo: 'error/404', pathMatch: 'full'},
  {path: '**', redirectTo: 'error/404', pathMatch: 'full'}
];

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forRoot(routes),
  ],
  exports: [RouterModule],
})
export class AppRoutingModule {
}
