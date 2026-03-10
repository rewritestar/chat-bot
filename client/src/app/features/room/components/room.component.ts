import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { RoomApiService } from '../services/room-api.service';
import { AsyncPipe } from '@angular/common';
import { Observable, startWith, Subject, switchMap, tap } from 'rxjs';
import { AuthService } from '../../../core/services/auth.service';

@Component({
  selector: 'room',
  templateUrl: 'room.html',
  imports: [ReactiveFormsModule, AsyncPipe],
})
export class Room {
  private reload$ = new Subject<void>();
  room$: Observable<any>;
  roomId: number = 0;
  workerId: number;

  form = new FormGroup({
    content: new FormControl(''),
  });

  defaultForm = {
    content: '',
  };

  constructor(
    private roomApi: RoomApiService,
    private authService: AuthService,
  ) {
    this.room$ = this.reload$.pipe(
      startWith(void 0),
      switchMap(() => this.roomApi.getRoom()),
      tap((room: any) => (this.roomId = room.id)),
    );
    this.workerId = this.authService.getId();
  }

  roadData() {
    this.reload$.next();
  }

  onSubmit() {
    if (this.form.invalid) {
      console.log(this.form.errors);
      return;
    }

    const req = {
      roomId: this.roomId,
      content: this.form.value.content?.trim(),
    };

    this.roomApi.saveChat(req).subscribe(() => {
      this.roadData();
      this.resetForm();
    });
  }

  resetForm() {
    this.form.reset(this.defaultForm);
  }

  isMyChat(chat: any): boolean {
    return chat.creatorId === this.workerId;
  }
}
