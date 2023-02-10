import React, { useEffect, useRef, useState } from "react";

export default function TrackingForm() {
  const [tracker, setTracker] = useState("");
  const trackerRef = useRef();
  useEffect(() => {
    trackerRef.current.focus();
  }, [])
  return (
    <form className="relative text-center z-20">
      <input
        className="py-6 outline-none border-none ring-none px-8 lg:w-[300px] text-neutral-500 font-semibold focus:border-0 focus:-outline-2 focus:outline-secondary"
        type="text"
        placeholder="Enter Tracking ID"
        value={tracker}
        ref={trackerRef}
        onChange={(e) => setTracker(e.target.value.toLocaleUpperCase())}
      />
      <button className="bg-secondary py-6 text-white font-bold px-8 hover:scale-105 transition duration-200">
        Track
      </button>
    </form>
  );
}
