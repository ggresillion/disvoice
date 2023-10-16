import { Injectable } from '@angular/core';
import { from, Observable, ObservableInput } from 'rxjs';
import { Config } from '../models/config.model';

@Injectable({
  providedIn: 'root'
})
export class BackendService {

  constructor() { }

  public getBackend(): any {
    // @ts-ignore
    return window.backend;
  }

  public getConfig(): Observable<Config> {
    return from<ObservableInput<Config>>(this.getBackend().getConfig());
  }

  public toggleEffect(id: string): Observable<void> {
    return from<ObservableInput<void>>(this.getBackend().toggleEffect(id));
  }

  public showSettings(id: string): Observable<void> {
    return from<ObservableInput<void>>(this.getBackend().showSettings(id));
  }
}
