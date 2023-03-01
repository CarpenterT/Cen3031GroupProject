/// <reference types="cypress" />

describe('Example test', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('Visit login page', () => {
    cy.contains('Login!').click()
    cy.contains('Log In').should('exist')
  })

  it('Visit account creation page', () => {
    cy.contains('Login!').click()
    cy.contains('Create Account').click()
    cy.contains('Create Account').should('exist')
  })

  it('Return from account creation page', () => {
    cy.contains('Login!').click()
    cy.contains('Create Account').click()
    cy.contains('Cancel').click()
    cy.contains('ClusterC').should('exist')
  })
})
