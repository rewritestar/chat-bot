import { Routes } from '@angular/router';
import { AuthLayout } from './features/auth/layout/auth-layout.component';
import { AuthGuard } from './core/services/auth-guard';
import { Signin } from './features/auth/components/signin/signin.component';
import { Login } from './features/auth/components/login/login.component';
import { RoomList } from './features/room/components/room-list/room-list.component';
import { AppLayout } from './core/layout/app-layout.component';
import { Room } from './features/room/components/room/room.component';

export const routes: Routes = [
  {
    path: '',
    component: AppLayout,
    children: [
      {
        path: 'room',
        component: RoomList,
        canActivate: [AuthGuard],
      },
      {
        path: 'room/:roomId',
        component: Room,
        canActivate: [AuthGuard],
      },
    ],
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
