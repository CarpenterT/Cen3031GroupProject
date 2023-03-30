import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { NavTestComponent } from './nav-test/nav-test.component';
import { LayoutModule } from '@angular/cdk/layout';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { DashboardTestComponent } from './dashboard-test/dashboard-test.component';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatCardModule } from '@angular/material/card';
import { MatMenuModule } from '@angular/material/menu';
import { LoginPageComponent } from './login-page/login-page.component';
import { RouterModule } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { AccountCreateComponent } from './account-create/account-create.component';
import { ErrorPageComponent } from './error-page/error-page.component';
import { HomePageComponent } from './home-page/home-page.component';
import { ChatBoxTestComponent } from './chat-box-test/chat-box-test.component';



@NgModule({
  declarations: [
    AppComponent,
    NavTestComponent,
    DashboardTestComponent,
    LoginPageComponent,
    AccountCreateComponent,
    ErrorPageComponent,
    HomePageComponent,
    ChatBoxTestComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    HttpClientModule,
    RouterModule.forRoot([
      { path: '', component: DashboardTestComponent},
      { path: 'login', component: LoginPageComponent},
      { path: 'account-create', component: AccountCreateComponent },
      { path: 'home', component: HomePageComponent },
      { path: 'chat-page', component: ChatBoxTestComponent },
      //Not sure if below line is working. Might be related to Cypress error.
      { path: 'login?**', component: LoginPageComponent},

      { path: '**', component: ErrorPageComponent },
    ]),
    FormsModule,
    MatSlideToggleModule,
    LayoutModule,
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatIconModule,
    MatListModule,
    MatGridListModule,
    MatCardModule,
    MatMenuModule,
    MatInputModule,
    HttpClientModule,
    
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
