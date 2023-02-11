import StatusFilter from "./StatusFilter";
import CountryFilter from "./CountryFilter";

export default function FiltersContainer() {
  return (
    <div className="flex flex-col lg:flex-row lg:items-center gap-y-4 lg:gap-x-4 mb-12">
      <StatusFilter />
      <CountryFilter />
    </div>
  );
}
