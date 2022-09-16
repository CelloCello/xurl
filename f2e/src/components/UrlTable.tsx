import { deleteLink, fetchLinks } from '../api/links';
import { LinkPayload } from '../types';
import { useStore } from '../store';
import { ROOT_URL } from '../constants';

interface Props {
  links: LinkPayload[];
}

const UrlTable = ({ links }: Props) => {
  const [setLinks] = useStore((state) => [state.setLinks]);

  const onCopy = (code: string) => {
    navigator.clipboard.writeText(`${ROOT_URL}/${code}`);
  };

  const onDelete = (link: LinkPayload) => {
    deleteLink(link.id).then((isSuccess) => {
      fetchLinks().then((payload) => {
        setLinks(payload);
      });
    });
  };

  return (
    <div className="overflow-x-auto relative shadow-md sm:rounded-lg">
      <table className="w-full table-fixed text-sm text-left text-gray-500 dark:text-gray-400">
        <thead className="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
          <tr>
            <th scope="col" className="w-1/6 py-3 px-6">
              Code
            </th>
            <th scope="col" className="w-4/6 py-3 px-6">
              Link
            </th>
            <th scope="col" className="w-1/6 py-3 px-6">
              <span className="sr-only">Edit</span>
            </th>
          </tr>
        </thead>
        <tbody>
          {links.map((link) => (
            <tr
              key={link.code}
              className="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
            >
              <td className="py-2 px-4">
                <a href={`${ROOT_URL}/${link.code}`} target="_blank" rel="noreferrer">
                  {link.code}
                </a>
              </td>
              <td
                scope="row"
                className="py-2 px-4 break-all font-medium text-gray-900 dark:text-white"
              >
                <a href={link.url} target="_blank" rel="noreferrer">
                  {link.url}
                </a>
              </td>
              <td className="flex py-2 px-4 items-center content-center justify-center">
                {/* copy button */}
                <a
                  href="#"
                  className="font-medium text-blue-600 dark:text-blue-500 hover:underline"
                  onClick={() => {
                    onCopy(link.code);
                  }}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    className="w-5 h-5"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M16.5 8.25V6a2.25 2.25 0 00-2.25-2.25H6A2.25 2.25 0 003.75 6v8.25A2.25 2.25 0 006 16.5h2.25m8.25-8.25H18a2.25 2.25 0 012.25 2.25V18A2.25 2.25 0 0118 20.25h-7.5A2.25 2.25 0 018.25 18v-1.5m8.25-8.25h-6a2.25 2.25 0 00-2.25 2.25v6"
                    />
                  </svg>
                </a>

                {/* delete button */}
                <a
                  href="#"
                  className="font-medium text-red-600 dark:text-red-500 hover:text-red-700"
                  onClick={() => {
                    onDelete(link);
                  }}
                >
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    strokeWidth="1.5"
                    stroke="currentColor"
                    className="w-5 h-5"
                  >
                    <path
                      strokeLinecap="round"
                      strokeLinejoin="round"
                      d="M14.74 9l-.346 9m-4.788 0L9.26 9m9.968-3.21c.342.052.682.107 1.022.166m-1.022-.165L18.16 19.673a2.25 2.25 0 01-2.244 2.077H8.084a2.25 2.25 0 01-2.244-2.077L4.772 5.79m14.456 0a48.108 48.108 0 00-3.478-.397m-12 .562c.34-.059.68-.114 1.022-.165m0 0a48.11 48.11 0 013.478-.397m7.5 0v-.916c0-1.18-.91-2.164-2.09-2.201a51.964 51.964 0 00-3.32 0c-1.18.037-2.09 1.022-2.09 2.201v.916m7.5 0a48.667 48.667 0 00-7.5 0"
                    />
                  </svg>
                </a>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
};

export default UrlTable;
