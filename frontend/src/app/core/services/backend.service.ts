import { Injectable } from '@angular/core';
import { from, Observable, ObservableInput } from 'rxjs';
import { Config } from '../models/config.model';

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor() { }

  public getConfig(): Observable<Config> {
    // @ts-ignore
    return from<ObservableInput<void>>(getConfig());
  }

  public toggleEffect(id: string): Observable<void> {
    // @ts-ignore
    return from<ObservableInput<void>>(toggleEffect(id));
  }

  public showSettings(id: string): Observable<void> {
    // @ts-ignore
    return from<ObservableInput<void>>(showSettings(id));
  }

  public startAudio(): Observable<void> {
    // @ts-ignore
    return from<ObservableInput<void>>(startAudio());
  }
}
