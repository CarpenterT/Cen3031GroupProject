import { DashboardTestComponent } from "./dashboard-test.component";

describe('loginButtonTest.cy.ts', () => {
  it('mounts', () => {
    cy.mount(DashboardTestComponent, {

    })

    cy.get('[data-cy=Lb]').click()
  })
  
})