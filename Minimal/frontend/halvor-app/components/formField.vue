<template>
  <v-container>
    <v-form v-model="valid">
      <v-col cols="12">
        <v-text-field
          v-model="streetName"
          label="Gatenavn"
          required
        ></v-text-field>
      </v-col>
      <v-col cols="12">
        <v-text-field
          v-model="streetNumber"
          label="Gatenummer"
          required
        ></v-text-field>
      </v-col>
      <v-col cols="12">
        <v-text-field v-model="addressLetter" label="Bokstav"></v-text-field>
      </v-col>
      <v-col cols="12">
        <v-text-field
          v-model="zipCode"
          label="Zip code"
          required
        ></v-text-field>
      </v-col>
      <v-col cols="12">
        <v-text-field v-model="city" label="By" required></v-text-field>
      </v-col>
      <v-btn tile color="blue lighten-1" @click="postInformation()"
        >Legg til addresse</v-btn
      >
    </v-form>
  </v-container>
</template>

<script>
export default {
  name: "formFields",
  data() {
    return {
      valid: false,
      streetName: "",
      streetNumber: "",
      addressLetter: "",
      zipCode: "",
      city: "",
      namedRules: [v => !!v || "Field one is required"]
    };
  },
  methods: {
    postInformation() {
      if (this.streetName == "" || this.streetNumber == "") {
        return false;
      }

      let informationObject = {
        information: {
          streetName: this.streetName,
          streetNumber: this.streetNumber,
          addressLetter: this.addressLetter,
          zipCode: this.zipCode,
          city: this.city
        }
      };

      this.streetName = "";
      this.streetNumber = "";
      this.addressLetter = "";
      this.zipCode = "";
      this.city = "";

      console.log("adding information:" + informationObject);
      this.$store.commit("addTravel", informationObject);
      //   this.$axios
      //     .$post("localhost:8080/info", informationObject)
      //     .then(response => {
      //       this.streetName = "";
      //       this.streetNumber = "";
      //       this.addressLetter = "";
      //       this.zipCode = "";
      //       this.city = "";

      //       console.log(response);
      //     })
      //     .catch(error => {
      //       console.log(error);
      //     });
    }
  }
};
</script>
