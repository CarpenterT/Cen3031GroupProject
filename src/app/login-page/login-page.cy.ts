/// <reference types="cypress" />
// @ts-check

import { LoginPageComponent } from "./login-page.component";

describe('loginPageComponent', () => {
  const sampleUser = 'janeDoe'
  const samplePass = 'password123'

  it('mounts', () => {
    
    cy.mount(LoginPageComponent, {
      componentProperties: {
        
      },
    })
    cy.get('[data-cy=un]').type(sampleUser)
    cy.get('[data-cy=pw]').type(samplePass)
    
    cy.get('[data-cy=button]').click()

    cy.get('[data-cy=un]').should('have.text', sampleUser)
    cy.get('[data-cy=pw]').should('have.text', samplePass)
  })
})