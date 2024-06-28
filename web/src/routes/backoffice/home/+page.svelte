<script lang="ts">
  import { onMount } from 'svelte';
  import axios from 'axios';
  import Navbar from '../Navbar.svelte';
  import { PUBLIC_API_URL } from '$env/static/public';
  import Dropdown from '../Dropdown.svelte';

  let today = 0;
  let thisWeek = 0;
  let thisMonth = 0;
  let total = 0;
  let errorMessage = '';
  let staffId: any = null;
  let staffs: any = [];

  const fetchInitialSummary = async () => {
    try {
      const summary = await axios.get(`${PUBLIC_API_URL}/data/summary`);
      const data = summary.data.data;

      today = data.today;
      thisWeek = data.this_week;
      thisMonth = data.this_month;
      total = data.total;
    } catch (error) {
      console.error('Failed to fetch summary data:', error);
      errorMessage = 'Failed to fetch summary data.';
    }
  };

  onMount(async () => {
    fetchInitialSummary();

    const staffList = await axios.post(`${PUBLIC_API_URL}/staff/list-order-name`);
    staffs = staffList.data.data.map(s => ({
      label: `${s.first_name} ${s.last_name}`,
      value: s.id
    }));
  });

  const fetchStaffSummary = async (value) => {
    if (value === null) {
      fetchInitialSummary();
      return;
    }
    try {
      const response = await axios.get(`${PUBLIC_API_URL}/data/summary/${value}`);
      const data = response.data.data;
      today = data.today;
      thisWeek = data.this_week;
      thisMonth = data.this_month;
      total = data.total;
    } catch (error) {
      console.error('Failed to fetch staff summary data:', error);
      errorMessage = 'Failed to fetch staff summary data.';
    }
  };

  const handleInputChange = ({ target }) => {
    staffId = target.value;
    fetchStaffSummary(staffId);
  };
</script>

<div class="flex h-screen">
  <Navbar />
  <div class="flex-1 overflow-auto">
    <div class="bg-white pt-14 ml-64 shadow-md rounded-lg max-w-2xl mx-auto">
      <div class="flex justify-between items-center">
        <!-- <h2 class="text-3xl font-bold">Report</h2> -->
      </div>
      {#if errorMessage}
        <div class="text-red-500">{errorMessage}</div>
      {:else}
        <div class="pl-1">
          <div class="mb-4">
            <h2 class="text-3xl font-bold">Data Registration Summary</h2>
            <Dropdown
              label="Staff"
              name="staff"
              bind:value={staffId}
              options={staffs}
              handleInputChange={handleInputChange}
            />
            <div class="bg-gray-100 p-4 rounded-lg mt-4">
              <div class="grid grid-cols-3 gap-x-4 gap-y-2 items-center">
                <div class="font-bold">Today:</div>
                <div class="text-right">{today}</div>
                <div class="font-bold text-center">Record</div>

                <div class="font-bold">This Week:</div>
                <div class="text-right">{thisWeek}</div>
                <div class="font-bold text-center">Record</div>

                <div class="font-bold">This Month:</div>
                <div class="text-right">{thisMonth}</div>
                <div class="font-bold text-center">Record</div>

                <div class="font-bold">Total:</div>
                <div class="text-right">{total}</div>
                <div class="font-bold text-center">Record</div>
              </div>
            </div>
          </div>
        </div>
      {/if}
    </div>
  </div>
</div>
