import { ComponentFixture, TestBed } from '@angular/core/testing';

import { WhyhiremeComponent } from './whyhireme.component';

describe('WhyhiremeComponent', () => {
  let component: WhyhiremeComponent;
  let fixture: ComponentFixture<WhyhiremeComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [WhyhiremeComponent]
    })
    .compileComponents();

    fixture = TestBed.createComponent(WhyhiremeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
