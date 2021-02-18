import { TestBed } from '@angular/core/testing';

import { EffectService } from './effect.service';

describe('EffectService', () => {
  beforeEach(() => TestBed.configureTestingModule({}));

  it('should be created', () => {
    const service: EffectService = TestBed.get(EffectService);
    expect(service).toBeTruthy();
  });
});
