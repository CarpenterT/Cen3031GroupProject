import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ChatBoxTestComponent } from './chat-box-test.component';

describe('ChatBoxTestComponent', () => {
  let component: ChatBoxTestComponent;
  let fixture: ComponentFixture<ChatBoxTestComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ChatBoxTestComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ChatBoxTestComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
