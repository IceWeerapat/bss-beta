<script lang="ts">
  import { onMount } from 'svelte';
  import axios from 'axios';
  import Navbar from '../Navbar.svelte';
  import AuditLogTable from './AudiLogTable.svelte';
  import { PUBLIC_API_URL } from '$env/static/public';

  let auditLogs:any = [];
  const headers = [
    'ID', 'Staff', 'Data (Lincese No.)', 'Action', 'Effect Field', 'Created At'
  ];

  const onEdit = (a:any) => {
  };

  const onDelete = async (a:any) => {
  };

  const actions = [
    { label: 'Edit', callback: onEdit },
    { label: 'Delete', callback: onDelete }
  ];

  onMount(async () => {
    try {
      const response = await axios.post(`${PUBLIC_API_URL}/audit-log/list`, {
      });

      auditLogs = response.data.data;
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
        <h2 class="text-3xl font-bold">Audit Log</h2>
      </div>
      <AuditLogTable {headers} data={auditLogs} {actions} />
    </div>
  </div>
</div>

<style>
  :global(body) {
    margin: 0;
  }
</style>
