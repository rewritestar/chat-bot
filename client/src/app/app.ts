import { Component, signal } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { WsService } from './core/services/ws.service';

@Component({
  selector: 'app-root',
  imports: [RouterOutlet],
  templateUrl: './app.html',
})
export class App {
  protected readonly title = signal('client');
  constructor(private wsService: WsService) {
    this.wsService.connect();
  }
}
