import React from "react";

import styled from "styled-components"
import { Box, TextInput, CheckBox, Text, Select, Grid } from "grommet";
import {ItemGrid} from "../HolyGrail/GrailItemView"
import { BrowserRouter as Router, Switch, Route, Redirect, useHistory, useRouteMatch, Link , useParams} from "react-router-dom";

import { useEffect, useState } from "react";
import * as overmind from "../../overmind";
import { SubMenu, Content } from "../../pages/Page";

import {InventoryGrid, Stash} from "../../components"

export const CharacterPage = () => {
  let { path, url } = useRouteMatch();

  let state = overmind.useState();
  let actions = overmind.useActions();

  useEffect(() => {
    // GET request using fetch inside useEffect React hook
    fetch("/api/characters")
      .then((response) => response.json())
      .then((data) => {
        actions.setCharacters(data);
      });

    // empty dependency array means this effect will only run once (like componentDidMount in classes)
  }, []);

  return (
    <React.Fragment>
      <SubMenu></SubMenu>
      <Switch>
        <Route exact path={path}>
          <CharacterOverview characters={state.characterList} />
        </Route>
        <Route path={`${path}/:name`}>
          <CharacterDetails />
        </Route>
      </Switch>
    </React.Fragment>
  );
};

export const CharacterOverview = ({ characters }) => {
  return (
    <Content>
      <Grid pad="small" gap="small" margin="small">
        {characters.map((char) => <CharacterItem key={char.name} character={char} />)}
      </Grid>
    </Content>
  );
};

export const CharacterDetails = () => {

    const state = overmind.useState();
    const actions = overmind.useActions();

    const { name } = useParams();

    useEffect(() => {
        console.log("set selected character id ", name);
          actions.setSelectedCharacterId(name);
        }, [name]);

  return <Content><Grid pad="small">
      {state.selectedCharacter && (<Box>
          <Text>{state.selectedCharacter.name}</Text>
          <Stash width={10} height={10} pages={state.selectedCharacter.personalStash.pages} />
         <InventoryGrid width={10} height={4} items={state.selectedCharacter.items.filter((item) => item.location.location === "inventory")} />
      </Box>)}
  </Grid>
  </Content>;
};

export const CharacterItem = ({character}) => {

    let history = useHistory();

    return (
        <Box>
            <CharacterItemDiv onClick={() => history.push("/characters/"+ character.name)}>
                <Text >{character.name}</Text>
                <Text>{character.class}</Text>
                <Text>Level {character.level}</Text>
            </CharacterItemDiv>            
        </Box>
    );
}

const CharacterItemDiv = styled.div `
    cursor:pointer;
    position:relative;
    background-color: #141D24;
    border: 1px solid;
    border-color: #C19965;
    padding:10px;
    margin: 5px;
    display: flex;
    flex-direction: column;
    font-size: 12pt;
  `