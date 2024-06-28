<script>
  import InputField from './InputField.svelte';
  import CustomDropdown from './CustomDropdown.svelte';
  import UploadPicture from './UploadPicture.svelte';
  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import { getObject } from '../../helper/local-storage';
  import Calendar from './Calendar.svelte';
  import Modal from './Modal.svelte';
  import axios from 'axios'; 
  import { PUBLIC_API_URL } from '$env/static/public';

  let staff;
  let existData;
  let isExistData = false;

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleDateString('th-TH', { day: '2-digit', month: '2-digit', year: 'numeric' });
  };

  let inputs = {
    date_registration: '',
    license_number: '',
    province: '',
    type: '',
    characteristic: '',
    brand: '',
    model: '',
    year: '',
    color: '',
    vehicle_number: '',
    location_vehicle: '',
    engine_brand: '',
    engine_number: '',
    location_engine: '',
    fuel: '',
    gas_tank_number: '',
    cylinder_count: '',
    cc: '',
    horsepower: '',
    axle_and_wheel_count: '',
    vehicle_weight: '',
    load_weight: '',
    total_weight: '',
    seat_count: '',
    latest_owner_number: '',
    ownership_date: '',
    owner: '',
    id_card_number: '',
    birth_date: '',
    nationality: '',
    address: '',
    phone: '',
    lease_contract_number: '',
    contract_date: '',
    last_tax_paid_date: '',
    tax_due_date: '',
    receipt_number: '',
    tax_amount: '',
    additional_amount: '',
    recorder: '',
    registrar: '',
    last_update_date: new Date().toISOString().split('.')[0] + '.000Z',
    display_last_update_date: formatDate(new Date()),
    picture_front: '',
    picture_back: '',
    picture_side: '',
    created_by: '',
    updated_by: '',
    staff_id: '',
  };

  let inputValidity = {
    date_registration: true,
    license_number: true,
    province: true,
    type: true,
    characteristic: true,
    brand: true,
    model: true,
    year: true,
    color: true,
    vehicle_number: true,
    location_vehicle: true,
    engine_brand: true,
    engine_number: true,
    location_engine: true,
    fuel: true,
    gas_tank_number: true,
    cylinder_count: true,
    cc: true,
    horsepower: true,
    axle_and_wheel_count: true,
    vehicle_weight: true,
    load_weight: true,
    total_weight: true,
    seat_count: true,
    latest_owner_number: true,
    ownership_date: true,
    owner: true,
    id_card_number: true,
    birth_date: true,
    nationality: true,
    address: true,
    phone: true,
    lease_contract_number: true,
    contract_date: true,
    last_tax_paid_date: true,
    tax_due_date: true,
    receipt_number: true,
    tax_amount: true,
    additional_amount: true,
    recorder: true,
    registrar: true,
    last_update_date: true,
    picture_front: true, 
    picture_back: true, 
    picture_side: true,
  };

  let provinceList = ['กรุงเทพฯ', 'นนทบุรี', 'สมุทรปราการ']; 
  let typeList = ['รถบรรทุก', 'รถตู้', 'รถอเนกประสงค์']; 
  let characteristicList = ['รถส่วนตัว', 'รถบริษัท', 'รถสาธารณะ'];
  let modelList= ['Camry','Jazz','Atto'];
  let brandList = ['Toyota', 'Honda', 'Tesla']; 
  let years = [2022, 2023, 2024]; 
  let colors = ['ดำ', 'ขาว', 'แดง'];
  let cylinderCounts = [4, 6, 8]; 
  let ccList = [1000, 3000, 3500]; 
  let axleAndWheelCountList = [2, 4, 6]; 
  let seatCountList = [2, 4, 5];

  let showModal = false;
  let modalMessage = '';

  const handleInputChange = (event) => {
    const { name, value } = event.target;
    inputs[name] = value;
    inputValidity[name] = true;

    if (name === 'license_number' && value.length >=5) {
      handleExistData(value)
    }
  };

  const handleExistData = async (license_number) => {
  try {
    const response = await axios.get(`${PUBLIC_API_URL}/data/${license_number}`);
    if (response) {
      isExistData = true;
      existData = response.data.data;
      inputs = JSON.parse(JSON.stringify(existData));
      inputs.created_by = `${staff.first_name} ${staff.last_name}`
      inputs.updated_by = `${staff.first_name} ${staff.last_name}`
      inputs.staff_id = staff.id

      inputs.display_date_registration = formatDate(inputs.date_registration);
      inputs.display_ownership_date = formatDate(inputs.ownership_date);
      inputs.display_birth_date = formatDate(inputs.birth_date);
      inputs.display_contract_date = formatDate(inputs.contract_date);
      inputs.display_last_tax_paid_date = formatDate(inputs.last_tax_paid_date);
      inputs.display_tax_due_date = formatDate(inputs.tax_due_date);
      inputs.display_last_update_date = formatDate(inputs.last_update_date);

    }
  } catch (error) {
    console.log("Not Found Exist License Number");
  }
};

  const handleValidateAndSubmit = () => {
    let isValid = true;
    for (const key in inputs) {
      if (inputs[key] === '') {
        inputValidity[key] = false;
        isValid = false;
      } else {
      }
    }

    if (isValid && !isExistData) {
      handleSaveNewData();
    } else if (isValid && isExistData) {
      // Handle case where isValid is true and isExistData is true
      handleSaveEditedData();
    } else {
      showModal = true;
      modalMessage = 'กรุณากรอกข้อมูลให้ครบถ้วน';
    }
  };

  const handleSaveNewData = async () => {
  try {
    const cleanedInputs = {
      date_registration: inputs.date_registration,
      license_number: inputs.license_number,
      province: inputs.province,
      type: inputs.type,
      characteristic: inputs.characteristic,
      brand: inputs.brand,
      model: inputs.model,
      year: Number(inputs.year),
      color: inputs.color,
      vehicle_number: inputs.vehicle_number,
      location_vehicle: inputs.location_vehicle,
      engine_brand: inputs.engine_brand,
      engine_number: inputs.engine_number,
      location_engine: inputs.location_engine,
      fuel: inputs.fuel,
      gas_tank_number: inputs.gas_tank_number,
      cylinder_count: Number(inputs.cylinder_count),
      cc: Number(inputs.cc), 
      horsepower: Number(inputs.horsepower), 
      axle_and_wheel_count: Number(inputs.axle_and_wheel_count), 
      vehicle_weight: Number(inputs.vehicle_weight), 
      load_weight: Number(inputs.load_weight), 
      total_weight: Number(inputs.total_weight), 
      seat_count: Number(inputs.seat_count),
      latest_owner_number: inputs.latest_owner_number,
      ownership_date: inputs.ownership_date,
      owner: inputs.owner,
      id_card_number: inputs.id_card_number,
      birth_date: inputs.birth_date,
      nationality: inputs.nationality,
      address: inputs.address,
      phone: inputs.phone,
      lease_contract_number: inputs.lease_contract_number,
      contract_date: inputs.contract_date,
      last_tax_paid_date: inputs.last_tax_paid_date,
      tax_due_date: inputs.tax_due_date,
      receipt_number: inputs.receipt_number,
      tax_amount: Number(inputs.tax_amount), // Ensure tax_amount is a number
      additional_amount: Number(inputs.additional_amount), // Ensure additional_amount is a number
      recorder: inputs.recorder,
      registrar: inputs.registrar,
      last_update_date: inputs.last_update_date,
      picture_front: inputs.picture_front,
      picture_back: inputs.picture_back,
      picture_side: inputs.picture_side,
      created_by: inputs.created_by,
      updated_by: inputs.updated_by,
      staff_id: inputs.staff_id
    };

    const response = await axios.post(`${PUBLIC_API_URL}/data`, cleanedInputs);
    if (response) {
      showModal = true;
      modalMessage = 'บันทึกข้อมูลสำเร็จ';
      setTimeout(() => {
        goto('/menu');
      }, 2000);
    }
  } catch (error) {
    showModal = true;
    modalMessage = 'บันทึกข้อมูลไม่สำเร็จ โปรดลองใหม่อีกครั้ง';
  }
};

const handleSaveEditedData = async () => {
  try {
    let effect_fields = [];
    for (const key in inputs) {
      if (inputs[key] !== existData[key] && !key.startsWith('display_') && key != 'created_by'&& key != 'staff_id') {
        effect_fields.push(key);
      }
    }

    const cleanedInputs = {
      date_registration: inputs.date_registration,
      license_number: inputs.license_number,
      province: inputs.province,
      type: inputs.type,
      characteristic: inputs.characteristic,
      brand: inputs.brand,
      model: inputs.model,
      year: Number(inputs.year),
      color: inputs.color,
      vehicle_number: inputs.vehicle_number,
      location_vehicle: inputs.location_vehicle,
      engine_brand: inputs.engine_brand,
      engine_number: inputs.engine_number,
      location_engine: inputs.location_engine,
      fuel: inputs.fuel,
      gas_tank_number: inputs.gas_tank_number,
      cylinder_count: Number(inputs.cylinder_count),
      cc: Number(inputs.cc),
      horsepower: Number(inputs.horsepower),
      axle_and_wheel_count: Number(inputs.axle_and_wheel_count),
      vehicle_weight: Number(inputs.vehicle_weight),
      load_weight: Number(inputs.load_weight),
      total_weight: Number(inputs.total_weight),
      seat_count: Number(inputs.seat_count),
      latest_owner_number: inputs.latest_owner_number,
      ownership_date: inputs.ownership_date,
      owner: inputs.owner,
      id_card_number: inputs.id_card_number,
      birth_date: inputs.birth_date,
      nationality: inputs.nationality,
      address: inputs.address,
      phone: inputs.phone,
      lease_contract_number: inputs.lease_contract_number,
      contract_date: inputs.contract_date,
      last_tax_paid_date: inputs.last_tax_paid_date,
      tax_due_date: inputs.tax_due_date,
      receipt_number: inputs.receipt_number,
      tax_amount: Number(inputs.tax_amount),
      additional_amount: Number(inputs.additional_amount),
      recorder: inputs.recorder,
      registrar: inputs.registrar,
      last_update_date: inputs.last_update_date,
      picture_front: inputs.picture_front,
      picture_back: inputs.picture_back,
      picture_side: inputs.picture_side,
      updated_by: inputs.updated_by,
      staff_id: inputs.staff_id,
      effect_field: effect_fields.join('|') // Join the changed fields with '|'
    };
    const response = await axios.patch(`${PUBLIC_API_URL}/data/${inputs.id}`, cleanedInputs);
    if (response) {
      showModal = true;
      modalMessage = 'บันทึกข้อมูลที่แก้ไขสำเร็จ';
      setTimeout(() => {
        goto('/menu');
      }, 2000);
    }
  } catch (error) {
    showModal = true;
    modalMessage = 'บันทึกข้อมูลไม่สำเร็จ โปรดลองใหม่อีกครั้ง';
  }
};


  const navigateToMenu = () => {
    goto('/menu');
  };

  const fillWithMockData = () => {
    inputs = {
      date_registration: new Date('2024-06-07T00:00:00Z').toISOString(),
      display_date_registration: '07/06/2567',
      license_number: 'บบ1234',
      province: 'กรุงเทพฯ',
      type: 'รถบรรทุก',
      characteristic: 'รถส่วนตัว',
      brand: 'Toyota',
      model: 'Camry',
      year: 2023,
      color: 'ดำ',
      vehicle_number: '5678',
      location_vehicle: 'Bangkok',
      engine_brand: 'Toyota',
      engine_number: '91011',
      location_engine: 'Bangkok',
      fuel: 'Gasoline',
      gas_tank_number: '121314',
      cylinder_count: 4,
      cc: 3000,
      horsepower: 200,
      axle_and_wheel_count: 4,
      vehicle_weight: 1500.0,
      load_weight: 500,
      total_weight: 2000,
      seat_count: 4,
      latest_owner_number: '1',
      ownership_date: new Date('2024-06-07T00:00:00Z').toISOString(),
      display_ownership_date: '07/06/2567',
      owner: 'John Doe',
      id_card_number: '1234567890123',
      birth_date: new Date('2024-06-07T00:00:00Z').toISOString(),
      display_birth_date: '07/06/2567',
      nationality: 'Thai',
      address: '123 Bangkok Street',
      phone: '0812345678',
      lease_contract_number: '12345',
      contract_date: new Date('2024-06-07T00:00:00Z').toISOString(),
      display_contract_date: '07/06/2567',
      last_tax_paid_date: new Date('2024-06-07T00:00:00Z').toISOString(),
      display_last_tax_paid_date: '07/06/2567',
      tax_due_date: new Date('2024-06-07T00:00:00Z').toISOString(),
      display_tax_due_date: '07/06/2567',
      receipt_number: '67890',
      tax_amount: 1000.0,
      additional_amount: 200.0,
      recorder: 'Jane Smith',
      registrar: 'Tom Johnson',
      last_update_date: new Date('2024-06-07T00:00:00Z').toISOString(),
      display_last_update_date: '07/06/2567',
      created_by: `${staff.first_name} ${staff.last_name}`,
      updated_by: `${staff.first_name} ${staff.last_name}`,
      staff_id: staff.id,
      picture_front: '',
      picture_back: '',
      picture_side: ''
    };
  };

  onMount(() => {
    isExistData = false
    staff = getObject('staff');
    if (staff == null) {
      goto('/');
    }
    inputs.created_by = `${staff.first_name} ${staff.last_name}`
    inputs.updated_by = `${staff.first_name} ${staff.last_name}`
    inputs.staff_id = staff.id
  });
</script>

<div class="container mx-auto p-4">
  <div class="flex justify-end">
    <button type="button" class="py-2 px-4 bg-blue-500 text-white rounded-md" on:click={fillWithMockData}>Demo</button>
  </div>
  <form>
    <InputField
      label="เลขทะเบียน"
      name="license_number"
      bind:value={inputs.license_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.license_number}
    />
    <Calendar
      label="วันที่จดทะเบียน"
      name="date_registration"
      bind:selectedDate={inputs.date_registration}
      bind:displayDate={inputs.display_date_registration}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.date_registration}
    />
    <CustomDropdown
      label="จังหวัด"
      name="province"
      bind:value={inputs.province}
      options={provinceList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.province}
    />
    <CustomDropdown
      label="ประเภท"
      name="type"
      bind:value={inputs.type}
      options={typeList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.type}
    />
    <CustomDropdown
      label="ลักษณะ"
      name="characteristic"
      bind:value={inputs.characteristic}
      options={characteristicList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.characteristic}
    />
    <CustomDropdown
      label="ยี่ห้อรถ"
      name="brand"
      bind:value={inputs.brand}
      options={brandList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.brand}
    />
    <CustomDropdown
      label="แบบ"
      name="model"
      bind:value={inputs.model}
      options={modelList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.characteristic}
    />
    <CustomDropdown
      label="รุ่นปี"
      name="year"
      bind:value={inputs.year}
      options={years}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.year}
    />
    <CustomDropdown
      label="สี"
      name="color"
      bind:value={inputs.color}
      options={colors}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.color}
    />
    <InputField
      label="เลขตัวรถ"
      name="vehicle_number"
      bind:value={inputs.vehicle_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.vehicle_number}
    />
    <InputField
      label="อยู่ที่"
      name="location_vehicle"
      bind:value={inputs.location_vehicle}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.location_vehicle}
    />
    <InputField
      label="ยี่ห้อเครื่องยนต์"
      name="engine_brand"
      bind:value={inputs.engine_brand}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.engine_brand}
    />
    <InputField
      label="เลขเครื่องยนต์"
      name="engine_number"
      bind:value={inputs.engine_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.engine_number}
    />
    <InputField
      label="อยู่ที่"
      name="location_engine"
      bind:value={inputs.location_engine}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.location_engine}
    />
    <InputField
      label="เชื้อเพลิง"
      name="fuel"
      bind:value={inputs.fuel}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.fuel}
    />
    <InputField
      label="เลขถังแก๊ส"
      name="gas_tank_number"
      bind:value={inputs.gas_tank_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.gas_tank_number}
    />
    <CustomDropdown
      label="จำนวนสูบ"
      name="cylinder_count"
      bind:value={inputs.cylinder_count}
      options={cylinderCounts}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.cylinder_count}
    />
    <CustomDropdown
      label="ซีซี"
      name="cc"
      bind:value={inputs.cc}
      options={ccList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.cc}
    />
    <InputField
      label="แรงม้า"
      name="horsepower"
      bind:value={inputs.horsepower}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.horsepower}
    />
    <CustomDropdown
      label="จำนวนเพลาและล้อ"
      name="axle_and_wheel_count"
      bind:value={inputs.axle_and_wheel_count}
      options={axleAndWheelCountList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.axle_and_wheel_count}
    />
    <InputField
      label="น้ำหนักรถ(กก.)"
      name="vehicle_weight"
      bind:value={inputs.vehicle_weight}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.vehicle_weight}
    />
    <InputField
      label="น้ำหนักบรรทุก/น้ำหนักลงเพลา(กก.)"
      name="load_weight"
      bind:value={inputs.load_weight}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.load_weight}
    />
    <InputField
      label="น้ำหนักรวม(กก.)"
      name="total_weight"
      bind:value={inputs.total_weight}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.total_weight}
    />
    <CustomDropdown
      label="ที่นั่ง(คน)"
      name="seat_count"
      bind:value={inputs.seat_count}
      options={seatCountList}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.seat_count}
    />
    <InputField
      label="ลำดับที่เข้าของรถคนล่าสุด"
      name="latest_owner_number"
      bind:value={inputs.latest_owner_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.latest_owner_number}
    />
    <Calendar
      label="วันที่ครอบครองรถ"
      name="ownership_date"
      bind:selectedDate={inputs.ownership_date}
      bind:displayDate={inputs.display_ownership_date}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.ownership_date}
    />
    <InputField
      label="ผู้ถือกรรมสิทธิ์"
      name="owner"
      bind:value={inputs.owner}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.owner}
    />
    <InputField
      label="เลขที่บัตร"
      name="id_card_number"
      bind:value={inputs.id_card_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.id_card_number}
    />
    <Calendar
      label="วันเกิด"
      name="birth_date"
      bind:selectedDate={inputs.birth_date}
      bind:displayDate={inputs.display_birth_date}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.birth_date}
    />
    <InputField
      label="สัญชาติ"
      name="nationality"
      bind:value={inputs.nationality}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.nationality}
    />
    <InputField
      label="ที่อยู่"
      name="address"
      bind:value={inputs.address}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.address}
    />
    <InputField
      label="โทร"
      name="phone"
      bind:value={inputs.phone}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.phone}
    />
    <InputField
      label="สัญญาเช่าซื้อเลขที่"
      name="lease_contract_number"
      bind:value={inputs.lease_contract_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.lease_contract_number}
    />
    <Calendar
      label="ลงวันที่"
      name="contract_date"
      bind:selectedDate={inputs.contract_date}
      bind:displayDate={inputs.display_contract_date}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.contract_date}
    />
    <Calendar
      label="วันที่จ่ายภาษีล่าสุด"
      name="last_tax_paid_date"
      bind:selectedDate={inputs.last_tax_paid_date}
      bind:displayDate={inputs.display_last_tax_paid_date}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.last_tax_paid_date}
    />
    <Calendar
      label="วันครบกำหนดเสียภาษี"
      name="tax_due_date"
      bind:selectedDate={inputs.tax_due_date}
      bind:displayDate={inputs.display_tax_due_date}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.tax_due_date}
    />
    <InputField
      label="ใบเสร็จรับเงินเลขที่/เลขที่"
      name="receipt_number"
      bind:value={inputs.receipt_number}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.receipt_number}
    />
    <InputField
      label="ค่าภาษี บาท/สต."
      name="tax_amount"
      bind:value={inputs.tax_amount}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.tax_amount}
    />
    <InputField
      label="เงินเพิ่ม บาท/สต."
      name="additional_amount"
      bind:value={inputs.additional_amount}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.additional_amount}
    />
    <InputField
      label="ผู้บันทึก"
      name="recorder"
      bind:value={inputs.recorder}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.recorder}
    />
    <InputField
      label="นายทะเบียน"
      name="registrar"
      bind:value={inputs.registrar}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.registrar}
    />
    <Calendar
      label="วันที่บันทึกล่าสุด"
      name="last_update_date"
      bind:selectedDate={inputs.last_update_date}
      bind:displayDate={inputs.display_last_update_date}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.last_update_date}
    />
    <UploadPicture
      label="ภาพรถด้านหน้า"
      name="picture_front"
      bind:value={inputs.picture_front}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.picture_front}
    />
    <UploadPicture
      label="ภาพรถด้านหลัง"
      name="picture_back"
      bind:value={inputs.picture_back}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.picture_back}
    />
    <UploadPicture
      label="ภาพรถด้านข้าง"
      name="picture_side"
      bind:value={inputs.picture_side}
      handleInputChange={handleInputChange}
      isInvalid={!inputValidity.picture_side}
    />

    <div class="flex justify-between mt-4 space-x-2">
      <button type="button" class="py-4 px-8 bg-gray-300 text-white rounded-md w-40" on:click={navigateToMenu}>ยกเลิก</button>
      <button type="submit" class="py-4 px-8 bg-orange-500 text-white rounded-md w-40" on:click={handleValidateAndSubmit}>บันทึก</button>
    </div>
  </form>

  <Modal 
    message={modalMessage} 
    isVisible={showModal} 
    onConfirm={() => {
      showModal = false 
      // goto('/menu')
    }} />
</div>
