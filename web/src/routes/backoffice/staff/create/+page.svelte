<script lang="ts">
    import { onMount } from 'svelte';
    import axios from 'axios';
    import Navbar from '../../Navbar.svelte';
    import { PUBLIC_API_URL } from '$env/static/public';
    import { goto } from '$app/navigation';
    import Modal from '../../Modal.svelte';  
    import { getObject } from '../../../../helper/local-storage';

    let firstName = '';
    let lastName = '';
    let pinCode = '';
    let showModal = false;
    let modalMessage = '';
    let modalTitle = '';

    let admin: any;

    onMount(async () => {
        admin = getObject('admin');
    });

    const handleSubmit = async (event: Event) => {
        event.preventDefault();

        if (!Number.isInteger(Number(pinCode))) {
            alert('Pin Code must be a valid integer.');
            return;
        }

        try {
            const response = await axios.post(`${PUBLIC_API_URL}/staff`, {
                first_name: firstName,
                last_name: lastName,
                pin_code: Number(pinCode),
                created_by: admin.first_name + ' ' + admin.last_name,
                updated_by: admin.first_name + ' ' + admin.last_name,
            });

            if (response.status === 200) {
                modalTitle = 'Success';
                modalMessage = 'Staff created successfully!';
                showModal = true;
            } else {
                modalTitle = 'Error';
                modalMessage = 'Failed to create staff. Please try again.';
                showModal = true;
            }
        } catch (error) {
            console.error('An error occurred during creation:', error);
            modalTitle = 'Error';
            modalMessage = 'An error occurred during creation. Please try again with another pincode.';
            showModal = true;
        }
    };

    const handleModalConfirm = () => {
        showModal = false;
        goto('/backoffice/staff');
    };
</script>

<div class="flex h-screen">
    <Navbar />
    <div class="flex-1 overflow-auto mx-5">
        <div class="bg-white pt-14 ml-64 shadow-md rounded-lg max-w-md mx-auto">
            <h2 class="text-3xl font-bold mb-4">Create Staff</h2>
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
                    <input type="text" id="pinCode" bind:value={pinCode} class="shadow appearance-none border rounded-lg w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required pattern="\d*" title="Please enter a valid integer">
                </div>
                <div class="flex items-center justify-between">
                    <button type="submit" class="bg-orange-500 hover:bg-orange-700 text-white font-bold py-2 px-4 rounded-lg focus:outline-none focus:shadow-outline">
                        Create
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
