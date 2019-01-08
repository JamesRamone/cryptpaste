<template>
  <section class="relative container mx-auto py-4">
    <div class="pt-12 pb-24">
      <h1 class="mb-12 font-bold text-4xl text-white">Share your thoughts easily and securely</h1>
      <markdown-editor v-model="text" ref="markdownEditor" :configs="mdeconf"></markdown-editor>
      <div class="flex flex-col-reverse md:flex-row mt-8 align-middle items-end justify-between">
        <div class="mt-6 md:mt-0 flex items-stretch flex-col">
          <div class="flex justify-between">
            <div>
            <input v-show="submitted" type="checkbox" name="addkey" v-model="addkey" class="">
            <label v-show="submitted">Append key to link</label>
            </div>
          </div>
          <div class="bg-white rounded px-4 py-4 shadow">
            <nuxt-link v-show="submitted" :to="relurl" class=" no-underline text-teal-dark hover:text-teal-darker mt-4">{{ link }}</nuxt-link>
            <span v-show="!link" class="text-grey-dark">{{linktext}}</span>
            </div>
        </div>
        <div class="inline-flex shadow-md">
          <input
            type="text"
            v-model="pass"
            placeholder="Enter encryption key..."
            id="key"
            class="py-4 px-2 bg-white text-grey-darkest rounded-l "
          >
          <button
            @click.prevent="submitData"
            class="px-4 py-4 bg-teal-dark rounded-r text-white  flex items-center"
            v-bind:class="{ disabled: !enableSubmit }"
            :disabled="!enableSubmit"
          ><svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" class="h-4 mr-2 fill-current text-white"><path d="M224 420c-11 0-20-9-20-20v-64c0-11 9-20 20-20s20 9 20 20v64c0 11-9 20-20 20zm224-148v192c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V272c0-26.5 21.5-48 48-48h16v-64C64 71.6 136-.3 224.5 0 312.9.3 384 73.1 384 161.5V224h16c26.5 0 48 21.5 48 48zM96 224h256v-64c0-70.6-57.4-128-128-128S96 89.4 96 160v64zm320 240V272c0-8.8-7.2-16-16-16H48c-8.8 0-16 7.2-16 16v192c0 8.8 7.2 16 16 16h352c8.8 0 16-7.2 16-16z"/></svg>
          Encrypt & store
          </button>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import markdownEditor from "vue-simplemde/src/markdown-editor";
import Navbar from "~/components/Navbar.vue";
// import axios from "~/plugins/axios";

let CryptoJS = require("crypto-js");

export default {
  components: {
    markdownEditor,
    Navbar
  },
  data: function() {
    return {
      text: "",
      pass: "",
      uuid: "",
      addkey: false,
      submitted: false,
      mdeconf: {
        toolbar: false,
        status: false,
        placeholder: "# Share your thoughts privately"
      }
    };
  },
  computed: {
    encrypted: function() {
      return CryptoJS.AES.encrypt(this.text, this.pass).toString();
    },
    link: function() {
      return this.uuid !== "" ? "https://cryptpaste.xyz" + this.relurl : "";
    },
    relurl: function() {
      
      return this.addkey ? this.suffixkey : this.suffix;
   
    },
    suffix: function() {
      
      return this.uuid !== "" ? "/pad/" + this.uuid  : "";
 
    },
    suffixkey: function(){
      if (this.submitted) {
        return this.suffix + "$" + this.pass;
      }
    },
    enableSubmit: function() {
      return (this.text.length && this.pass.length) && !this.link.length;
    },
    linktext: function() {
      if (this.link === "") {
        return "Store your note to get link."
      } else {
        return "Your link is ready:"
      }
    }
  },
  methods: {
    submitData() {
        this.$axios
          .post("https://api.cryptpaste.xyz/pad", { data: this.encrypted })
          .then(response => {
            console.log(response);
            this.uuid = response.data;
            this.submitted = true;
          });
      },
  }
};
</script>

<style>
@import "~simplemde/dist/simplemde.min.css";

.disabled {
  @apply opacity-50 cursor-not-allowed
}
</style>
