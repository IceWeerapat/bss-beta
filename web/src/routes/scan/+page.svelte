<script lang="ts">
  import { onMount } from 'svelte';
  import QrScanner from 'qr-scanner';
  import { goto } from '$app/navigation';
  import {getObject} from '../../helper/local-storage' 

  let videoElement: HTMLVideoElement | null = null;
  let result: string | null = null;

  onMount(() => {
    let staff = getObject('staff');
    if (staff == null) {
      goto('/');
    }

    if (videoElement) {
      QrScanner.WORKER_PATH = '/node_modules/qr-scanner/qr-scanner-worker.min.js';
      const qrScanner = new QrScanner(videoElement, (scannedResult: string) => {
        result = scannedResult;
        console.log('Scanned QR code:', scannedResult);
      });
      qrScanner.start();
    }
  });

  const goBack = () => {
    goto('/menu');
  };
</script>

<style lang="postcss">
  :global(html) {
    @apply bg-gray-100;
  }
</style>

<div class="flex flex-col justify-center items-center min-h-screen bg-gradient-to-b from-white to-gray-100">
  <div class="text-base text-black mb-5">แสดงรหัสข้อมูลเพื่อทำการตรวจสอบ</div>
  <div class="relative w-11/12 max-w-md h-96 mb-5">
    <video bind:this={videoElement} class="w-full h-full rounded-lg object-cover"></video>
    <div class="absolute w-8 h-8 border-4 border-orange top-0 left-0 rounded-tl-lg" style="border-right: none; border-bottom: none;"></div>
    <div class="absolute w-8 h-8 border-4 border-orange top-0 right-0 rounded-tr-lg" style="border-left: none; border-bottom: none;"></div>
    <div class="absolute w-8 h-8 border-4 border-orange bottom-0 left-0 rounded-bl-lg" style="border-top: none; border-right: none;"></div>
    <div class="absolute w-8 h-8 border-4 border-orange bottom-0 right-0 rounded-br-lg" style="border-top: none; border-left: none;"></div>
  </div>
  <button class="bg-orange text-white rounded-bss py-3 px-6 text-base transition-colors duration-300 hover:bg-red-500" on:click={goBack}>กลับ</button>
  {#if result}
    <div class="mt-5 text-base text-black">result : {result}</div>
  {/if}
</div>
