import React from "react";

export default function CountryFilter() {
  return (
    <select
      className="cursor-pointer border-2 border-slate-800 font-semibold py-2 px-6 text-sonic-silver transition-all duration-100"
      name="country-filter"
    >
      <option selected={true} value="">
        Filter by country
      </option>
      <option value="pending">Nigeria</option>
      <option value="canceled">Angola</option>
      <option value="delivered">America</option>
      <option value="in-transit">China</option>
    </select>
  );
}
