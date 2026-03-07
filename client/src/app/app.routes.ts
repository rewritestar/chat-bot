import { Routes } from '@angular/router';
import { Signin } from './features/auth/signin/signin.component';
import { Login } from './features/auth/login/login.component';
import { Dashboard } from './features/dashboard/dashboard.component';
import { AuthLayout } from './features/auth/layout/auth-layout.component';
import { AuthGuard } from './core/services/auth-guard';

export const routes: Routes = [
  {
    path: '',
    component: Dashboard,
    canActivate: [AuthGuard],
  },
  {
    path: '',
    component: AuthLayout,
    children: [
      {
        path: 'signin',
        component: Signin,
      },
      {
        path: 'login',
        component: Login,
      },
    ],
  },
];
