<template>
  <section class="relative container mx-auto py-4">
    <div
      class=" bg-white min-h-1/2 mt-6 rounded shadow-md px-12 py-12"
      
    >
      <div v-html="md"></div>
      <div v-show="!md" class="relative w-full h-full flex justify-center items-center">
        <div class="inline-flex">
        <input
          type="text"
          v-model="key"
          placeholder="Enter decryption key..."
          id="key"
          class="py-4 px-2 bg-white text-grey-darkest rounded-l shadow-md focus:shadow-outline"
        >
        <input
          type="submit"
          value="Decrypt"
          @click.prevent="decrypt"
          class="px-4 py-4 bg-teal-darker rounded-r text-white"
        >
        </div>
      </div>
    </div>
  </section>
</template>

<script>
let CryptoJS = require("crypto-js");
import marked from "marked";
import axios from "~/plugins/axios";

import Navbar from "~/components/Navbar.vue";
export default {
    components: {
        Navbar,
    },
  asyncData({ params, $axios }) {
    return $axios
      .get(`/api/pad/${params.id}`)
      .then(res => {
        return { encrypted: res.data.data };
      });
  },
  data: function() {
    return {
      key: "",
      decrypted: ""
    };
  },
  computed: {
    md: function() {
      return marked(this.decrypted);
    }
  },
  methods: {
    decrypt() {
      var bytes = CryptoJS.AES.decrypt(this.encrypted, this.key);
      this.decrypted = bytes.toString(CryptoJS.enc.Utf8);
    }
  }
};
</script>

