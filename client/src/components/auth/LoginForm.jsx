import React from "react";

export default function LoginForm() {
  return (
    <form className="bg-slate-200 py-10 rounded-lg px-8 w-4/5 lg:w-2/5 shadow-2xl">
      <h3 className="font-bold text-2xl text-center text-slate-800 mb-6">
        Login
      </h3>
      <div className="flex flex-col space-y-4">
        <input
          className="py-3 px-3"
          type="text"
          placeholder="Enter username or email"
        />
        <input
          className="py-3 px-3"
          type="password"
          placeholder="Enter password"
        />
        <button className="py-3 bg-primary text-white font-bold">Login</button>
      </div>

      <button className="mt-4 italic underline">Reset password</button>
    </form>
  );
}
