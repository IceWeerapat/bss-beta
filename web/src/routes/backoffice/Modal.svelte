<script>
    export let title = '';
    export let message = '';
    export let isVisible = false;
    export let onConfirm = () => {};
  
    const handleConfirm = () => {
      onConfirm();
    };
  </script>
  
  {#if isVisible}
    <div class="fixed inset-0 bg-gray-600 bg-opacity-75 flex items-center justify-center">
      <div class="bg-white p-8 rounded-lg shadow-lg">
        <h2 class="text-2xl font-bold mb-4">{title}</h2>
        <p class="mb-4">{message}</p>
        <button class="bg-orange-500 text-white px-4 py-2 rounded-md" on:click={handleConfirm}>
          Back
        </button>
      </div>
    </div>
  {/if}
  