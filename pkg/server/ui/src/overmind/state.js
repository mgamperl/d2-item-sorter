import { derived } from "overmind";

export const state = {
  filter: {
    quality: "",
    typeQuality: "",
    searchString: "",
    itemSelection: "all",
  },
  itemExchangeFilter: {
    searchString: "",
  },
  characters: {},
  characterList: derived((state) => Object.values(state.characters)),
  selectedCharacterId: "",
  selectedCharacter: derived((state) => state.characters[state.selectedCharacterId] || null),
  sharedStash: {},
  currentPickedUpPageIdx: -1,
  currentPickedUpItemIdx: -1,
  currentPickedUpItem: derived((state) => {
    if (state.sharedStash && state.sharedStash.pages && state.currentPickedUpPageIdx > -1) {
      const page = state.sharedStash.pages[state.currentPickedUpPageIdx];
      if (page && state.currentPickedUpItemIdx > -1) {
        const item = state.sharedStash.pages[state.currentPickedUpPageIdx].items[state.currentPickedUpItemIdx];
        return item;
      }
    }
    return null;
  }),
};
