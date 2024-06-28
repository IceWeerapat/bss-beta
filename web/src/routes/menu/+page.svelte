<script lang="ts">
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import {getObject} from '../../helper/local-storage' 

  const navigateToForm = () => {
    goto('/form');
  };

  const navigateToScan = () => {
    goto('/scan');
  };

  onMount(() => {
    let staff = getObject('staff');
    if (staff == null) {
      goto('/');
    }
  });
</script>

<style lang="postcss">
  :global(html) {
    @apply bg-gray-100;
  }
</style>

<div class="flex flex-col justify-center items-center h-screen bg-gradient-to-b from-white to-gray-100">
  <button class="bg-orange text-white rounded-bss px-6 py-3 text-xl transition duration-300 mt-5 hover:bg-orange-500 w-60 h-24" on:click={navigateToForm}>สร้างรหัสข้อมูล</button>
  <button class="bg-orange text-white rounded-bss px-6 py-3 text-xl transition duration-300 mt-5 hover:bg-orange-500 w-60 h-24" on:click={navigateToScan}>ตรวจสอบรหัสข้อมูล</button>
</div>
