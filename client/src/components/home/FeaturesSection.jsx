import FeaturesCard from "./FeaturesCard";

export default function () {
  const features = [
    "Primary care for several medical conditions",
    "Quick access to medical specialists",
    "Bookings for in house health services",
    "Affordable and reliable healthcare services",
    "Quick and affordable pharmaceutical services",
    "Access to health education and tips",
  ];
  return (
    <>
      {" "}
      <div class="grid gap-x-5 gap-y-4 md:grid-cols-2 lg:grid-cols-3">
        {features.map((feature) => (
          <FeaturesCard feature={feature} />
        ))}
      </div>
    </>
  );
}
