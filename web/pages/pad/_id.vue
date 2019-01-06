<template>
<section class="container mx-auto py-4">
    <input type="text" v-model="key" id="">
    <input type="submit" value="Decrypt" @click.prevent="decrypt">
    <div v-html="md" class="border-1 border-grey-dark rounded-sm px-4 py-6"></div>
</section>
</template>

<script>
let CryptoJS = require("crypto-js");
import marked from 'marked';

export default {
  asyncData ({ params, $axios }) {
    return $axios.get(`http://localhost:3000/api/pad/${params.id}`)
    .then((res) => {
      return { encrypted: res.data.data }
    })
  },
  data: function() {
      return {
          key: '',
          decrypted: ''
      }
  },
  computed: {
      md: function() {
          return marked(this.decrypted)
      }
  },
  methods: {
      decrypt() {
        var bytes  = CryptoJS.AES.decrypt(this.encrypted, this.key);
        this.decrypted = bytes.toString(CryptoJS.enc.Utf8);
      }
  }
}
</script>

