import React from "react";
import { theme } from "./theme";
import "./App.css";
import styled from "styled-components";
import { BrowserRouter as Router, Switch, Route, Redirect } from "react-router-dom";

import { Menu, Footer, GridLayout, SubMenu, Content } from "./pages/Page";
import { HolyGrailPage, HolyGrailSubMenu, HolyGrailFooter } from "./pages/HolyGrail";
import { CharacterPage } from "./pages/Characters";
import { useEffect, useState } from "react";
import { MenuItem, MenuBar, Separator } from "./components/Menu";
import { ItemExchangePage } from "./pages/ItemExchange";
import { RunCounterPage } from "./pages/RunCounter";

function App() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(false);
  const [sortingSharedStash, setSortingSharedStash] = useState(false);


  function sortSharedStash() {
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
    };

    setSortingSharedStash(true);
    fetch("/api/sharedStash/sort", requestOptions)
      .then((response) => response.json())
      .then(() => setSortingSharedStash(false));
  }

  const routes = {
    "/": {
      redirect: "/holyGrail",
      exact: true,
    },
    "/holyGrail": {
      label: "Holy Grail",
      subMenu: <HolyGrailSubMenu />,
      page: <HolyGrailPage />,
      footer: <HolyGrailFooter />,
    },
    "/characters": {
      label: "Characters",
      page: <CharacterPage />,
      footer: <div>Characters footer</div>,
    },
    "/itemExchange": {
      label: "Item Exchange",
      page: <ItemExchangePage />,
    },
    "/mfRuns": {
      label: "Mf Runs",
      page: <RunCounterPage />,
    },
  };

  return (
    <Router>
      <GridLayout>
        <Menu>
          <MenuBar>
            {Object.keys(routes).map((k, idx) => {
              let entry = routes[k];
              let lastEntry = idx == Object.keys(routes).length - 1;
              let isRedirect = entry.redirect != undefined;

              return [<MenuItem key={k} to={k} label={entry.label} />, lastEntry || isRedirect ? null : <Separator key={idx} />];
            })}
          </MenuBar>
        </Menu>
        {/*<SubMenu>
          <Switch>
            {Object.keys(routes).map((k, idx) => {
              let entry = routes[k];
              let lastEntry = idx == Object.keys(routes).length - 1;
              let isRedirect = entry.redirect != undefined;

              if (isRedirect) {
                return (
                  <Route exact={entry.exact} path={k}>
                    <Redirect to={entry.redirect} />
                  </Route>
                );
              }

              if (entry.subMenu) {
                return <Route path={k}>{entry.subMenu}</Route>;
              } else {
                return null;
              }
            })}
          </Switch>
        </SubMenu>*/}
        <Switch>
          {Object.keys(routes).map((k, idx) => {
            let entry = routes[k];
            let lastEntry = idx == Object.keys(routes).length - 1;
            let isRedirect = entry.redirect != undefined;

            if (isRedirect) {
              return (
                <Route key={k} exact={entry.exact} path={k}>
                  <Redirect to={entry.redirect} />
                </Route>
              );
            }

            if (entry.page) {
              return (
                <Route key={k} path={k}>
                  {entry.page}
                </Route>
              );
            } else {
              return null;
            }
          })}
        </Switch>
      </GridLayout>
    </Router>
  );
}

export default App;

/*

<Switch>
            {Object.keys(routes).map((k, idx) => {
                let entry = routes[k];
                let lastEntry = idx == Object.keys(routes).length - 1;
                let isRedirect = entry.redirect != undefined;
  
                if (isRedirect) {
                  return (
                    <Route exact={entry.exact} path={k}>
                      <Redirect to={entry.redirect} />
                    </Route>
                  );
                }
  
                if (entry.footer) {
                  return <Route path={k}>{entry.footer}</Route>;
                } else {
                  return null;
                }
              })}
            </Switch>

            */
