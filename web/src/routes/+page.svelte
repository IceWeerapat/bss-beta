<script lang="ts">
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import {setObject} from '../helper/local-storage' 

  const navigateToPinCode = () => {
    goto('/pin-code');
  };

  onMount(() => {
    setObject('staff',null)
  });
</script>

<style lang="postcss">
  :global(html) {
    @apply bg-gray-100;
  }
</style>

<div class="flex flex-col justify-center items-center h-screen bg-gradient-to-b from-white to-gray-100">
    <div class="text-2xl text-orange mb-10">ยินดีต้อนรับเข้าสู่ระบบ</div>
    <button class="bg-orange text-white rounded-bss px-6 py-4 text-lg cursor-pointer transition duration-300 mt-5 hover:bg-orange-500" on:click={navigateToPinCode}>เข้าสู่ระบบ</button>
</div>
