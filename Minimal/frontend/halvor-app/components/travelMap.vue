<template>
  <div>
    {{ travels }}
    <!-- <no-ssr> -->
    <l-map class="mini-map" :zoom="13" :center="center" :draggable="true">
      <l-tile-layer
        url="http://{s}.tile.osm.org/{z}/{x}/{y}.png"
      ></l-tile-layer>
      <div v-for="(dest, idx) in travels" v-bind:key="idx">
        <l-marker :lat-lng="center" :draggable="true">
          <l-popup :content="dest.toString()"></l-popup>
        </l-marker>
      </div>
    </l-map>
    <!-- </no-ssr> -->
  </div>
</template>

<script>
export default {
  name: "travelMap",
  data() {
    return {
      center: [59.3, 11.3],
      travels: [
        [59.4231276, 11.3052155],
        [59.4231276, 11.305216]
      ]
    };
  },
  methods: {
    calculateCenter() {
      if (this.travels.length < 1) {
        this.center = [59.4231276, 11.3052155];
        return;
      }

      if (this.travels.length === 1) {
        this.center = this.travels[0];
        return;
      }

      let x = 0.0;
      let y = 0.0;
      let z = 0.0;

      for (let coordCouple in this.travels) {
        console.log("LAT " + coordCouple[0]);
        let lat = (coordCouple[0] * Math.PI) / 180;

        console.log("LON " + coordCouple[1]);
        let lon = (coordCouple[1] * Math.PI) / 180;

        x += Math.cos(lat) * Math.cos(lon);
        y += Math.cos(lat) * Math.sin(lon);
        z += Math.sin(lat);

        console.log("x: " + x, "y: " + y, "z: " + z);
      }

      let tot = this.travels.length;

      x = x / tot;
      y = y / tot;
      z = z / tot;

      this.center[0] = Math.atan2(y, x);
      let centralSquare = Math.sqrt(x * x + y * y);
      this.center[1] = Math.atan2(z, centralSquare);
    },
    plotAllDestinations() {}
  },
  mounted() {
    this.calculateCenter();
  }
};
</script>

<style src="leaflet/dist/leaflet.css"></style>
<style>
.mini-map {
  width: 100%;
  height: 600px !important;
}
</style>
