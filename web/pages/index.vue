<template>
  <section class="container mx-auto py-4">
    <div>
      <h1 class="title">
        cryptpad-web
      </h1>
      <markdown-editor v-model="text" ref="markdownEditor"></markdown-editor>
      <input type="text" v-model="pass" id="key">
      <input type="submit" value="Save" @click.prevent="submitData" class="px-4 py-4 bg-teal-darker rounded text-white">
    </div>
  </section>
</template>

<script>
import markdownEditor from 'vue-simplemde/src/markdown-editor'
let CryptoJS = require("crypto-js");

  export default {
    components: {
      markdownEditor
    },
    data: function() {
      return {
        text: '',
        pass: 'secretpass'
      }
    },
    computed: {
      encrypted: function() {
        return CryptoJS.AES.encrypt(this.text, this.pass).toString();
      }
    },
    methods:{
      submitData() {
        this.$axios.post('/api/pad', this.encrypted)
        console.log("sent "+ this.encrypted)
      }
    }
  }
</script>

<style>
  @import '~simplemde/dist/simplemde.min.css';
</style>
