import { useEffect, useState } from "react";

const Shortener = () => {
  useEffect(() => {}, []);

  return (
    <div className="my-4">
      <input className="placeholder:text-slate-400 block bg-white w-full border border-slate-300 rounded-md py-2 pl-9 pr-3 shadow-sm focus:outline-none focus:border-sky-500 focus:ring-sky-500 focus:ring-1 sm:text-sm" placeholder="http(s)://..." type="text"/>
      <button className="ml-2 my-4 border-black">Create</button>
      <button className="ml-2 my-4 border-black">Clear</button>
    </div>
  );
};

export default Shortener;
