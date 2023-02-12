import React from "react";

export default function FeaturesCard({ feature }) {
  return (
    <div class="items-centers flex gap-x-4 rounded-lg bg-primary py-7 px-9 font-bold text-white">
      <p>{feature}</p>
    </div>
  );
}
