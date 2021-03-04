import React from "react";

import styled from "styled-components";
import { Box, TextInput, CheckBox, Text, Select, Grid, Button } from "grommet";
import { BrowserRouter as Router, Switch, Route, Redirect, useHistory, useRouteMatch, Link, useParams } from "react-router-dom";

import { useEffect, useState } from "react";
import * as overmind from "../../overmind";
import { SubMenu, Content } from "../../pages/Page";

import { InventoryGrid, Stash } from "../../components";

const POSTRequestOptions = {
  method: "POST",
  headers: { "Content-Type": "application/json" },
};

export const RunCounterPage = () => {
  const [recorderRunning, setRecorderRunning] = useState(false);

  let { path, url } = useRouteMatch();

  let state = overmind.useState();
  let actions = overmind.useActions();

  function startRunRecorder() {
    setRecorderRunning(true);
    fetch("/api/startRunRecorder", POSTRequestOptions)
      .then((response) => response.json())
      .then((data) => {
        setRecorderRunning(true);
      });
  }

  function stopRunRecorder() {
    fetch("/api/stopRunRecorder", POSTRequestOptions)
      .then((response) => response.json())
      .then((data) => {
        setRecorderRunning(false);
      });
  }

  return (
    <React.Fragment>
      <SubMenu></SubMenu>
      <Content>
        <Grid pad="small" gap="small" margin="small">
            <Box>
          {recorderRunning && <Box >Running</Box>}
          {!recorderRunning && <Box >Stopped</Box>}
          </Box>
          <Box direction="row">
            <Button onClick={() => startRunRecorder()} primary={!recorderRunning} secondary={recorderRunning} label="Start Recorder" />
            <Button onClick={() => stopRunRecorder()} secondary={!recorderRunning} primary={recorderRunning} label="Stop Recorder" />
          </Box>
        </Grid>
      </Content>
    </React.Fragment>
  );
};
