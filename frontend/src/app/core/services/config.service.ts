import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { Config } from '../models/config.model';
import { BackendService } from './backend.service';

@Injectable({
  providedIn: 'root'
})
export class ConfigService {

  constructor(private readonly backendService: BackendService) { }

  public getConfig(): Observable<Config> {
    return this.backendService.getConfig();
  }
}
