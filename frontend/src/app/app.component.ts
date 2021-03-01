import { Component, OnInit } from '@angular/core';
import { AudioService } from './core/services/audio.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit {

  constructor(private readonly audioService: AudioService) { }

  public ngOnInit(): void {
    console.log('Starting audio...');
    this.audioService.startAudio().subscribe(() => {
      console.log('Started audio');
    });
  }
}
