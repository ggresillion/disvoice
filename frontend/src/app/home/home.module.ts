import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { HomeComponent } from './home.component';
import { RouterModule, Routes } from '@angular/router';
import { EffectsComponent } from './components/effects/effects.component';
import { CoreModule } from '../core/core.module';

const routes: Routes = [
  {
    path: '',
    component: HomeComponent
  }
];

@NgModule({
  declarations: [HomeComponent, EffectsComponent],
  imports: [
    CommonModule,
    RouterModule.forChild(routes),
    CoreModule
  ],
  exports: [
    RouterModule
  ]
})
export class HomeModule { }
