import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { AsyncPipe, NgClass } from '@angular/common';
import { Router } from '@angular/router';

import { BehaviorSubject, startWith, Subject, switchMap, tap } from 'rxjs';
import { ToastrService } from 'ngx-toastr';

import { RoomApiService } from '../../services/room-api.service';
import { WsService } from '../../../../core/services/ws.service';

@Component({
  selector: 'room-list',
  templateUrl: 'room-list.html',
  imports: [ReactiveFormsModule, AsyncPipe, NgClass],
})
export class RoomList {
  private reload$ = new Subject<void>();
  roomList$ = new BehaviorSubject<any[]>([]);
  isOpenNewRoom: boolean = false;
  isOpenDeleteRoom: boolean = false;
  selectedRoomId: number | null = null;

  roomForm = new FormGroup({
    name: new FormControl('', [Validators.required]),
    aiSystem: new FormControl(''),
  });

  isNameValid: boolean = true;

  constructor(
    private roomApi: RoomApiService,
    private router: Router,
    private wsService: WsService,
    private toastr: ToastrService,
  ) {
    this.reload$
      .pipe(
        startWith(void 0),
        switchMap(() => this.roomApi.indexRoom()),
        tap((roomList: any) => {
          this.roomList$.next(roomList.list ?? []);
        }),
      )
      .subscribe();

    this.wsService.leaveRoom();

    this.wsService.getMessage().subscribe((data) => {
      const currentList = this.roomList$.value || [];
      const selectedRoom = currentList.find((room) => room.id === data.roomId);
      if (!selectedRoom) {
        return;
      }

      const updatedRoom = {
        ...selectedRoom,
        unreadCount: selectedRoom.unreadCount + 1,
        chatList: [data.content],
      };

      const updatedList = [updatedRoom, ...currentList.filter((room) => room.id !== data.roomId)];
      this.roomList$.next(updatedList);
    });

    this.roomForm.get('name')?.statusChanges.subscribe((value) => {
      this.isNameValid = value === 'VALID' || this.isNameValid;
    });
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
      this.isNameValid = this.roomForm.get('name')?.valid || false;
      return;
    }

    const req = {
      name: this.roomForm.value.name?.trim(),
      aiSystem: this.roomForm.value.aiSystem?.trim(),
    };

    this.roomApi.saveRoom(req).subscribe((res: any) => {
      this.toastr.success('성공적으로 생성되었습니다.');
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
      this.toastr.success('성공적으로 삭제되었습니다.');
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
