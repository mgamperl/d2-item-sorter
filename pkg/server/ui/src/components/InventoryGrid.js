import React from "react";
import { Grid, Box, Tip, Text, InfiniteScroll } from "grommet";
import styled from "styled-components";
import { getItemColor } from "./ItemGrid";

const imageUrl = process.env.PUBLIC_URL + "/assets/items/";
const size = 28;

const InventorySlot = styled.div`
  width: 28px;
  height: 28px;
  border: 1px solid white;
`;

const InventoryItemDiv = styled.div.attrs((props) => ({
  style: {
    backgroundImage: `url(${imageUrl + props.item.imageId}.png)`,
    width: `${props.item.type.sizeX * size}px`,
    height: `${props.item.type.sizeY * size}px`,
    left: `${props.item.location.x * size}px`,
    top: `${props.item.location.y * size}px`,
    opacity: `${props.item.etheral ? 50 : 100}%`
  },
}))`
  cursor: pointer;
  width: 100%;
  position: absolute;
  background-color: ${props => props.highlight ? "rgba(4, 73, 0, 0.5)" : "transparent"};
  pointer-events: ${props => props.clickable === false  ?  "none" : "inherit"};
`;

const Inventory = styled.div`
  position: relative;
  width: ${(props) => props.width * size}px;
  height: ${(props) => props.height * size}px;
  border: 1px solid white;
  background-color: #141a20;
  background-image: repeating-linear-gradient(#8c8c8d 0 1px, transparent 1px 100%), repeating-linear-gradient(90deg, #8c8c8d 0 1px, transparent 1px 100%);
  background-size: 28px 28px;
`;

const searchMatches = (item, filter) => {
    if(filter && filter.searchString){
  return filter.searchString != "" && (
    item.name.toLowerCase().indexOf(filter.searchString.toLowerCase()) >= 0
  );
    }else{
        return false;
    }
};

export const InventoryGrid = ({ width, height, items, filter, onItemClick, onGridClick }) => {
  let rows = [];
  let columns = [];
  for (let i = 0; i < width; i++) {
    columns.push("28px");
  }

  for (let i = 0; i < height; i++) {
    rows.push("28px");
  }

  return (
      <Box>
    <Inventory width={width} height={height} onClick={(e) => {
      console.log(e.relatedTarget, e);
      onGridClick && onGridClick()
    }}>
      {items &&
        items.map((item, idx) => (
          <InventoryItem key={idx} item={item} highlight={searchMatches(item, filter)} onItemClick={() => onItemClick && onItemClick(item,idx)} />
        ))}
    </Inventory>
    </Box>
  );
};

export const InventoryItem = ({item,highlight, onItemClick, clickable}) => {
  return (
    <Tip content={<ItemInfo item={item} />}>
      <InventoryItemDiv clickable={clickable} highlight={highlight} item={item} onClick={(e) => {
        console.log("clicking item ", item) 
        if(onItemClick){
          e.stopPropagation();
          onItemClick(); 
        }
      }
    } />
    </Tip>
  )
}

export const ItemInfo = ({item}) => {
  return (<Box>
    <Box ><Text color={getItemColor(item)}>{item.name}</Text></Box>
    <Box ><Text color={getItemColor(item)}>{item.type.name}</Text></Box>
  </Box>);
}

export const Stash = ({ pages, width, height, filter, onItemClick, onGridClick }) => {
  return (
    <Box>
      <Grid columns="280px" rows="auto" pad="small">
        {/*<InfiniteScroll items={pages} step={20}> */}
        {pages &&
          pages.map((page,pageIdx) => (
            <Box key={page.number} width="280px">
              <StashPage filter={filter} key={page.number} width={width} height={height} page={page} onGridClick={(position) => onGridClick && onGridClick(position)} onItemClick={(item,idx) => onItemClick && onItemClick(item,idx,pageIdx)} />
            </Box>
          ))}
        {/*</InfiniteScroll> */}
      </Grid>
    </Box>
  );
};

export const StashPage = ({ page, width, height, filter, onItemClick, onGridClick }) => {
  return (
    <Box pad="small">
      <Text>{page.name || "Stash Page - #" + page.number}</Text>
      <InventoryGrid filter={filter} width={width} height={height} items={page.items} onGridClick={(position) => onGridClick && onGridClick(position)} onItemClick={(item,idx) => onItemClick && onItemClick(item,idx)}/>
    </Box>
  );
};

/*<Grid rows={rows} columns={columns}  >
            {
                columns.map
            }
        </Grid>
        */
