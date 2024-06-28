<script>
  import { onMount } from 'svelte';
  export let label = '';
  export let value = null;
  export let name = '';
  export let options = [];
  export let handleInputChange = () => {};
  export let isInvalid = false;

  let open = false;
  let selectedLabel = '';

  const toggleDropdown = () => {
    open = !open;
  };

  const selectOption = (option) => {
    value = option.value;
    selectedLabel = option.label;
    handleInputChange({ target: { name, value: option.value } });
    open = false;
  };

  const handleClickOutside = (event) => {
    if (event.target.closest('.select-container') === null) {
      open = false;
    }
  };

  // ฟังก์ชันเพื่อรีเซ็ตค่า dropdown
  export function clearSelection() {
    value = null;
    selectedLabel = '';
    handleInputChange({ target: { name, value: null } });
  }

  onMount(() => {
    document.addEventListener('click', handleClickOutside);
    return () => document.removeEventListener('click', handleClickOutside);
  });
</script>

<div class="p-2 bg-gray-100 rounded-md mb-4 shadow-lg text-black w-72">
  <label class="font-bold">{label}</label>
  <div class="relative select-container">
    <div
      class={`w-full mt-2 p-2 rounded-md bg-gray-100 text-black cursor-pointer ${isInvalid ? 'border-red-500' : 'border-gray-300'}`}
      on:click={toggleDropdown}
    >
      {selectedLabel || ''}
      <div class="absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
        <svg
          class="h-5 w-5 text-gray-400"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
          fill="currentColor"
          aria-hidden="true"
        >
          <path
            fill-rule="evenodd"
            d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 111.414 1.414l-4 4a1 1 01-1.414 0l-4-4a1 1 010-1.414z"
            clip-rule="evenodd"
          />
        </svg>
      </div>
    </div>
    {#if open}
      <div class="mt-1 absolute w-full rounded-md bg-gray-100 shadow-lg z-10 max-h-60 overflow-auto">
        <div class="px-4 py-2 cursor-pointer hover:bg-gray-100" on:click={() => selectOption({ label: 'None', value: null })}>
          None
        </div>
        {#each options as option}
          <div
            class={`px-4 py-2 cursor-pointer hover:bg-gray-100 ${option.value === value ? 'bg-blue-600 text-white' : ''}`}
            on:click={() => selectOption(option)}
          >
            {option.label}
          </div>
        {/each}
      </div>
    {/if}
  </div>
</div>
