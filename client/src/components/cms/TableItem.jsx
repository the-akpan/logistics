import React from "react";

export default function TableItem() {
  return (
    <div className="px-3 hover:rounded-md hover:bg-gunmetal transition-all duration-100 border-b py-4 cursor-pointer ">
      <div className="flex gap-x-12 lg:gap-x-4  text-left items-center">
        <p className=" basis-1/4">1234FDSA</p>
        <div className=" basis-1/4">
          <p>Alvin Chinedu</p>
        </div>
        <p className=" basis-1/4">
          <span className="text-sonic-silver text-sm font-semibold">
            America
          </span>
        </p>
        <p className=" basis-1/4">
          <span className="text-sonic-silver text-sm font-semibold">
            Nigeria
          </span>
        </p>
        <p className=" basis-1/4">
          <span className=" p-1.5 text-black bg-lavender-blue text-sm rounded-md font-semibold">
            Active
          </span>
        </p>
        <p className=" basis-1/4">21/04/2023</p>
        <span className=" basis-1/4 text-center">
          {" "}
          <button className="bg-slate-900 text-white text-sm px-3 py-2 rounded-lg">
            Edit
          </button>
        </span>
        <span className=" basis-1/4 text-center">
          {" "}
          <button className="bg-slate-900 text-white text-sm px-3 py-2 rounded-lg">
            Delete
          </button>
        </span>
      </div>
    </div>
  );
}
