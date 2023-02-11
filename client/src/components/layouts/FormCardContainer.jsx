import React from "react";

const FormCardContainer = ({ children }) => {
  return (
    <div className="my-20 flex flex-1 flex-grow items-center justify-center">
      {children}
    </div>
  );
};

export default FormCardContainer;
