import React from "react";
import styled, {css} from 'styled-components';
import {
    BrowserRouter as Router,
    Switch,
    Route,
    Link,
    useRouteMatch
  } from "react-router-dom";

export const MenuBar = styled.div `
    scroll:none;
    //position: fixed;
    //top: 0;    
    background-color: #111820;
    height: 60px;
    width:100%;
    display: flex;
    justify-content:center;
    justify-items:center;
    align-content:center;
    align-items:center;
    border-bottom: 1px solid #C4883B;
`

export const MenuItem = ({to, label, activeOnlyWhenExact}) => {
    let match = useRouteMatch({
        path: to,
        exact: activeOnlyWhenExact
      });

    return (
        <StyledMenuItem match={match} to={to}>
           {label}
        </StyledMenuItem>
    )
}

const StyledMenuItem = styled(Link) `
    cursor:pointer;
    text-align:center;
    text-decoration: none;
    padding-left:20px;      
    padding-right:20px;
    font-family: "Exocet Light";
    font-size:20pt;
    color:  ${props => props.match ? "#C19965" : "#8C8C8D"};
    text-transform: uppercase;
    transform: translate(0%);
    transition: 0.3s ease-out;
    &:hover {
        color:white;
    }
`

export const Separator = styled.span`
    border-right: 1px solid #8C8C8D;
    width:1px;
    height:50%;
    padding:0px;
`

export const SubMenuBar = styled(MenuBar) `
    grid-area: submenu;
    height:60px;
    border:none; 
    align-items:flex-end;
    align-content:center;
    vertical-align:middle;
`

export const StyledSubMenuItem = styled(StyledMenuItem) `
    height:70%;
    padding-top:10px;
    vertical-align:middle;
    text-align:center;
    font-size:16pt;
    -padding-top:10px;
    -padding-bottom:10px;
    -height:100%;
    width:180px;
    background-color: ${props => props.match ? "#0D1318" : "#111820"};
`

export const SubMenuItem = ({to, label, activeOnlyWhenExact}) => {
    let match = useRouteMatch({
        path: to,
        exact: activeOnlyWhenExact
      });

    return (
        <StyledSubMenuItem match={match} to={to}>
           {label}
        </StyledSubMenuItem>
    )
}

export const Menu = () => {
    return (
        <React.Fragment>
        <MenuBar>
            <MenuItem active={true}>HOLY GRAIL</MenuItem>
            <Separator/>
            <MenuItem>CHARACTERS</MenuItem>
            <Separator/>
            <MenuItem>ITEM EXCHANGE</MenuItem>
            <Separator/>
            <MenuItem>MF RUNS</MenuItem>
        </MenuBar>
        <SubMenuBar>
            <SubMenuItem active={true}>OVERVIEW</SubMenuItem>
            <SubMenuItem>UNIQUES</SubMenuItem>
            <SubMenuItem>SET ITEMS</SubMenuItem>
            <SubMenuItem>RUNES</SubMenuItem>
        </SubMenuBar>
        </React.Fragment>
    )
}