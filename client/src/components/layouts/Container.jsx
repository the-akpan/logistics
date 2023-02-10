import React from 'react'

export default function Container({children}) {
  return (
    <section className="max-w-xs sm:max-w-xl md:max-w-3xl lg:max-w-6xl mx-auto">
      {children}
    </section>
  );
}
