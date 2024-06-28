<script>
  import { goto } from '$app/navigation';
  
  export let message = '';
  export let isVisible = false;
  export let onConfirm = () => {};
  
  const handleConfirm = () => {
    onConfirm();
  };
</script>

{#if isVisible}
  <div class="fixed inset-0 bg-gray-600 bg-opacity-75 flex items-center justify-center">
    <div class="bg-white p-8 rounded-lg shadow-lg flex flex-col items-center">
      <p class="mb-4 text-2xl text-center">{message}</p>
      <div class="flex justify-center">
        <button class="bg-orange-500 text-white px-4 py-2 rounded-md" on:click={handleConfirm}>
          ตกลง
        </button>
      </div>
    </div>
  </div>
{/if}
