import { Component, OnInit } from '@angular/core';
import { Effect } from '../../../core/models/effect.model';
import { EffectService } from '../../../core/services/effect.service';

@Component({
  selector: 'app-effects',
  templateUrl: './effects.component.html',
  styleUrls: ['./effects.component.scss']
})
export class EffectsComponent implements OnInit {

  public effects: Effect[];

  constructor(private readonly effectService: EffectService) { }

  public ngOnInit(): void {
    this.effectService.getEffects().subscribe(e => {
      this.effects = e;
      console.log(e);
    });
  }

  public toggleEffect(effect: Effect) {
    this.effectService.toggleEffect(effect.id).subscribe(() => {
      console.log('Toggled effect: ' + effect.name);
    });
  }

  public showSettings(effect: Effect) {
    this.effectService.showSettings(effect.id).subscribe(() => {
      console.log('Opened settings: ' + effect.name);
    });
  }

}
