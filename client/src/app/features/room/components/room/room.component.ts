import {
  AfterViewChecked,
  AfterViewInit,
  Component,
  ElementRef,
  Input,
  OnInit,
  ViewChild,
} from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { AsyncPipe } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';

import { BehaviorSubject, Observable, startWith, Subject, switchMap, tap } from 'rxjs';
import { ToastrService } from 'ngx-toastr';

import { RoomApiService } from '../../services/room-api.service';
import { AuthService } from '../../../../core/services/auth.service';
import { WsService } from '../../../../core/services/ws.service';

@Component({
  selector: 'room',
  templateUrl: 'room.html',
  imports: [ReactiveFormsModule, AsyncPipe],
})
export class Room implements AfterViewChecked {
  @ViewChild('scrollBottom') scrollBottom!: ElementRef;
  private reload$ = new Subject<void>();

  room$: Observable<any>;
  roomId!: number;
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
    private wsService: WsService,
    private route: ActivatedRoute,
    private router: Router,
    private toastr: ToastrService,
  ) {
    this.roomId = Number(this.route.snapshot.paramMap.get('roomId'));
    this.room$ = this.reload$.pipe(
      startWith(void 0),
      switchMap(() => this.roomApi.getRoom(this.roomId)),
      tap((room: any) => {
        this.chatList$.next(room?.chatList);
        this.roomForm.patchValue(room);
      }),
    );
    this.workerId = this.authService.getId();

    this.wsService.joinRoom(this.roomId);

    this.wsService.getMessage().subscribe((data) => {
      if (data.roomId === this.roomId) {
        const current = this.chatList$.value || [];
        this.chatList$.next([...current, data.content]);
        this.scrollToBottom(true);
      }
    });
  }

  ngAfterViewChecked() {
    this.scrollToBottom(false);
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

    this.wsService.sendMessage(req);
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

    this.roomApi.updateRoom(Number(this.roomId), req).subscribe(() => {
      this.toastr.success('성공적으로 업데이트되었습니다.');
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

  goRoomList() {
    this.router.navigate(['room']);
  }

  scrollToBottom(smooth: boolean) {
    const element = this.scrollBottom?.nativeElement;
    if (element) {
      this.scrollBottom?.nativeElement.scrollIntoView({
        behavior: smooth ? 'smooth' : 'auto',
      });
    }
  }
}
