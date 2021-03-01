import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { CoreModule } from '../core.module';
import { Effect } from '../models/effect.model';
import { BackendService } from './backend.service';
import { ConfigService } from './config.service';

@Injectable({
  providedIn: CoreModule
})
export class EffectService {

  constructor(
    private readonly configService: ConfigService,
    private readonly backendService: BackendService) { }

  public getEffects(): Observable<Effect[]> {
    return this.configService.getConfig()
    .pipe(map(c => c.effects));
  }

  public toggleEffect(id: string): Observable<void> {
    return this.backendService.toggleEffect(id);
  }

  public showSettings(id: string): Observable<void> {
    return this.backendService.showSettings(id);
  }
}
