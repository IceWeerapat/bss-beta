<script lang="ts">
    export let headers: string[] = [];
    export let data: any[] = [];
    export let actions: Array<{ label: string, callback: (row: any) => void }> = [];
  
    const handleAction = (callback: (row: any) => void, row: any) => {
      callback(row);
    };
  
    const formatDateTime = (dateString: string) => {
      const options: Intl.DateTimeFormatOptions = {
        year: 'numeric',
        month: '2-digit',
        day: '2-digit',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
      };
      const date = new Date(dateString);
      return date.toLocaleDateString('en-GB', options).replace(',', '');
    };
  </script>
  
  <div class="relative overflow-x-auto">
    <table class="w-full text-sm text-left text-gray-500">
      <thead class="text-xs text-white uppercase bg-gray-700">
        <tr>
          {#each headers as headerText}
            <th scope="col" class="px-6 py-3"> {headerText} </th>
          {/each}
          <!-- {#if actions.length > 0}
            <th scope="col" class="px-6 py-3 text-center"> Action </th>
          {/if} -->
        </tr>
      </thead>
      <tbody>
        {#if data.length > 0}
          {#each data as row}
            <tr class="bg-white border-b text-gray-700">
              <td class="px-6 py-4">{row.id}</td>
              <td class="px-6 py-4">{row.staff}</td>
              <td class="px-6 py-4">{row.license_number}</td>
              <td class="px-6 py-4">{row.action}</td>
              <td class="px-6 py-4">{row.effect_field}</td>
              <td class="px-6 py-4">{formatDateTime(row.created_at)}</td>
              <!-- {#if actions.length > 0}
                <td class="px-6 py-4 min-w-[140px]">
                  <div class="flex gap-3 justify-center">
                    {#each actions as action}
                      <button class="hover:underline" on:click={() => handleAction(action.callback, row)}>
                        {action.label}
                      </button>
                    {/each}
                  </div>
                </td>
              {/if} -->
            </tr>
          {/each}
        {:else}
          <tr class="bg-white border-b text-gray-500">
            <td colspan="{headers.length + (actions.length > 0 ? 1 : 0)}" class="bg-gray-200 px-6 py-3">No Data</td>
          </tr>
        {/if}
      </tbody>
    </table>
  </div>
  