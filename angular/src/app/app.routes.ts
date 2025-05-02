import { Routes, RouterModule } from '@angular/router';
import { NgModule } from '@angular/core';
import { HomeComponent } from './features/public/home/home.component';
import { AboutComponent } from './features/public/about/about.component';
export const routes: Routes = [

    { path: '', component: HomeComponent },
    { path: 'about', component: AboutComponent },
    
];
