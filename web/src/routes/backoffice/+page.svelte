<script lang="ts">
	import { onMount } from 'svelte';
	import axios from 'axios';
	import { PUBLIC_API_URL } from '$env/static/public';
	import { goto } from '$app/navigation';
	import {getObject, setObject} from '../../helper/local-storage'

	let username = '';
	let password = '';

	const handleSubmit = async (event: Event) => {
		event.preventDefault();

		try {
			const response = await axios.post(`${PUBLIC_API_URL}/admin/login`, {
				username,
				password
			});

			if (response) {
				setObject('admin', response.data.data);
				goto('/backoffice/home');
			} else {
				alert('Login failed. Please check your username and password.');
			}
		} catch (error) {
			console.error('An error occurred:', error);
			alert('An error occurred during login. Please try again.');
		}
	};

	onMount(async () => {
	let admin = getObject('admin');
	if(admin != null) {
		goto('/backoffice/home')
	}
  });
</script>

<main class="min-h-screen flex items-center justify-center bg-orange-500">
	<div class="bg-white p-8 rounded-lg shadow-lg w-full max-w-md">
		<h1 class="text-2xl font-bold mb-6 text-orange-500">Login</h1>
		<form on:submit={handleSubmit}>
			<div class="mb-4">
				<label for="username" class="block text-gray-700 text-sm font-bold mb-2">Username</label>
				<input type="text" id="username" bind:value={username} class="shadow appearance-none border rounded-md w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" required>
			</div>
			<div class="mb-6">
				<label for="password" class="block text-gray-700 text-sm font-bold mb-2">Password</label>
				<input type="password" id="password" bind:value={password} class="shadow appearance-none border rounded-md w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" required>
			</div>
			<div class="flex items-center justify-between">
				<button type="submit" class="bg-orange-500 hover:bg-orange-700 text-white font-bold py-2 px-4 rounded-md focus:outline-none focus:shadow-outline">
					Login
				</button>
			</div>
		</form>
	</div>
</main>

<!-- <style>
	:global(body) {
		margin: 0;
	}
</style> -->
