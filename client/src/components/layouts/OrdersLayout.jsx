import { FiltersContainer, TableContainer, Pagination } from "../cms";

export default function OrdersLayout() {
  return (
    <>
      <h2 className="text-3xl font-bold mt-12 mb-12">Manage Orders</h2>
      <FiltersContainer />
      <div className="overflow-x-scroll">
        <TableContainer />
      </div>

      <Pagination />
    </>
  );
}
