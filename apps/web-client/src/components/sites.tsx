import AddForm from "./AddForm";
import StatusBadge from "./badge";
import DeleteButton from "./deleteButton";

const sites = [
  {
    name: "test",
    url: "test.com",
    status: { up: true },
  },
  {
    name: "google",
    url: "google.com",
    status: { up: false },
  },
  {
    name: "kirre",
    url: "google.com",
  },
];

const doDelete = () => {
  console.log("delete");
};

function SiteList() {
  return (
    <div className="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
      <div className="flex justify-between items-center px-4 py-2 bg-gray-50">
        <AddForm />
      </div>
      <table className="min-w-full divide-y divide-gray-300">
        <thead className="bg-gray-50">
          <tr>
            <th
              scope="col"
              className="px-3 py-5 text-left text-sm font-semibold text-gray-900"
            />
            <th scope="col" className="relative py-3.5 pl-3 pr-4 sm:pr-6">
              <span className="sr-only"></span>
            </th>
          </tr>
        </thead>
        {sites.map((site, index) => (
          <tbody key={index} className="bg-white divide-y divide-gray-200">
            <tr>
              <td className="px-3 py-4">
                <div>
                  <div className="flex items-center gap-2">
                    <span className="text-gray-900 font-semibold text-base">
                      {site.name}
                    </span>
                    <StatusBadge status={site.status} />
                  </div>
                  <div className="flex items-center gap-2">
                    <span className="text-sm text-gray-500">{site.url}</span>
                  </div>
                </div>
              </td>
              <td className="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
                <DeleteButton onClick={doDelete} />
              </td>
            </tr>
          </tbody>
        ))}
      </table>
    </div>
  );
}

export default SiteList;
