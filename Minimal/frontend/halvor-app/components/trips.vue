<template>
  <v-container>
    <v-list>
      <template v-for="(travel, index) in $store.state.travels">
        <v-list-item v-bind:key="index">
          <v-list-item-icon>
            <v-icon>mdi-map-marker</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title
              >{{ travel.information.streetName }}
              {{
                travel.information.streetNumber +
                  travel.information.addressLetter
              }}
            </v-list-item-title>
            {{ travel.information.zipCode }}, {{ travel.information.city }}
          </v-list-item-content>

          <v-btn @click="$store.commit('deleteTravel', index)">slett</v-btn>
          <v-btn @click="openAddressEdit(travel, index)">endre</v-btn>
        </v-list-item>
      </template>
    </v-list>

    <v-dialog v-model="changeAddress" v-if="changeAddressBool">
      <v-card>
        <v-card-title>
          Endre addresse
        </v-card-title>

        <v-card-text>
          <v-form>
            <v-col cols="12">
              <v-text-field
                v-model="streetName"
                :label="changeAddress.information.streetName"
                required
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="streetNumber"
                :label="changeAddress.information.streetNumber"
                required
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="addressLetter"
                :label="changeAddress.information.addressLetter"
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="zipCode"
                :label="changeAddress.information.zipCode"
                required
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="city"
                :label="changeAddress.information.city"
                required
              ></v-text-field>
            </v-col>
            <v-btn tile color="blue lighten-1" @click="editAddress()"
              >Endre addresse</v-btn
            >
            <v-btn tile color="red lighten-1" @click="changeAddressBool = false"
              >Avbryt</v-btn
            >
          </v-form>
        </v-card-text>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script>
export default {
  name: "listTrips",
  data() {
    return {
      travels: [],
      changeAddress: {},
      streetName: "",
      streetNumber: "",
      city: "",
      zipCode: "",
      addressLetter: "",
      idx: 0,
      changeAddressBool: false
    };
  },
  methods: {
    fetchTravels() {
      // get all travels from database.
      // this.$axios.$get("localhost:8080/allTravels").then().catch();
      return true;
    },
    openAddressEdit(travel, index) {
      this.changeAddress = travel;

      this.idx = index;
      this.changeAddressBool = true;
    },
    editAddress() {
      if (this.streetName == "" || this.streetName == NaN)
        this.streetName = this.changeAddress.information.streetName;
      if (this.streetNumber == "" || this.streetNumber == NaN)
        this.streetNumber = this.changeAddress.information.streetNumber;
      if (this.addressLetter == "" || this.addressLetter == NaN)
        this.addressLetter = this.changeAddress.information.addressLetter;
      if (this.zipCode == "" || this.zipCode == NaN)
        this.zipCode = this.changeAddress.information.zipCode;
      if (this.city == "" || this.city == NaN)
        this.city = this.changeAddress.information.city;

      let informationObject = {
        information: {
          streetName: this.streetName,
          streetNumber: this.streetNumber,
          addressLetter: this.addressLetter,
          zipCode: this.zipCode,
          city: this.city
        }
      };

      this.$store.commit("editAddress", {
        obj: informationObject,
        idx: this.idx
      });
      //   this.travels[this.idx] = informationObject;
      this.changeAddressBool = false;
      this.changeAddress = {};
    }
  }
};
</script>
