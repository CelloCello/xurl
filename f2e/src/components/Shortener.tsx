import { useEffect, useState, KeyboardEvent } from 'react';
import { createLink, fetchLinks } from '../api';

import { useStore } from '../store';

const Shortener = () => {
  const [orgUrl, setOrgUrl] = useState('');
  const [setLinks] = useStore((state) => [state.setLinks]);

  useEffect(() => {}, []);

  const onClear = () => {
    setOrgUrl('');
  };

  const onGenerate = () => {
    if (!orgUrl) return;
    createLink(orgUrl).then((link) => {
      fetchLinks().then((payload) => {
        setLinks(payload);
      });
    });
  };

  const onInputKeyPress = (e: KeyboardEvent) => {
    if (e.key === 'Enter') onGenerate();
    e.preventDefault();
  };

  return (
    <div className="my-4">
      <form className="flex items-center">
        <div className="relative w-full">
          <input
            type="text"
            className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg block w-full pl-4 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white "
            placeholder="https://"
            value={orgUrl}
            onChange={(e) => setOrgUrl(e.target.value)}
            onKeyPress={onInputKeyPress}
          />
          <button
            type="button"
            className="flex absolute inset-y-0 right-0 items-center pr-3 border-0 focus:outline-0"
            onClick={onClear}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              fill="none"
              viewBox="0 0 24 24"
              strokeWidth="1.5"
              stroke="currentColor"
              className="w-4 h-4 text-gray-500 dark:text-gray-400 hover:text-gray-900 dark:hover:text-white"
            >
              <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
            </svg>
          </button>
        </div>
        <button
          // type="submit"
          type="button"
          className="inline-flex items-center py-2.5 px-3 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
          onClick={onGenerate}
        >
          Generate
        </button>
      </form>
    </div>
  );
};

export default Shortener;
