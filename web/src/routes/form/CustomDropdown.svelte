<script>
    import { onMount } from 'svelte';
    export let label = '';
    export let value = '';
    export let name = '';
    export let options = [];
    export let handleInputChange = () => {};
    export let isInvalid = false;
  
    let open = false;
  
    const toggleDropdown = () => {
      open = !open;
    };
  
    const selectOption = (option) => {
      value = option;
      handleInputChange({ target: { name, value: option } });
      open = false;
    };
  
    const handleClickOutside = (event) => {
      if (event.target.closest('.select-container') === null) {
        open = false;
      }
    };
  
    onMount(() => {
      document.addEventListener('click', handleClickOutside);
      return () => document.removeEventListener('click', handleClickOutside);
    });
  </script>
  
  <style>
    .selector-wrapper {
      @apply p-4 bg-box rounded-md mb-4 shadow-lg text-gray-font;
    }
  
    .selector-box {
      @apply w-full mt-2 p-2 border rounded-md bg-box text-black;
    }
  
    .selector-box.invalid {
      @apply border-red-500;
    }
  
    .selector-box.valid {
      @apply border-gray-100;
    }
  
    .select-container {
      @apply relative;
    }
  
    .select-icon {
      @apply absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none;
    }
  
    .select-options {
      @apply mt-1 absolute w-full rounded-md bg-white shadow-lg z-10 max-h-60 overflow-auto;
    }
  
    .select-option {
      @apply px-4 py-2 cursor-pointer hover:bg-gray-100 text-black;
    }
  
    .select-option.selected {
      @apply bg-blue-600 text-white;
    }
  </style>
  
  <div class="selector-wrapper">
    <label>{label}</label>
    <div class="select-container">
      <div
        class="selector-box {isInvalid ? 'invalid' : 'valid'}"
        on:click={toggleDropdown}
      >
        {value || `กรุณาเลือก${label}`}
        <div class="select-icon">
          <svg
            class="h-5 w-5 text-gray-400"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 20 20"
            fill="currentColor"
            aria-hidden="true"
          >
            <path
              fill-rule="evenodd"
              d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
              clip-rule="evenodd"
            />
          </svg>
        </div>
      </div>
      {#if open}
        <div class="select-options">
          {#each options as option}
            <div
              class="select-option {option === value ? 'selected' : ''}"
              on:click={() => selectOption(option)}
            >
              {option}
            </div>
          {/each}
        </div>
      {/if}
    </div>
  </div>
  