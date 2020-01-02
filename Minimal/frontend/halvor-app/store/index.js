export const state = {
  data() {
    return {
      travels: [],
      chosenTravel: null
    };
  }
};

export const mutations = {
  addTravel(state, travel) {
    console.log("adding to travels");
    state.travels.unshift(travel);
  },
  getTravels(state) {
    return state.travels;
  },
  deleteTravel(state, index) {
    state.travels.splice(index, 1);
  },
  editAddress(state, informationObject) {
    state.travels[informationObject.idx] = informationObject.obj;
  }
};
