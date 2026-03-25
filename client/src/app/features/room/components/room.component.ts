import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { RoomApiService } from '../services/room-api.service';
import { AsyncPipe } from '@angular/common';
import { BehaviorSubject, Observable, startWith, Subject, switchMap, tap } from 'rxjs';
import { AuthService } from '../../../core/services/auth.service';
import { ChatService } from '../../../core/services/chat.service';

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
  chatList$ = new BehaviorSubject<any[]>([]);
  isOpenConfig: boolean = false;

  chatForm = new FormGroup({
    content: new FormControl(''),
  });

  roomForm = new FormGroup({
    name: new FormControl(''),
    aiSystem: new FormControl(''),
  });

  defaultChatForm = {
    content: '',
  };

  constructor(
    private roomApi: RoomApiService,
    private authService: AuthService,
    private chatService: ChatService,
  ) {
    this.room$ = this.reload$.pipe(
      startWith(void 0),
      switchMap(() => this.roomApi.getRoom()),
      tap((room: any) => {
        this.roomId = room.id;
        this.chatList$.next(room?.chatList);
        this.roomForm.patchValue(room);
      }),
    );
    this.workerId = this.authService.getId();

    this.chatService.connect();
    this.chatService.getMessage().subscribe((data) => {
      const current = this.chatList$.value || [];
      this.chatList$.next([...current, data]);
    });
  }

  roadData() {
    this.reload$.next();
  }

  onChatSubmit() {
    if (this.chatForm.invalid) {
      console.log(this.chatForm.errors);
      return;
    }

    const req = {
      roomId: this.roomId,
      content: this.chatForm.value.content?.trim(),
    };

    this.chatService.sendMessage(req);
    this.resetChatForm();
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

    this.roomApi.updateRoom(this.roomId, req).subscribe(() => {
      this.closeConfig();
      this.roadData();
    });
  }

  openConfig() {
    this.isOpenConfig = true;
  }

  closeConfig() {
    this.isOpenConfig = false;
  }

  resetChatForm() {
    this.chatForm.reset(this.defaultChatForm);
  }

  isMyChat(chat: any): boolean {
    return chat.creatorId === this.workerId;
  }
}
