<script lang="ts">
  import { onMount } from 'svelte';
  import { goto } from '$app/navigation';
  import axios from 'axios';
  import {setObject} from '../../helper/local-storage'
  import { PUBLIC_API_URL } from '$env/static/public';

  let pin = '';
  const maxPinLength = 6;
  let errorMessage = '';

  const handleButtonClick = (value: string) => {
    if (value === 'delete') {
      pin = pin.slice(0, -1);
    } else {
      if (pin.length < maxPinLength) {
        pin += value;
      }
    }

    if (pin.length === maxPinLength) {
      checkPincode();
    }
  };

  const checkPincode = async () => {
    try {
      const response = await axios.get(`${PUBLIC_API_URL}/staff/pin-code/${pin}`);
      if (response) {
        setObject('staff',response.data.data);
        goto('/menu')
      } 
		} catch (error) {
      errorMessage = 'pincode incorrect';
        pin = '';
        setTimeout(() => {
          errorMessage = '';
        }, 5000); 
		}
  };

  onMount(() => {
  });
</script>

<style lang="postcss">
  :global(html) {
    @apply bg-gray-100;
  }
</style>

<div class="flex flex-col justify-center items-center h-screen bg-gradient-to-b from-white to-gray-100">
  <div class="text-3xl text-orange mb-10">ใส่รหัสผ่าน</div>
  <div class="flex justify-center mb-4">
    {#each { length: maxPinLength } as _, i}
      <div class="w-5 h-5 mx-2 rounded-full border-2 border-orange {pin.length > i ? 'bg-orange' : ''}"></div>
    {/each}
  </div>
  <div class="grid grid-cols-3 gap-6 mb-5">
    {#each Array.from({ length: 9 }, (_, i) => i + 1) as number}
      <button class="w-20 h-20 rounded-full bg-white shadow-3xl text-3xl text-orange cursor-pointer transition duration-300 active:bg-gray-200" on:click={() => handleButtonClick(number.toString())}>{number}</button>
    {/each}
    <button class="w-20 h-20 rounded-full  bg-white shadow-3xl text-3xl text-orange cursor-pointer transition duration-300 active:bg-gray-200" on:click={() => handleButtonClick('delete')}>←</button>
    <button class="w-20 h-20 rounded-full  bg-white shadow-3xl text-3xl text-orange cursor-pointer transition duration-300 active:bg-gray-200" on:click={() => handleButtonClick('0')}>0</button>
  </div>
  {#if errorMessage}
    <div class="text-red-500 text-lg mt-2">{errorMessage}</div>
  {/if}
</div>
