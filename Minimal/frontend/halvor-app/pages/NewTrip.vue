<template>
  <v-container>
    <v-row>
      <v-col cols="6">
        <form-field></form-field>
      </v-col>

      <v-col cols="6">
        <list-trips></list-trips>
      </v-col>
    </v-row>
    <v-btn tile @click="showMap()" color="blue lighten-2"
      >Vis kart for kjøretur</v-btn
    >

    <div v-if="showMapBoolean">
      <travel-map></travel-map>
    </div>
  </v-container>
</template>

<script>
import formField from "../components/formField.vue";
import listTrips from "../components/trips.vue";
import travelMap from "../components/travelMap.vue";
export default {
  name: "newTrip",
  components: {
    formField,
    listTrips,
    travelMap
  },
  data() {
    return {
      showMapBoolean: false
    };
  },
  methods: {
    showMap() {
      this.showMapBoolean = true;
    }
  }
};
</script>
