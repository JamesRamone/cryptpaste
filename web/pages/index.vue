<template>
  <section class="container mx-auto py-4">
    <h1 class="title mb-4">cryptpad-web</h1>
    <div>
      <markdown-editor v-model="text" ref="markdownEditor" :configs="mdeconf"></markdown-editor>
      <div class="flex mt-4 align-middle items-center justify-between">
        <div>Your share link:
          <a v-bind:href="link">{{ link }}</a>
        </div>
        <div>
          <input
            type="text"
            v-model="pass"
            placeholder="Enter encryption key..."
            id="key"
            class="py-4 px-2 bg-white text-grey-darkest rounded-md shadow-md focus:shadow-outline"
          >
          <input
            type="submit"
            value="Encrypt & store"
            @click.prevent="submitData"
            class="px-4 py-4 bg-teal-darker rounded text-white"
            v-bind:class="{ disabled: !enableSubmit }"
            :disabled="!enableSubmit"
          >
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import markdownEditor from "vue-simplemde/src/markdown-editor";
let CryptoJS = require("crypto-js");

export default {
  components: {
    markdownEditor
  },
  data: function() {
    return {
      text: "",
      pass: "",
      uuid: "",
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
      return this.uuid !== "" ? "http://localhost:3333/pad/" + this.uuid : "";
    },
    enableSubmit: function() {
      return this.text.length && this.pass.length;
    }
  },
  methods: {
    submitData() {
      
        this.$axios
          .post("http://localhost:3000/api/pad", { data: this.encrypted })
          .then(response => {
            console.log(response);
            this.uuid = response.data;
          });
      }
    
  }
};
</script>

<style>
@import "~simplemde/dist/simplemde.min.css";

.disabled {
  @apply opacity-50 cursor-not-allowed
}
</style>
