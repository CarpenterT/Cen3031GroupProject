describe('Basic Home Test', () => {
  it('Checks if the home page loads', function () {
    cy.visit('http://localhost:4200')
    cy.contains("Login!").click()
    cy.url().should('include', 'http://localhost:4200/login')
    
    cy.contains('Create Account').click()
    cy.url().should('include', 'http://localhost:4200/account-create')

    cy.contains('Username').type('abc')
    cy.contains('Password').type('123')

    cy.get('form').submit()

    cy.on('window:alert', (t)=>{
      expect(t).toBe.contains('Username is already taken.')
    })
  })
})
