<script lang="ts">
  import { onMount } from 'svelte';
  import axios from 'axios';
  import Navbar from '../Navbar.svelte';
  import StaffTable from './StaffTable.svelte';
  import { PUBLIC_API_URL } from '$env/static/public';
  import { goto } from '$app/navigation';

  let staffs: any = [];
  let showModal = false;
  let modalMessage = '';
  let modalAction:any = null;
  let staffToDelete: any = null;

  const headers = ['ID', 'First Name', 'Last Name', 'Pin Code', 'Created At', 'Created By', 'Updated At', 'Updated By'];

  const onEdit = (s:any) => {
    goto(`/backoffice/staff/${s.id}/edit`);
  };

  const onDelete = (s:any) => {
    staffToDelete = s;
    modalMessage = `Are you sure you want to delete staff member ${s.first_name} ${s.last_name}?`;
    showModal = true;
    modalAction = confirmDelete;
  };

  const confirmDelete = async () => {
    if (staffToDelete) {
      try {
        const response = await axios.delete(`${PUBLIC_API_URL}/staff/${staffToDelete.id}`);
        if (response.status === 200) {
          staffs = staffs.filter((s: any) => s.id !== staffToDelete.id);
          modalMessage = 'Staff member deleted successfully.';
        } else {
          modalMessage = 'Failed to delete staff member. Please try again.';
        }
      } catch (error) {
        console.error('An error occurred:', error);
        modalMessage = 'An error occurred during deletion. Please try again.';
      } finally {
        staffToDelete = null;
        modalAction = () => showModal = false;
      }
    }
  };

  const actions = [
    { label: 'Edit', callback: onEdit },
    { label: 'Delete', callback: onDelete }
  ];

  const onCreate = () => {
    goto(`/backoffice/staff/create`);
  };

  onMount(async () => {
    try {
      const response = await axios.post(`${PUBLIC_API_URL}/staff/list`);
      staffs = response.data.data;
    } catch (error) {
      console.error('Failed to fetch users:', error);
    }
  });
</script>

<div class="flex h-screen">
  <Navbar />
  <div class="flex-1 overflow-auto">
    <div class="bg-white pt-14 ml-64 shadow-md rounded-lg">
      <div class="flex justify-between items-center m-2 mt-0">
        <h2 class="text-3xl font-bold">Staff</h2>
        <button class="bg-orange-500 text-white px-4 py-2 rounded-md" on:click={onCreate}>Create</button>
      </div>
      <StaffTable {headers} data={staffs} {actions} />
    </div>
  </div>
</div>

{#if showModal}
  <div class="fixed inset-0 bg-gray-600 bg-opacity-75 flex items-center justify-center">
    <div class="bg-white p-8 rounded-lg shadow-lg">
      <p class="mb-4">{modalMessage}</p>
      <div class="flex justify-between">
        {#if staffToDelete}
          <button class="bg-red-500 text-white px-4 py-2 rounded-md" on:click={modalAction}>
            Delete
          </button>
          <button class="bg-gray-500 text-white px-4 py-2 rounded-md" on:click={() => { showModal = false; staffToDelete = null; }}>
            Cancel
          </button>
        {:else}
          <button class="bg-orange-500 text-white px-4 py-2 rounded-md" on:click={() => { showModal = false; }}>
            OK
          </button>
        {/if}
      </div>
    </div>
  </div>
{/if}

<!-- <style>
  :global(body) {
    margin: 0;
  }
</style> -->
