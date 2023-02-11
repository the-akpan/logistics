import React from "react";

export default function Pagination() {
  return (
    <div className="flex items-center justify-center sm:gap-x-4 mx-auto mt-8 text-sm">
      <button className="px-3 py-1"> Previous </button>
      <button className="px-3 py-1 border-2 hover:bg-slate-600 hover:text-white border-slate-500">
        1
      </button>
      <button className="px-3 py-1 border-2 hover:bg-slate-600 hover:text-white border-slate-500">
        2
      </button>
      <button className="px-3 py-1 border-2 hover:bg-slate-600 hover:text-white border-slate-500">
        3
      </button>
      <button className="px-3 py-1 border-2 hover:bg-slate-600 hover:text-white border-slate-500">
        4
      </button>
      <button className="px-3 py-1 border-2 hover:bg-slate-600 hover:text-white border-slate-500">
        5
      </button>
      <button className="px-3 py-1"> Next </button>
    </div>
  );
}
