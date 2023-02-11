export default function StatusFilter() {
  return (
    <select
      className="cursor-pointer border-2 border-slate-800 font-semibold py-2 px-6 text-sonic-silver transition-all duration-100"
      name="status-filter"
      id="chart filter"
    >
      <option selected={true} value="">
        Filter by status
      </option>
      <option value="pending">Pending</option>
      <option value="canceled">Canceled</option>
      <option value="delivered">Delivered</option>
      <option value="in-transit">In Transit</option>
    </select>
  );
}
