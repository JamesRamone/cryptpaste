<template>
<div>
  <div class="bg-grey-lighter min-h-screen">
  <Navbar/>
    <nuxt/>
  </div>
</div>
</template>

<script>
import Navbar from "~/components/Navbar.vue";
export default {
    components: {
        Navbar,
    }
}
</script>

