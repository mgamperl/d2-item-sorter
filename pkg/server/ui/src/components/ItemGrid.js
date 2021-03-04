
import React from "react"; 
import {Text} from "grommet";
import styled from "styled-components"
import { theme } from "../theme";

const ItemGridEntryImageContainer = styled.div `
width: 80%;
height: 80%;           
align-items: center;
display: flex;
justify-content: center;
align-self: center;
`

const ItemName = styled.div `
color: ${props => cols[props.type || "default"]};
text-transform: uppercase;
font-family: "Exocet Light";
font-size: 12pt;
text-align: center;
`

const ItemDescription = styled.div `
color: ${props => cols[props.type || "default"]};
font-size: 8pt;
font-family: "Exocet Light";
text-transform: uppercase;
text-align: center;
`

export const getItemColor = (item) => cols[getItemType(item)];

const getItemType = (item) => item.quality.toLowerCase() || "default"

const cols = {
    unique: "#C19965",
    set: "#357E24",
    rune: "#8C8C8D",
    default: "#8C8C8D",
  }

const sizes= {
   unique: "200px",
   set: "200px",
   rune: "100px",
   normal: "200px",
   default: "200px"
}
  
const ItemGridEntryDiv = styled.div `
    position:relative;
    background-color: ${props => props.found ? "#141D24" : "#0F161B"};
    border: 1px solid;
    border-color: ${props => props.found ? theme.global.colors[props.type] : "#0E1419"};
    padding:10px;
    margin: 5px;
    display: flex;
    flex-direction: column;
    font-size: 12pt;
    width: ${props =>sizes[props.type] || "200px" };
    height: ${props =>sizes[props.type] || "200px"} ;
  `

const CountDisplay = styled.div `
  position:absolute;
  top:10px;
  left:10px;
  right:10px;
  bottom:10px;
  display:flex;
  align-items:flex-end;
  justify-content:flex-end;
  justify-items: center;
  align-content:center;
  color:white;
  opacity:15%;
  text-align:center;
  font-size:24pt;
`

const CountText = styled(Text) `
font-size:24pt;
`

const CountTimes = styled(Text) `
font-size:12pt;
`

export const ItemGridEntry = ({ item }) => {
    return (      
      <ItemGridEntryDiv
        type={getItemType(item)}
        found={item.count <= 0 ? false : true}
        key={item.imageId}
      >
        {item.count > 0 && <CountDisplay><CountTimes>x</CountTimes><CountText>{item.count}</CountText></CountDisplay> }
        <ItemName type={getItemType(item)}>{item.name}</ItemName>
        {(item.isUnique || item.isSet) && <ItemDescription type={getItemType(item)}>
          {item.typeName} ({item.typeQuality})
        </ItemDescription> }
        <div
          style={{
            width: "100%",
            height: "100%",
            alignItems: "center",
            display: "flex",
            justifyContent: "center",
            alignSelf: "center",
          }}
        >
          <ItemGridEntryImageContainer>
            <img
              src={
                process.env.PUBLIC_URL + "/assets/items/" + item.imageId + ".png"
              }
            />
          </ItemGridEntryImageContainer>
        </div>
      </ItemGridEntryDiv>
    );
  };