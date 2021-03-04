import React from "react";

import { Box, TextInput, CheckBox, Text, Select } from "grommet";

import { BrowserRouter as Router, Switch, Route, Redirect, useRouteMatch } from "react-router-dom";

import { useDebounce } from "../../hooks";
import { useEffect, useState } from "react";
import { OverviewPage } from "./OverviewPage";
import { GrailItemView } from "./GrailItemView";

import * as overmind from "../../overmind";

import { SubMenuBar, SubMenuItem } from "../../components/Menu";
import { SubMenu } from "../Page";

export const HolyGrailFooter = () => {
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
      actions.updateItemFilter({
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
      <Select
        options={["all","missing","already owned"       
        ]}
        defaultValue="all"
        size="small"       
        clear={true}
        value={state.filter.itemSelection}
        onChange={({ option }) => actions.updateItemFilter({ itemSelection: option })}
      />
      <Select
      options={[
        { value: "normal", label: "Normal" },
        { value: "exceptional", label: "Exceptional" },
        { value: "elite", label: "Elite" },
      ]}
      placeholder="Quality"
      labelKey="label"
      valueKey={{ key: 'value', reduce: true }}
      size="small"
      clear={true}
      labelKey="label"
      value={state.filter.typeQuality}
      onChange={(x) => actions.updateItemFilter({typeQuality: x.value})}  
      />
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

export const HolyGrailSubMenu = () => {
  let { path, url } = useRouteMatch();

  return (
    <SubMenuBar>
      <SubMenuItem to={`${url}/overview`} label="OVERVIEW"></SubMenuItem>
      <SubMenuItem to={`${url}/unique`} label="UNIQUES"></SubMenuItem>
      <SubMenuItem to={`${url}/set`} label="SET ITEMS"></SubMenuItem>
      <SubMenuItem to={`${url}/rune`} label="RUNES"></SubMenuItem>
    </SubMenuBar>
  );
};

export const HolyGrailPage = () => {
  let { path, url } = useRouteMatch();

  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    // GET request using fetch inside useEffect React hook
    setLoading(true);
    fetch("/api/grailStatus")
      .then((response) => response.json())
      .then((data) => {
        setItems(data);
        setLoading(false);
      });

    // empty dependency array means this effect will only run once (like componentDidMount in classes)
  }, []);

  return (
    <React.Fragment>
      <SubMenu>
        <HolyGrailSubMenu />
      </SubMenu>
      <Switch>
        <Route exact path={path}>
          <Redirect to={`${url}/overview`} />
        </Route>
        <Route path={`${path}/overview`}>
          <OverviewPage items={items} />
        </Route>
        <Route path={`${path}/:quality/:typeQuality?/:owned?`} children={<GrailItemView items={items} />}></Route>
      </Switch>
    </React.Fragment>
  );
};
