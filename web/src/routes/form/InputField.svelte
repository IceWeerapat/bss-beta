<script>
    export let label = '';
    export let value = '';
    export let name = '';
    export let handleInputChange = () => {};
    export let isInvalid = false;
  </script>
  
  <style>
    .input-wrapper {
      @apply p-4 bg-box rounded-md mb-4 shadow-lg text-gray-font;
    }
  
    .input-box {
      @apply w-full mt-2 p-2 border rounded-md bg-box text-black;
    }
  
    .input-box.invalid {
      @apply border-red-500;
    }
  
    .input-box.valid {
      @apply border-gray-100;
    }
  </style>
  
  <div class="input-wrapper">
    <label>{label}</label>
    <input 
      type="text" 
      name={name} 
      bind:value={value} 
      on:input={handleInputChange} 
      class="input-box {isInvalid ? 'invalid' : 'valid'}"
    />
  </div>
  