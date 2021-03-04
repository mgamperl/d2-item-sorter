import React from "react";
import styled from "styled-components";

export const GridLayout = styled.div`
  display: grid;
  height: 100vh;
  grid-template-columns: 50px 50px 50px 50px;
  grid-template-rows: auto auto 1fr auto;
  grid-template-areas:
    "header"
    "submenu"
    "main"
    "footer";
`;

export const Content = styled.div`
  grid-area: main;
  width: 100vw;
  overflow: auto;
  background-color: #0d1318;
`;

export const Menu = styled.div`
  grid-area: header;
  width: 100vw;
`;

export const SubMenu = styled.div`
  grid-area: submenu;
  width: 100vw;
`;

export const Footer = styled.div`
  border-top: 1px solid #C4883B;
  width: 100vw;
  grid-area: footer;
`;

export const Page = ({subMenu, content, footer}) => {
  return (
    <React.Fragment>
      <SubMenu>
          {subMenu}
      </SubMenu>
      <Content>
          {content}
      </Content>
      <Footer>
          {footer}
      </Footer>
    </React.Fragment>
  );
};
