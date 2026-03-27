import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { RoomApiService } from '../../services/room-api.service';
import { AsyncPipe } from '@angular/common';
import { Observable, startWith, Subject, switchMap, tap } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'room-list',
  templateUrl: 'room-list.html',
  imports: [ReactiveFormsModule, AsyncPipe],
})
export class RoomList {
  private reload$ = new Subject<void>();
  roomList$: Observable<any>;
  isOpenNewRoom: boolean = false;
  isOpenDeleteRoom: boolean = false;
  selectedRoomId: number | null = null;

  roomForm = new FormGroup({
    name: new FormControl(''),
    aiSystem: new FormControl(''),
  });

  constructor(
    private roomApi: RoomApiService,
    private router: Router,
  ) {
    this.roomList$ = this.reload$.pipe(
      startWith(void 0),
      switchMap(() => this.roomApi.indexRoom()),
    );
  }

  roadData() {
    this.reload$.next();
  }

  openRoom(roomId: number) {
    this.router.navigate([`room/${roomId}`]);
  }

  onRoomSubmit() {
    if (this.roomForm.invalid) {
      console.log(this.roomForm.errors);
      return;
    }

    const req = {
      name: this.roomForm.value.name?.trim(),
      aiSystem: this.roomForm.value.aiSystem?.trim(),
    };

    this.roomApi.saveRoom(req).subscribe((res: any) => {
      this.closeNewRoom();
      this.router.navigate([`room/${res.id}`]);
    });
  }

  onDeleteRoom() {
    if (!this.selectedRoomId) {
      this.isOpenDeleteRoom = false;
      return;
    }
    this.roomApi.deleteRoom(this.selectedRoomId).subscribe(() => {
      this.closeDeleteRoom();
      this.roadData();
    });
  }

  openNewRoom() {
    this.isOpenNewRoom = true;
  }

  closeNewRoom() {
    this.isOpenNewRoom = false;
  }

  openDeleteRoom(roomId: number) {
    this.selectedRoomId = roomId;
    this.isOpenDeleteRoom = true;
  }

  closeDeleteRoom() {
    this.isOpenDeleteRoom = false;
  }
}
