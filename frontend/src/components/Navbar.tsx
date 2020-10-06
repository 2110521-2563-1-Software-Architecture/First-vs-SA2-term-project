import styled from '@emotion/styled'

const NavbarContainer = styled.div`
  background: #232931;
  padding: 0.75rem 1.5rem;
`
const Logo = styled.h2`
  font-size: 36px;
  color: #4ecca3;
  font-family: 'Comfortaa', cursive;
  margin: 0;
`

const Navbar = () => {
  return (
    <NavbarContainer>
      <Logo>matcher</Logo>
    </NavbarContainer>
  )
}

export default Navbar
