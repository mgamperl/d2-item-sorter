import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import * as overmind from "../../overmind";
import { ItemGridEntry } from "../../components/ItemGrid";
import { Content, Footer } from "../Page";
import { HolyGrailFooter } from ".";

import { Grid, InfiniteScroll, Box, Text } from "grommet";

import { generatePath, useHistory, useRouteMatch } from "react-router-dom";

const useParamsAsState = () => {
  const { path, params } = useRouteMatch();
  const history = useHistory();

  const updateParams = (updatedParams) => {
    Object.assign(params, updatedParams);
    history.push(generatePath(path, params));
  };
  return [params, updateParams];
};

export const GrailItemView = ({ items }) => {
  const state = overmind.useState();
  const actions = overmind.useActions();

  const { quality, typeQuality, owned } = useParams();

  useEffect(() => {
    let filter = {};
    if (quality !== "" || typeQuality != "" || owned != "") {
      if (owned === "owned") {
        filter = { ...filter, itemSelection: "already owned" };
      }

      if (owned === "missing") {
        filter = { ...filter, itemSelection: "missing" };
      }

      if (typeQuality === "all") {
        filter = { ...filter, typeQuality: "" };
      }

      filter = { ...filter, quality, typeQuality };

      actions.updateItemFilter(filter);
    }
  }, [quality, typeQuality, owned]);

  const shownItems = items.filter((item) => {
    if (state.filter.itemSelection === "missing" && item.count > 0) {
      return false;
    }

    if (state.filter.itemSelection === "already owned" && item.count <= 0) {
      return false;
    }

    if (state.filter.typeQuality !== undefined && (state.filter.typeQuality !== "") & (state.filter.typeQuality !== "all")) {
      if (state.filter.typeQuality.toLowerCase() !== item.typeQuality.toLowerCase()) {
        return false;
      }
    }

    if (state.filter.searchString !== undefined) {
      if (item.name.toLowerCase().indexOf(state.filter.searchString.toLowerCase()) < 0 && item.typeName.toLowerCase().indexOf(state.filter.searchString.toLowerCase()) < 0) {
        return false;
      }
    }

    if (state.filter.quality !== undefined && state.filter.quality !== "") {
      if (item.quality.toLowerCase() !== state.filter.quality.toLowerCase()) {
        return false;
      }
    }

    return true;
  });

  return (
    <React.Fragment>
      <Content>
        <ItemGrid items={shownItems} />
      </Content>
      <Footer>
        <HolyGrailFooter />
      </Footer>
    </React.Fragment>
  );
};

export const ItemGrid = ({items}) => {
  return <Grid columns="220px" rows="auto" pad="small">
  <InfiniteScroll items={items} step={50}>
    {(item) => (
      <Box key={item.id} width="200px">
        <ItemGridEntry item={item} />
      </Box>
    )}
  </InfiniteScroll>
</Grid>
}
