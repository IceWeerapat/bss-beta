<script lang="ts">
  import { onMount } from 'svelte';
  import axios from 'axios';
  import Navbar from '../Navbar.svelte';
  import DataTable from './DataTable.svelte';
  import { PUBLIC_API_URL } from '$env/static/public';
  import Dropdown from '../Dropdown.svelte';
  import Calendar from '../Calendar.svelte';

  let datas: any = [];

  const headers = [
    'ID', 'เลขทะเบียน', 'วันจดทะเบียน', 'จังหวัด', 'ประเภท', 
    'ลักษณะ', 'ยี่ห้อรถ', 'แบบ', 'รุ่นปี', 'สี', 
    'เลขตัวรถ', 'อยู่ที่', 'ยี่ห้อเครื่องยนต์', 'เลขเครื่องยนต์', 'อยู่ที่',
    'เชื้อเพลิง', 'เลขถังแก๊ส', 'จำนวนสูบ', 'ซีซี', 'แรงม้า', 
    'จำนวนเพลาและล้อ', 'น้ำหนักรถ(กก.)','น้ำหนักบรรทุก/น้ำหนักลงเพลา (กก.)', 'น้ำหนักรวม(กก.)','ที่นั่ง(คน)',
    'ลำดับที่เจ้าของรถล่าสุด', 'วันที่ครอบครองรถ', 'ผู้ถือกรรมสิทธิ์', 'เลขที่บัตร', 'วันเกิด',
    'สัญชาติ', 'ที่อยู่', 'โทร', 'สัญญาเช่าซื้อเลขที่', 'ลงวันที่', 
    'วันจ่ายภาษีล่าสุด', 'วันครบกำหนดเสียภาษี', 'ใบเสร็จรับเงินเลขที่คุม/เลขที่', 'ค่าภาษี บาท/สต.', 'เงินเพิ่ม บาท/สต.', 
    'ผู้บันทึก', 'นายทะเบียน',  'วันที่บันทึกล่าสุด', 'ภาพรถด้านหน้า', 'ภาพรถด้านข้าง',
    'ภาพรถด้านหลัง', 'วันที่สร้าง', 'สร้างโดย', 'วันที่แก้ไข', 'แก้ไขโดย'
  ];

  const onEdit = (d: any) => {
  };

  const onDelete = async (d: any) => {
  };

  const actions = [
    { label: 'Edit', callback: onEdit },
    { label: 'Delete', callback: onDelete }
  ];

  let staffId: any = null;
  let staffs: any = [];
  let startDate: any = null;
  let startDisplayDate = '';
  let endDate: any = null;
  let endDisplayDate = '';

  let currentPage = 1;
  let totalPages = 1;

  let dropdownComponent: any;

  const handleInputChange = ({ target }) => {
    const { name, value } = target;
    if (name === 'staff') {
      staffId = value;
    } else if (name === 'start_date') {
      startDate = value;
    } else if (name === 'end_date') {
      endDate = new Date(value);
      endDate.setHours(23, 59, 59, 999); // Set time to 23:59:59
    }
  };

  const fetchData = async (page = 1) => {
    try {
      const response = await axios.get(`${PUBLIC_API_URL}/data/list-criteria`, {
        params: {
          staff_id: staffId,
          start_date: startDate,
          end_date: endDate,
          order_by: 'created_at',
          sort: 'desc',
          page,
          limit: 20
        }
      });
      datas = response.data.data;
      const pagination = response.data.pagination;
      currentPage = pagination.current_page;
      totalPages = pagination.total_pages;
    } catch (error) {
      console.error('Failed to fetch data:', error);
    }
  };

  const handleFilter = () => {
    fetchData(1);
  };

  const handleClear = () => {
    staffId = null;
    startDate = null;
    startDisplayDate = '';
    endDate = null;
    endDisplayDate = '';
    currentPage = 1;

    // Clear the dropdown selection
    dropdownComponent.clearSelection();

    // Fetch data with cleared filters
    fetchData();
  };

  const handlePageChange = (page) => {
    fetchData(page);
  };

  onMount(async () => {
    const staffList = await axios.post(`${PUBLIC_API_URL}/staff/list-order-name`);
    staffs = staffList.data.data.map(s => ({
      label: `${s.first_name} ${s.last_name}`,
      value: s.id
    }));

    // Fetch initial data
    fetchData();
  });
</script>

<div class="flex h-screen">
  <Navbar />
  <div class="flex-1 overflow-auto">
    <div class="bg-white pt-14 ml-64 shadow-md rounded-lg flex flex-col justify-between">
      <div class="flex justify-between items-center m-2 mt-0">
        <h2 class="text-3xl font-bold">Data</h2>
        <!-- <button class="bg-orange-500 text-white px-4 py-2 rounded-md" on:click={onCreate}>Create</button> -->
      </div>
      <div class="mx-2 flex space-x-2">
        <Dropdown
          label="Staff"
          name="staff"
          bind:value={staffId}
          options={staffs}
          handleInputChange={handleInputChange}
          bind:this={dropdownComponent}
        />
        <Calendar
          label="Start Date"
          name="start_date"
          bind:selectedDate={startDate}
          bind:displayDate={startDisplayDate}
          handleInputChange={handleInputChange}
        />
        <Calendar
          label="End Date"
          name="end_date"
          bind:selectedDate={endDate}
          bind:displayDate={endDisplayDate}
          handleInputChange={handleInputChange}
        />
        <div class="flex flex-col justify-end pb-2">
          <button
            class="py-2 px-2 mb-2 bg-blue-400 text-white rounded-md w-20"
            on:click={handleFilter}
          >
            Filter
          </button>
        </div>
        <div class="flex flex-col justify-end pb-2">
          <button
            class="py-2 px-2 mb-2 bg-red-400 text-white rounded-md w-20"
            on:click={handleClear}
          >
            Clear
          </button>
        </div>
      </div>
      <DataTable {headers} data={datas} {actions} />
      <div class="flex justify-center mt-4 mb-4">
        {#if currentPage > 1}
          <button class="px-3 py-1 bg-gray-300 rounded-md mx-1" on:click={() => handlePageChange(currentPage - 1)}>Previous</button>
        {/if}
        <span class="px-3 py-1 mx-1">{currentPage} / {totalPages}</span>
        {#if currentPage < totalPages}
          <button class="px-3 py-1 bg-gray-300 rounded-md mx-1" on:click={() => handlePageChange(currentPage + 1)}>Next</button>
        {/if}
      </div>
    </div>
  </div>
</div>
