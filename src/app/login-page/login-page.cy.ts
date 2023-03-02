/// <reference types="cypress" />
// @ts-check

import { LoginPageComponent } from "./login-page.component";
import { DashboardTestComponent } from "../dashboard-test/dashboard-test.component";
import { async } from "@angular/core/testing";
import { TestBed } from "@angular/core/testing";
import { MatCard } from "@angular/material/card";
import { AppComponent } from "../app.component";

describe('loginPageComponent', () => {
  const sampleUser = 'janeDoe'
  const samplePass = 'password123'
  const sampleUrl = '/login?username=janeDoe&password=password123'

  //trying to get the UI elements to properly render
  //but i'm getting errors out the ass.
  //The below code is from https://stackoverflow.com/a/44508549
  //but I coudln't get it to work.
  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [
        MatCard,
        AppComponent,
        
      ]
    }).compileComponents();
  }));

  it('mounts', () => {
    cy.mount(DashboardTestComponent, {})
    cy.mount(LoginPageComponent, {})
    
    
    cy.get('[data-cy=un]').type(sampleUser)
    cy.get('[data-cy=pw]').type(samplePass)
    //Cypress throws a huge error after clicking.
    //See login-page.components.ts
    cy.get('[data-cy=button]').click()

    cy.get('[data-cy=un]').should('have.text', sampleUser)
    cy.get('[data-cy=pw]').should('have.text', samplePass)

    
  })
})