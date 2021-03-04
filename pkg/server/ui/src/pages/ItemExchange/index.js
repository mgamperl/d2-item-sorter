import React from "react";


import styled from "styled-components";
import { useDebounce, useMousePosition} from "../../hooks";
import { Box, TextInput, CheckBox, Text, Select, Grid } from "grommet";
import {ItemGrid} from "../HolyGrail/GrailItemView"
import { BrowserRouter as Router, Switch, Route, Redirect, useRouteMatch, Link , useParams} from "react-router-dom";

import { useEffect, useState } from "react";
import * as overmind from "../../overmind";
import { SubMenu, Content, Footer } from "../../pages/Page";

import {InventoryGrid, InventoryItem, Stash} from "../../components"

export const ItemExchangePage = () => {
  let { path, url } = useRouteMatch();
  let [sharedStash, setSharedStash] = useState({ pages: []});

  let state = overmind.useState();
  let actions = overmind.useActions();

  useEffect(() => {
    // GET request using fetch inside useEffect React hook
    fetch("/api/sharedStash")
      .then((response) => response.json())
      .then((data) => {
        setSharedStash(data);
      });

    // empty dependency array means this effect will only run once (like componentDidMount in classes)
  }, []);

  return (
    <React.Fragment>
      <SubMenu></SubMenu>
      <Content>      
         <PickedUpItem sharedStash={sharedStash}  />
        <Stash filter={state.itemExchangeFilter} width={10} height={10} pages={sharedStash.pages} onGridClick={(position) => actions.dropItem({position})} onItemClick={(item,idx,pageIdx) => actions.pickupItem({item,idx,pageIdx})}/>
      </Content>
      <Footer>
        <ItemExchangeFooter />
      </Footer>     
    </React.Fragment>
  );
};

const PickedUpItem = ({sharedStash}) => {

  let [currentPickedUpItem, setCurrentPickedUpItem] = useState(null);

  let reaction = overmind.useReaction();

  useEffect(() => reaction(
    ({ currentPickedUpPageIdx, currentPickedUpItemIdx }) => ({currentPickedUpPageIdx,currentPickedUpItemIdx}),
    ({ currentPickedUpPageIdx, currentPickedUpItemIdx }) => {
      
      if (sharedStash && sharedStash.pages && currentPickedUpPageIdx > -1) {
        const page = sharedStash.pages[currentPickedUpPageIdx];
        if (page && currentPickedUpItemIdx > -1) {
          const item = sharedStash.pages[currentPickedUpPageIdx].items[currentPickedUpItemIdx];
          console.log(currentPickedUpPageIdx, currentPickedUpItemIdx, item);
          return setCurrentPickedUpItem(item);          
        }
      }
      return setCurrentPickedUpItem(null);
    }
  ))

  if(currentPickedUpItem){
  return (
    <FollowMouse><InventoryItem clickable={false} item={currentPickedUpItem} /></FollowMouse>
  )
  }else{
    return null;
  }
}

const FollowMouse = ({children}) => {

  const { clientX, clientY, pageX, pageY, } = useMousePosition();

  return <FollowMouseDiv x={clientX} y={clientY} >{children && children}</FollowMouseDiv>
}

const FollowMouseDiv = styled.div.attrs((props) => ({
  style: {
    top: props.y,
    left: props.x,
  },
}))` 
position:fixed;
z-index:9999;
`

export const ItemExchangeFooter = () => {
  const state = overmind.useState();
  const actions = overmind.useActions();

  const [value, setValue] = useState(state.filter.searchString);
  const [checked, setChecked] = React.useState(true);
  const debouncedSearchTerm = useDebounce(value, 500);

  // Here's where the API call happens
  // We use useEffect since this is an asynchronous action
  useEffect(
    () => {
      // Make sure we have a value (user has entered something in input)
      console.log(debouncedSearchTerm);
      actions.updateItemExchangeFilter({
        searchString: debouncedSearchTerm,
      });
    },
    // This is the useEffect input array
    // Our useEffect function will only execute if this value changes ...
    // ... and thanks to our hook it will only change if the original ...
    // value (searchTerm) hasn't changed for more than 500ms.
    [debouncedSearchTerm]
  );

  return (
    <Box direction="row" gap="xsmall" pad="xsmall" align="center">
      <Text>Search:</Text>
      <TextInput placeholder="type here" value={value} onChange={(event) => setValue(event.target.value)} />      
      {/*<Select
        options={[
          { value: "unique", label: "Unique" },
          { value: "set", label: "Set Items" },
        ]}
        placeholder="Type"
        labelKey="label"
        valueKey={{ key: 'value', reduce: true }}
        defaultValue="all"
        size="small"
        clear={true}
        labelKey="label"
        value={state.filter.quality}
        onChange={(x) => actions.updateItemFilter({quality: x.value})}
      /> */}
    </Box>
  );
};
