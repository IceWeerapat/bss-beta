<script lang="ts">
  import { onMount } from 'svelte';
  import axios from 'axios';
  import Navbar from '../../../Navbar.svelte';
  import { PUBLIC_API_URL } from '$env/static/public';
  import { page } from '$app/stores';
  import { goto } from '$app/navigation';
  import Modal from '../../../Modal.svelte';
  import { getObject } from '../../../../../helper/local-storage';

  let staffId: string;
  let firstName = '';
  let lastName = '';
  let pinCode: number;
  let showModal = false;
  let modalMessage = '';
  let modalTitle = '';

  let admin: any;

  $: staffId = $page.params.id;

  const fetchStaff = async () => {
    try {
      const response = await axios.get(`${PUBLIC_API_URL}/staff/${staffId}`);
      const staff = response.data.data;
      firstName = staff.first_name;
      lastName = staff.last_name;
      pinCode = staff.pin_code;
    } catch (error) {
      console.error('Failed to fetch staff data:', error);
    }
  };

  const handleSubmit = async (event: Event) => {
    event.preventDefault();
    try {
      const response = await axios.patch(`${PUBLIC_API_URL}/staff/${staffId}`, {
        first_name: firstName,
        last_name: lastName,
        pin_code: pinCode,
        updated_by: admin.first_name + ' ' + admin.last_name,   
      });
      if (response.status === 200) {
        modalTitle = 'Success';
        modalMessage = 'Staff member updated successfully!';
        showModal = true;
        setTimeout(() => {
          goto('/backoffice/staff');
        }, 2000);
      } else {
        modalTitle = 'Error';
        modalMessage = 'Failed to update staff member. Please try again with another pin code.';
        showModal = true;
      }
    } catch (error) {
      console.error('An error occurred during the update:', error);
      modalTitle = 'Error';
      modalMessage = 'Failed to update staff member. Please try again with another pin code.';
      showModal = true;
    }
  };

  const handleModalConfirm = () => {
    showModal = false;
    goto('/backoffice/staff');
  };

  onMount(() => {
    fetchStaff();
    admin = getObject('admin');
  });
</script>

<div class="flex h-screen">
  <Navbar />
  <div class="flex-1 overflow-auto mx-5">
    <div class="bg-white pt-14 ml-64 shadow-md rounded-lg max-w-md mx-auto">
      <h2 class="text-3xl font-bold mb-4">Edit Staff</h2>
      <form on:submit={handleSubmit}>
        <div class="mb-4">
          <label for="firstName" class="block text-gray-700 text-sm font-bold mb-2">First Name</label>
          <input type="text" id="firstName" bind:value={firstName} class="shadow appearance-none border rounded-lg w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required>
        </div>
        <div class="mb-4">
          <label for="lastName" class="block text-gray-700 text-sm font-bold mb-2">Last Name</label>
          <input type="text" id="lastName" bind:value={lastName} class="shadow appearance-none border rounded-lg w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required>
        </div>
        <div class="mb-4">
          <label for="pinCode" class="block text-gray-700 text-sm font-bold mb-2">Pin Code</label>
          <input type="number" id="pinCode" bind:value={pinCode} class="shadow appearance-none border rounded-lg w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required pattern="\d*" title="Please enter a valid integer">
        </div>
        <div class="flex items-center justify-between">
          <button type="submit" class="bg-orange-500 hover:bg-orange-700 text-white font-bold py-2 px-4 rounded-lg focus:outline-none focus:shadow-outline">
            Save
          </button>
        </div>
      </form>
    </div>
  </div>
</div>

<Modal title={modalTitle} message={modalMessage} isVisible={showModal} on:close={() => showModal = false} onConfirm={handleModalConfirm} />

<!-- <style>
  :global(body) {
    margin: 0;
  }
</style> -->
