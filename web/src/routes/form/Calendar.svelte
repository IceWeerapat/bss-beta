<script>
  import { onMount } from 'svelte';
  export let selectedDate = '';
  export let displayDate = '';
  export let label = '';
  export let name = '';
  export let handleInputChange = () => {};
  export let isInvalid = false;

  let showCalendar = false;
  let currentMonth = new Date().getMonth();
  let currentYear = new Date().getFullYear();
  let selectingYear = false;

  const daysInWeek = ['อา.', 'จ.', 'อ.', 'พ.', 'พฤ.', 'ศ.', 'ส.'];

  let dates = [];
  let years = [];

  const toggleCalendar = (event) => {
    event.stopPropagation();
    showCalendar = !showCalendar;
  };

  const closeCalendar = (event) => {
    event.stopPropagation();
    showCalendar = false;
  };

  const handleDateClick = (event, date) => {
    event.stopPropagation();
    const selected = new Date(currentYear, currentMonth, date);
    selectedDate = selected.toISOString();
    displayDate = selected.toLocaleDateString('th-TH', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric'
    });
    handleInputChange({ target: { name, value: selectedDate } });
    closeCalendar(event);
  };

  const nextMonth = (event) => {
    event.stopPropagation();
    event.preventDefault(); // Prevent default behavior
    if (currentMonth === 11) {
      currentMonth = 0;
      currentYear++;
    } else {
      currentMonth++;
    }
    generateCalendar();
  };

  const prevMonth = (event) => {
    event.stopPropagation();
    event.preventDefault(); // Prevent default behavior
    if (currentMonth === 0) {
      currentMonth = 11;
      currentYear--;
    } else {
      currentMonth--;
    }
    generateCalendar();
  };

  const nextYear = (event) => {
    event.stopPropagation();
    event.preventDefault(); // Prevent default behavior
    currentYear += 10;
    generateYears();
  };

  const prevYear = (event) => {
    event.stopPropagation();
    event.preventDefault(); // Prevent default behavior
    currentYear -= 10;
    generateYears();
  };

  const handleYearClick = (event, year) => {
    event.stopPropagation();
    currentYear = year - 543; // Convert B.E. to A.D.
    selectingYear = false;
    generateCalendar();
  };

  const generateCalendar = () => {
    dates = [];
    let firstDay = new Date(currentYear, currentMonth, 1).getDay();
    let daysInMonth = new Date(currentYear, currentMonth + 1, 0).getDate();

    for (let i = 0; i < firstDay; i++) {
      dates.push('');
    }
    for (let i = 1; i <= daysInMonth; i++) {
      dates.push(i);
    }
  };

  const generateYears = () => {
    years = [];
    let startYear = currentYear - 5;
    let endYear = currentYear + 4;
    for (let i = startYear; i <= endYear; i++) {
      years.push(i + 543); // Convert A.D. to B.E.
    }
  };

  onMount(() => {
    generateCalendar();
    generateYears();
  });
</script>

<div class="p-4 bg-gray-100 rounded-md mb-4 shadow-lg text-gray-400">
  <label>{label}</label>
  <div class="relative">
    <input
      type="text"
      readonly
      value={displayDate || ''}
      on:click={toggleCalendar}
      class="w-full mt-2 p-2 rounded-md bg-gray-100 text-black cursor-pointer border {isInvalid ? 'border-red-500' : 'border-gray-100'} placeholder-black"
      placeholder="กรุณาเลือกวันที่"
    />
    <div class="absolute inset-y-0 right-0 flex items-center pr-3 pointer-events-none">
      <svg class="h-5 w-5 text-gray-400" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor" aria-hidden="true">
        <path d="M6 2a1 1 0 011 1v1h6V3a1 1 0 112 0v1h1a2 2 0 012 2v10a2 2 0 01-2 2H5a2 2 0 01-2-2V6a2 2 0 012-2h1V3a1 1 0 011-1zm8 4H6a1 1 0 00-1 1v8a1 1 0 001 1h8a1 1 0 001-1V7a1 1 0 00-1-1zm-1 2a1 1 0 010 2H7a1 1 0 110-2h6z" />
      </svg>
    </div>
    {#if showCalendar}
      <div class="absolute mt-1 w-full bg-white shadow-lg rounded-md z-10" on:click|stopPropagation>
        {#if selectingYear}
          <div class="flex justify-between items-center p-2 bg-gray-100 border-b border-gray-300">
            <button on:click={prevYear} class="text-blue-500">ก่อนหน้า</button>
            <div>{currentYear - 5 + 543} - {currentYear + 4 + 543}</div>
            <button on:click={nextYear} class="text-blue-500">ถัดไป</button>
          </div>
          <div class="grid grid-cols-3 gap-1 p-2">
            {#each years as year}
              <div
                class="text-center cursor-pointer p-2 rounded-md hover:bg-blue-100"
                on:click={(event) => handleYearClick(event, year)}
              >
                {year}
              </div>
            {/each}
          </div>
        {:else}
          <div class="flex justify-between items-center p-2 bg-gray-100 border-b border-gray-300">
            <button on:click={prevMonth} class="text-blue-500">ก่อนหน้า</button>
            <div on:click={(event) => {event.stopPropagation(); selectingYear = true;}} class="cursor-pointer">
              {new Date(currentYear, currentMonth).toLocaleString('th-TH', { month: 'long' })} {currentYear + 543}
            </div>
            <button on:click={nextMonth} class="text-blue-500">ถัดไป</button>
          </div>
          <div class="grid grid-cols-7 gap-1 p-2">
            {#each daysInWeek as day}
              <div class="text-center font-bold">{day}</div>
            {/each}
            {#each dates as date}
              <div
                class="text-center cursor-pointer p-2 rounded-md hover:bg-blue-100 {date ? '' : 'invisible'}"
                on:click={(event) => date && handleDateClick(event, date)}
              >
                {date}
              </div>
            {/each}
          </div>
        {/if}
      </div>
    {/if}
  </div>
</div>

<style>
input::placeholder {
    color: black;
}
</style>
