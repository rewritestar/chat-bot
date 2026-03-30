import { Router } from '@angular/router';
import {
  HttpErrorResponse,
  HttpEvent,
  HttpHandler,
  HttpInterceptor,
  HttpRequest,
  HttpStatusCode,
} from '@angular/common/http';
import { Injectable } from '@angular/core';

import { catchError, Observable, throwError } from 'rxjs';
import { ToastrService } from 'ngx-toastr';

import { WsService } from './ws.service';
import { AuthService } from './auth.service';

@Injectable({
  providedIn: 'root',
})
export class AuthInterceptor implements HttpInterceptor {
  constructor(
    private authService: AuthService,
    private router: Router,
    private toastr: ToastrService,
    private wsService: WsService,
  ) {}
  intercept(req: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    const authReq = this.setAuthHeader(req);

    return next.handle(authReq).pipe(
      catchError((error: HttpErrorResponse) => {
        this.handleError(error);
        return throwError(() => error);
      }),
    );
  }

  setAuthHeader(req: HttpRequest<any>): HttpRequest<any> {
    const token = this.authService.getToken();
    if (token) {
      return req.clone({
        headers: req.headers.set('Authorization', 'Bearer ' + token),
      });
    }
    return req;
  }

  handleError(error: HttpErrorResponse) {
    this.showMessage(error);

    switch (error.status) {
      case HttpStatusCode.Unauthorized:
        this.handleUnauthorized();
    }
  }

  handleUnauthorized() {
    this.authService.logout();
    this.wsService.close();
    this.router.navigate(['login']);
  }

  showMessage(error: HttpErrorResponse) {
    this.toastr.error(error.error?.message, 'Error');
  }
}
