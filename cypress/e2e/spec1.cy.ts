/// <reference types="cypress" />

describe('Navigation', () => {
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
    cy.contains('Create an Account').should('exist')
  })

  it('Return from account creation page', () => {
    cy.contains('Login!').click()
    cy.contains('Create Account').click()
    cy.contains('Home').click()
    cy.contains('ClusterC').should('exist')
  })
})

describe('Log In', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('Invalid login', () => {
    cy.contains('Login!').click()
    cy.contains('Username').click().type('user')
    cy.contains('Password').click().type('1234')
    cy.contains('Login').click()
    cy.contains('Log In').should('exist')
  })

  /*
  it('Create account and log in', () => {
    cy.contains('Login!').click()
    cy.contains('Create Account').click()
    cy.contains('Username').click().type('user')
    cy.contains('Password').click().type('1234')
    cy.contains('Create').click()
    cy.contains('Welcome, user!').should('exist')
  })
  */
})
