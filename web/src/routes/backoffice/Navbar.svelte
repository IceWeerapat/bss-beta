<script lang="ts">
	import { onMount } from 'svelte';
	import { getObject, setObject } from '../../helper/local-storage';
	import { goto } from '$app/navigation';
  
	let activeMenu: string;
	let admin: any;
	let isMobile: boolean = false;
	let isOpenSide = false;
  
	onMount(() => {
	  admin = getObject('admin');
	  if (!admin) {
		goto('/backoffice');
	  }
	  activeMenu = window.location.pathname.substring(
		(window.location.pathname ?? '')?.lastIndexOf('/') + 1
	  );
	  checkIsMobile();
	  window.addEventListener('resize', checkIsMobile);
	});
  
	const checkIsMobile = () => {
	  isMobile = window.innerWidth < 1024;
	};
  
	const menus = [
	  {
		id: 1,
		title: 'Staff',
		href: '/backoffice/staff',
	  },
	  {
		id: 2,
		title: 'Data',
		href: '/backoffice/data',
	  },
	  {
		id: 3,
		title: 'Audit Log',
		href: '/backoffice/audit-log',
	  },
	//   {
	// 	id: 4,
	// 	title: 'Report',
	// 	href: '/backoffice/report',
	//   },
	];
  
	const logout = async () => {
	  setObject('admin', null);
	  goto('/backoffice');
	};
  </script>
  
  <nav class="fixed top-0 z-40 w-full bg-orange border-b border-gray-100">
	<div class="px-3 py-3 lg:px-5 lg:pl-3">
	  <div class="flex items-center justify-between">
		<div class="flex items-center justify-start">
		  <!-- if mobile size -->
		  {#if isMobile}
			<button
			  on:click={() => {
				isOpenSide = !isOpenSide;
			  }}
			  data-drawer-target="logo-sidebar"
			  data-drawer-toggle="logo-sidebar"
			  aria-controls="logo-sidebar"
			  type="button"
			  class="inline-flex items-center p-2 text-sm text-gray-500 rounded-lg lg:invisible">
			  <span class="sr-only">Open sidebar</span>
			  <svg
				class="w-6 h-6 fill-white"
				aria-hidden="true"
				viewBox="0 0 20 20"
				xmlns="http://www.w3.org/2000/svg">
				<path
				  clip-rule="evenodd"
				  fill-rule="evenodd"
				  d="M2 4.75A.75.75 0 012.75 4h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 4.75zm0 10.5a.75.75 0 01.75-.75h7.5a.75.75 0 010 1.5h-7.5a.75.75 0 01-.75-.75zM2 10a.75.75 0 01.75-.75h14.5a.75.75 0 010 1.5H2.75A.75.75 0 012 10z" />
			  </svg>
			</button>
		  {/if}
		  <!-- end if mobile size -->
		  <a href="/backoffice/home" class="flex ml-2 md:mr-24">
			<span class="self-center text-base font-semibold sm:text-lg whitespace-nowrap text-white">
			  BSS
			</span>
		  </a>
		</div>
  
		<div class="flex gap-3 items-center">
		  {#if admin}
			<span class="self-center text-base font-semibold sm:text-lg whitespace-nowrap text-white">
			  {admin.first_name} {admin.last_name}
			</span>
		  {/if}
		  <button
			class="text-gray-700 bg-white hover:bg-gray-200 focus:ring-4 focus:outline-none focus:ring-gray-200
					  rounded-lg border border-gray-100 text-sm font-medium px-3 py-1"
			on:click={logout}>Log Out</button>
		</div>
	  </div>
	</div>
  </nav>
  
  <aside
	id="logo-sidebar"
	class="fixed top-0 left-0 z-30 w-64 h-screen transition-transform -translate-x-full pt-20 bg-orange border-r
	  border-gray-100 lg:translate-x-0"
	class:translate-x-0={isOpenSide}
	aria-label="Sidebar">
	<div class="h-full px-3 pb-4 overflow-y-auto bg-orange">
	  <ul class="space-y-2 font-medium">
		{#each menus as menu}
		  <li>
			<a
			  href={menu.href}
			  class:bg-gray-100={activeMenu ===
				menu.href.substring((menu.href ?? '')?.lastIndexOf('/') + 1)}
			  class:bg-opacity-25={activeMenu ===
				menu.href.substring((menu.href ?? '')?.lastIndexOf('/') + 1)}
			  class="flex items-center p-2 text-white rounded-lg hover:bg-gray-100 hover:bg-opacity-25">
			  <span class="flex-1 text-base ml-3 whitespace-nowrap">{menu.title}</span>
			</a>
		  </li>
		{/each}
	  </ul>
	</div>
  </aside>
  