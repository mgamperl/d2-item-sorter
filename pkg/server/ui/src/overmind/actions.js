import { Actions } from "grommet-icons";


export const updateItemFilter = ({state}, value) => {    
    state.filter = {...state.filter, ...value};
}

export const updateItemExchangeFilter = ({state}, value) => {    
    state.itemExchangeFilter = {...state.itemExchangeFilter, ...value};
}

export const setCharacters = ({state}, value) => {
    state.characters = value;
}

export const setSelectedCharacterId = ({state}, value) => {
    state.selectedCharacterId = value;
}

export const setSharedStash = ({state}, value) => {
    state.sharedStash = value;
}

export const pickupItem = ({state,actions}, value) => {

    if(state.currentPickedUpItem !== null){
        state.currentPickedUpPageIdx = value.pageIdx;
        state.currentPickedUpItemIdx = value.idx;
        console.log("item picked up ", value);
    }else{
        actions.switchItem(value);
    }
}

export const switchItem = ({state,actions},value) => {
    state.currentPickedUpPageIdx = value.pageIdx;
    state.currentPickedUpItemIdx = value.idx;
}

export const dropItem = ({state,actions},value) => {
    state.currentPickedUpPageIdx = -1;
    state.currentPickedUpItemIdx = -1;
}