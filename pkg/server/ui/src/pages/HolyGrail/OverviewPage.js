import { Box, Grid, Card, CardHeader, CardBody, CardFooter, Button, ResponsiveContext } from "grommet";
import * as Icons from "grommet-icons";
import { useHistory } from "react-router-dom";
import React from "react";
import styled, { css } from "styled-components";
import { theme } from "../../theme";
import { HolyGrailFooter } from ".";

import {Content,Footer} from "../Page";

const GridItem = styled.div`
  background-color: #141d24;
  border: 1px solid;
  border-color: #0e1419;
  padding: 10px;
  margin: 5px;
  display: flex;
  flex-direction: column;
  font-size: 12pt;
`;

const ProgressBarDiv = styled.div`
  border: 1px solid ${(props) => theme.global.colors[props.type]};
  width: 100%;
  height: 25px;
  margin: 5px;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const ProgressDiv = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  cursor: ${(props) => (props.clickable ? "pointer" : "default")};
  background-color: ${(props) => theme.global.colors[props.type]};
  opacity: 50%;
  width: ${(props) => props.progress}%;
  height: 25px;
  transform: translate(0%);
  transition: 0.3s ease-out;
  &:hover {
    ${(props) =>
      props.clickable &&
      css`
        opacity: 100%;
      `}
  }
`;

const MissingDiv = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  cursor: ${(props) => (props.clickable ? "pointer" : "default")};
  background-color: ${(props) => "#141D24"};
  opacity: 50%;
  width: ${(props) => 100 - props.progress}%;
  height: 25px;
  transform: translate(0%);
  transition: 0.3s ease-out;
  &:hover {
    ${(props) =>
      props.clickable &&
      css`
        background-color: ${(props) => "#1E2731"};
        opacity: 100%;
      `}
  }
`;

const ProgressLabel = styled.div`
  text-transform: uppercase;
  font-family: ${(props) => theme.global.fontFamily};
  color: ${(props) => theme.global.colors[props.type]};
  min-width: 150px;
`;

const ProgressContainer = styled.div`
  display: flex;
  white-space: nowrap;
  align-items: center;
`;

const ProgressCount = styled.div`
  color: ${(props) => (props.dark ? "white" : "#141D24")};
  justify-content: center;
  text-align: center;
  text-transform: uppercase;
  font-family: ${(props) => theme.global.fontFamily};
`;

const ProgressBar = ({ progress, type, label, count, max, onClickFilled, onClickNotFilled }) => {
  return (
    <React.Fragment>
      <ProgressContainer>
        <ProgressLabel type={type}>{label}</ProgressLabel>
        <ProgressBarDiv type={type}>
          <ProgressDiv clickable={onClickFilled != undefined} onClick={() => onClickFilled && onClickFilled()} progress={progress} type={type}>
            {count > 0 && <ProgressCount>{count}</ProgressCount>}
          </ProgressDiv>
          <MissingDiv clickable={onClickNotFilled != undefined} onClick={() => onClickNotFilled && onClickNotFilled()} progress={progress} type={type}>
            {max - count > 0 && <ProgressCount dark>{max - count}</ProgressCount>}
          </MissingDiv>
        </ProgressBarDiv>
      </ProgressContainer>
    </React.Fragment>
  );
};

const GridItemHeader = styled.h3`
  color: #8c8c8d;
  padding-top: 0px;
  margin-top: 3px;
`;

const ItemProgressBar = ({ items, type, label, filterFunc, onClickFilled, onClickNotFilled }) => {
  const maxCount = items.filter(filterFunc).length;
  const count = items.filter((item) => filterFunc(item) && item.count > 0).length;
  const percent = Math.round((count / maxCount) * 100);

  return <ProgressBar label={label} type={type} progress={percent} count={count} max={maxCount} onClickFilled={onClickFilled} onClickNotFilled={onClickNotFilled} />;
};

export const OverviewPage = ({ items }) => {
  let history = useHistory();

  const uniqueMaxCount = items.filter((item) => item.isUnique).length;
  const uniqueCount = items.filter((item) => item.isUnique && item.count > 0).length;
  const uniquePercent = Math.round((uniqueCount / uniqueMaxCount) * 100);

  const setMaxCount = items.filter((item) => item.isSet).length;
  const setCount = items.filter((item) => item.isSet && item.count > 0).length;
  const setPercent = Math.round((setCount / setMaxCount) * 100);

  const runeMaxCount = items.filter((item) => item.isRune).length;
  const runeCount = items.filter((item) => item.isRune && item.count > 0).length;
  const runePercent = Math.round((runeCount / runeMaxCount) * 100);

  const overallMaxCount = uniqueMaxCount + setMaxCount + runeMaxCount;
  const overallCount = uniqueCount + setCount + runeCount;
  const overallPercent = Math.round((overallCount / overallMaxCount) * 100);

  const responsiveColumns = {
    xsmall: ["1fr"],
    small: ["1fr"],
    medium: ["1fr", "1fr"],
    large: ["1fr", "1fr", "1fr"],
  };

  return (
    <React.Fragment>
      <Content>
        <ResponsiveContext.Consumer>
          {(size) => (
            <Grid
              margin="small"
              gap="small"
              fill="vertical"
              columns={responsiveColumns[size] ? responsiveColumns[size] : ["1fr", "1fr", "1fr"]}
              justify="stretch"
              align="start"
              alignContent="start"
              justifyContent="start"
              rows={["1fr", "1fr", "*"]}
            >
              <Box background="light-blue" pad="small">
                <GridItemHeader>Overall Grail Status</GridItemHeader>
                <ProgressBar
                  label="Uniques"
                  type="unique"
                  progress={uniquePercent}
                  count={uniqueCount}
                  max={uniqueMaxCount}
                  onClickFilled={() => history.push("/holyGrail/unique/all/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/unique/all/missing")}
                />
                <ProgressBar
                  label="Set Items"
                  type="set"
                  progress={setPercent}
                  count={setCount}
                  max={setMaxCount}
                  onClickFilled={() => history.push("/holyGrail/set/all/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/set/all/missing")}
                />
                <ProgressBar
                  label="Runes"
                  type="rune"
                  progress={runePercent}
                  count={runeCount}
                  max={runeMaxCount}
                  onClickFilled={() => history.push("/holyGrail/runes/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/runes/missing")}
                />
                <ProgressBar label="Overall" type="overall" progress={overallPercent} count={overallCount} max={overallMaxCount} />
              </Box>

              <Box background="light-blue" pad="small">
                <GridItemHeader>Unique Grail Status</GridItemHeader>
                <ItemProgressBar
                  items={items}
                  label="Uniques"
                  type="unique"
                  filterFunc={(item) => item.quality === "Unique"}
                  onClickFilled={() => history.push("/holyGrail/unique/all/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/unique/all/missing")}
                />
                <ItemProgressBar
                  items={items}
                  label="Normal"
                  type="unique"
                  filterFunc={(item) => item.quality === "Unique" && item.typeQuality === "Normal"}
                  onClickFilled={() => history.push("/holyGrail/unique/normal/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/unique/normal/missing")}
                />
                <ItemProgressBar
                  items={items}
                  label="Exceptional"
                  type="unique"
                  filterFunc={(item) => item.quality === "Unique" && item.typeQuality === "Exceptional"}
                  onClickFilled={() => history.push("/holyGrail/unique/exceptional/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/unique/exceptional/missing")}
                />
                <ItemProgressBar
                  items={items}
                  label="Elite"
                  type="unique"
                  filterFunc={(item) => item.quality === "Unique" && item.typeQuality === "Elite"}
                  onClickFilled={() => history.push("/holyGrail/unique/elite/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/unique/elite/missing")}
                />
              </Box>

              <Box background="light-blue" pad="small">
                <GridItemHeader>Set Grail Status</GridItemHeader>
                <ItemProgressBar items={items} label="Set Items" type="set" filterFunc={(item) => item.isSet} 
                onClickFilled={() => history.push("/holyGrail/set/all/owned")}
                onClickNotFilled={() => history.push("/holyGrail/set/all/missing")}/>
                <ItemProgressBar items={items} label="Normal" type="set" filterFunc={(item) => item.isSet && item.typeQuality === "Normal"}
                onClickFilled={() => history.push("/holyGrail/set/normal/owned")}
                onClickNotFilled={() => history.push("/holyGrail/set/normal/missing")} />
                <ItemProgressBar items={items} label="Exceptional" type="set" filterFunc={(item) => item.isSet && item.typeQuality === "Exceptional"}
                onClickFilled={() => history.push("/holyGrail/set/exceptional/owned")}
                onClickNotFilled={() => history.push("/holyGrail/set/exceptional/missing")}
                 />
                <ItemProgressBar items={items} label="Elite" type="set" filterFunc={(item) => item.isSet && item.typeQuality === "Elite"}
                  onClickFilled={() => history.push("/holyGrail/set/elite/owned")}
                  onClickNotFilled={() => history.push("/holyGrail/set/elite/missing")} />
              </Box>

              <Box background="light-blue" pad="small">
                <GridItemHeader>Rune Grail Status</GridItemHeader>
                <ItemProgressBar items={items} label="Runes" type="rune" filterFunc={(item) => item.isSet} />
                <ItemProgressBar items={items} label="Low Tier" type="rune" filterFunc={(item) => item.isRune && item.order < 20} />
                <ItemProgressBar items={items} label="Mid Tier" type="rune" filterFunc={(item) => item.isRune && item.order < 30 && item.order >= 20} />
                <ItemProgressBar items={items} label="High Tier" type="rune" filterFunc={(item) => item.isRune && item.order >= 30} />
              </Box>
            </Grid>
          )}
        </ResponsiveContext.Consumer>
      </Content>
      {/*<Footer><HolyGrailFooter /></Footer> */}
    </React.Fragment>
  );
};
