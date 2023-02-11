import TableItem from "./TableItem";

export default function TableContainer() {
  return (
    <section>
      <div className="flex text-left items-center px-3 border-b-2 pb-5 font-bold gap-x-14 lg:gap-x-4 ">
        <p className=" lg:basis-1/4">Tracking ID</p>
        <p className=" lg:basis-1/4">Recipient Name</p>
        <p className=" lg:basis-1/4">From</p>
        <p className=" lg:basis-1/4">To</p>
        <p className=" lg:basis-1/4">Status</p>
        <p className=" lg:basis-1/4">Date Created</p>
        <p className=" lg:basis-1/4 text-center">Actions</p>
        <p className="basis-1/4"></p>
      </div>

      <TableItem />
      <TableItem />
      <TableItem />
      <TableItem />
      <TableItem />
      <TableItem />
      <TableItem />
      <TableItem />
    </section>
  );
}
