import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Directive, Input, HostListener } from '@angular/core';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { MatCardModule } from '@angular/material/card';
import { MatInputModule } from '@angular/material/input';
import { AccountCreateComponent } from './account-create.component';

describe('AccountCreateComponent', () => {
  let component: AccountCreateComponent;
  let fixture: ComponentFixture<AccountCreateComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [
        AccountCreateComponent,
        RouterLinkDirectiveStub
      ],
      imports: [
        BrowserAnimationsModule,
        FormsModule,
        MatCardModule,
        MatInputModule
      ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(AccountCreateComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});

@Directive({selector: '[routerLink]'})
export class RouterLinkDirectiveStub {
  @Input('routerLink') linkParams: any;
  navigatedTo: any = null;

  @HostListener('click')
  onClick() {
    this.navigatedTo = this.linkParams;
  }
}