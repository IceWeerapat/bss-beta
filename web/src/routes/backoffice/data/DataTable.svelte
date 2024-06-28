<script lang="ts">
  export let headers: string[] = [];
  export let data: any[] = [];
  export let actions: Array<{ label: string, callback: (row: any) => void }> = [];

  let isModalOpen = false;
  let modalImageSrc = '';

  const handleAction = (callback: (row: any) => void, row: any) => {
    callback(row);
  };

  const formatDate = (dateString: string) => {
    const options: Intl.DateTimeFormatOptions = {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
    };
    const date = new Date(dateString);
    return date.toLocaleDateString('en-GB', options).replace(',', '');
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

  const openModal = (src: any) => {
    modalImageSrc = src;
    isModalOpen = true;
  };

  const closeModal = () => {
    isModalOpen = false;
    modalImageSrc = '';
  };
</script>

<div class="relative overflow-x-auto">
  <table class="w-full text-sm text-left text-gray-500">
    <thead class="text-xs text-white uppercase bg-gray-700">
      <tr>
        {#each headers as headerText}
          <th scope="col" class="px-6 py-3 whitespace-nowrap text-center"> {headerText} </th>
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
            <td class="px-2 py-4 text-center align-middle">{row.id}</td>
            <td class="px-2 py-4 text-center align-middle">{row.license_number}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDate(row.date_registration)}</td>
            <td class="px-2 py-4 text-center align-middle">{row.province}</td>
            <td class="px-2 py-4 text-center align-middle">{row.type}</td>
            <td class="px-2 py-4 text-center align-middle">{row.characteristic}</td>
            <td class="px-2 py-4 text-center align-middle">{row.brand}</td>
            <td class="px-2 py-4 text-center align-middle">{row.model}</td>
            <td class="px-2 py-4 text-center align-middle">{row.year}</td>
            <td class="px-2 py-4 text-center align-middle">{row.color}</td>
            <td class="px-2 py-4 text-center align-middle">{row.vehicle_number}</td>
            <td class="px-2 py-4 text-center align-middle">{row.location_vehicle}</td>
            <td class="px-2 py-4 text-center align-middle">{row.engine_brand}</td>
            <td class="px-2 py-4 text-center align-middle">{row.engine_number}</td>
            <td class="px-2 py-4 text-center align-middle">{row.location_engine}</td>
            <td class="px-2 py-4 text-center align-middle">{row.fuel}</td>
            <td class="px-2 py-4 text-center align-middle">{row.gas_tank_number}</td>
            <td class="px-2 py-4 text-center align-middle">{row.cylinder_count}</td>
            <td class="px-2 py-4 text-center align-middle">{row.cc}</td>
            <td class="px-2 py-4 text-center align-middle">{row.horsepower}</td>
            <td class="px-2 py-4 text-center align-middle">{row.axle_and_wheel_count}</td>
            <td class="px-2 py-4 text-center align-middle">{row.vehicle_weight}</td>
            <td class="px-2 py-4 text-center align-middle">{row.load_weight}</td>
            <td class="px-2 py-4 text-center align-middle">{row.total_weight}</td>
            <td class="px-2 py-4 text-center align-middle">{row.seat_count}</td>
            <td class="px-2 py-4 text-center align-middle">{row.latest_owner_number}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDate(row.ownership_date)}</td>
            <td class="px-2 py-4 text-center align-middle">{row.owner}</td>
            <td class="px-2 py-4 text-center align-middle">{row.id_card_number}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDate(row.birth_date)}</td>
            <td class="px-2 py-4 text-center align-middle">{row.nationality}</td>
            <td class="px-2 py-4 text-center align-middle">{row.address}</td>
            <td class="px-2 py-4 text-center align-middle">{row.phone}</td>
            <td class="px-2 py-4 text-center align-middle">{row.lease_contract_number}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDate(row.contract_date)}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDate(row.last_tax_paid_date)}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDate(row.tax_due_date)}</td>
            <td class="px-2 py-4 text-center align-middle">{row.receipt_number}</td>
            <td class="px-2 py-4 text-center align-middle">{row.tax_amount}</td>
            <td class="px-2 py-4 text-center align-middle">{row.additional_amount}</td>
            <td class="px-2 py-4 text-center align-middle">{row.recorder}</td>
            <td class="px-2 py-4 text-center align-middle">{row.registrar}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDate(row.last_update_date)}</td>
            <td class="px-0 py-4 text-center align-middle">
              <div class="flex justify-center items-center">
                <img src={`${row.picture_front}`} alt="Picture Front" class="w-16 h-16 object-cover cursor-pointer border-2" on:click={() => openModal(row.picture_front)}>
              </div>
            </td>
            <td class="px-0 py-4 text-center align-middle">
              <div class="flex justify-center items-center">
                <img src={`${row.picture_side}`} alt="Picture Side" class="w-16 h-16 object-cover cursor-pointer border-2" on:click={() => openModal(row.picture_side)}>
              </div>
            </td>
            <td class="px-0 py-4 text-center align-middle">
              <div class="flex justify-center items-center">
                <img src={`${row.picture_back}`} alt="Picture Back" class="w-16 h-16 object-cover cursor-pointer border-2" on:click={() => openModal(row.picture_back)}>
              </div>
            </td>
            <td class="px-2 py-4 text-center align-middle">{formatDateTime(row.created_at)}</td>
            <td class="px-2 py-4 text-center align-middle">{row.created_by}</td>
            <td class="px-2 py-4 text-center align-middle">{formatDateTime(row.updated_at)}</td>
            <td class="px-2 py-4 text-center align-middle">{row.updated_by}</td>
          </tr>
        {/each}
      {:else}
        <tr class="bg-white border-b text-gray-500">
          <td colspan="{headers.length + (actions.length > 0 ? 1 : 0)}" class="bg-gray-200 px-6 py-3 text-center align-middle">No Data</td>
        </tr>
      {/if}
    </tbody>
  </table>
</div>

{#if isModalOpen}
  <div class="fixed inset-0 bg-black bg-opacity-75 flex items-center justify-center z-50">
    <div class="relative bg-white p-4 rounded-lg">
      <button class="absolute top-0 right-0 mt-5 mr-2 text-red-500" on:click={closeModal}>X</button>
      <img src={modalImageSrc} alt="Enlarged Image" class="max-w-full max-h-screen" />
    </div>
  </div>
{/if}
