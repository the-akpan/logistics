import { useState } from "react";

import { makeRequest, Methods, URLS } from "../../api/api";

export default function LoginForm() {
  const [email, setEmail] = useState("");
  const [pwd, setPwd] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await makeRequest(URLS.auth.login, Methods.POST, {
        email: email,
        password: pwd,
      });
      localStorage.setItem("token", JSON.stringify(response?.data?.token));

      console.log(response);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <form
      onSubmit={handleSubmit}
      className="bg-slate-200 py-10 rounded-lg px-8 w-4/5 lg:w-2/5 shadow-2xl"
    >
      <h3 className="font-bold text-2xl text-center text-slate-800 mb-6">
        Login
      </h3>
      <div className="flex flex-col space-y-4">
        <input
          onChange={(e) => setEmail(e.target.value)}
          className="py-3 px-3"
          type="text"
          placeholder="Enter email"
        />
        <input
          onChange={(e) => setPwd(e.target.value)}
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
