import { Component } from '@angular/core';
import { FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import { RoomApiService } from '../services/room-api.service';

@Component({
  selector: 'room',
  templateUrl: 'room.html',
  imports: [ReactiveFormsModule],
})
export class Room {
  room: any;
  chatList: any[] = [];

  form = new FormGroup({
    content: new FormControl(''),
  });

  constructor(private roomApi: RoomApiService) {
    this.roomApi.getRoom().subscribe((res: any) => {
      this.room = res;
      this.chatList = res?.chatList;
    });
  }

  onSubmit() {
    if (this.form.invalid) {
      console.log(this.form.errors);
      return;
    }

    const req = {
      content: this.form.value.content?.trim(),
    };

    this.roomApi.saveChat(req).subscribe();
  }
}
