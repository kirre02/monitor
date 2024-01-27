import { useEffect, useState } from "react";
import { Site, SiteApi, Check} from "monitor-sdk";
import AddForm from "./AddForm";
import StatusBadge from "./badge";

function SiteList() {
  const [data, setSites] = useState<Site[]>([]);
  const [siteStatus] = useState<{ [key: string]: Check }>({});
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const api = new SiteApi();
        const response = await api.getSites();
        
        if (Array.isArray(response)) {
          setSites(response);
          setIsLoading(false);
        } else {
          throw new Error("Unexpected response format.");
        }
      } catch (error) {
        setError(error instanceof Error ? error.message : String(error));
        setIsLoading(false);
      }
    };

    fetchData();
  }, []);

  const handleDelete = async (site: Site) => {
     const id = site.id

     try {
      const api = new SiteApi()
      const response = await api.deleteSite({id})

      console.log("deleted site:", response.url)
     } catch (error) {
      console.error("error deleting site", error)
     }
  };

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (error) {
    return <div className="text-red-600">{error}</div>;
  }


  return (
    <>
    <div className="sm:flex sm:items-center">
        <div className="sm:flex-auto">
            <h1 className="text-xl font-semibold text-gray-900">
                Monitored Webdata
            </h1>
            <p className="mt-2 text-sm text-gray-700">
                A list of all the webdata being monitored, their current status,
                and when they were last checked.
            </p>
        </div>
        <div className="mt-4 sm:mt-0 sm:ml-16 sm:flex-none">
            <AddForm />
        </div>
    </div>

    <div className="mt-8 flex flex-col">
        <div className="-my-2 -mx-4 overflow-x-auto sm:-mx-6 lg:-mx-8">
            <div className="inline-block min-w-full py-2 align-middle md:px-6 lg:px-8">
                <div className="overflow-hidden shadow ring-1 ring-black ring-opacity-5 md:rounded-lg">
                    <table className="min-w-full divide-y divide-gray-300">
                        <thead className="bg-gray-50">
                            <tr>
                                <th
                                    scope="col"
                                    className="px-3 py-3.5 text-left text-sm font-semibold text-gray-900"
                                >
                                    Site
                                </th>
                                <th
                                    scope="col"
                                    className="relative py-3.5 pl-3 pr-4 sm:pr-6"
                                >
                                    <span className="sr-only"></span>
                                </th>
                            </tr>
                        </thead>
                        <tbody className="divide-y divide-gray-200 bg-white">
                            {data.length === 0 && (
                                <tr>
                                    <td
                                        colSpan={2}
                                        className={"text-center text-gray-400 py-8"}
                                    >
                                        Nothing to monitor yet. Add a website to see it here.
                                    </td>
                                </tr>
                            )}
                            {data.map((site) => {
                                const st = siteStatus[site.id];
                                return (
                                    <tr key={site.id}>
                                        <td className="px-3 py-4 text-sm">
                                            <div className="flex items-center gap-2">
                                                <span className="text-gray-700">{site.url}</span>
                                                <StatusBadge status={st} />
                                            </div>
                                        </td>
                                        <td className="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
                                            <button
                                                className="text-indigo-600 hover:text-indigo-900"
                                                onClick={() => handleDelete(site)}
                                            >
                                                Delete <span className="sr-only">{site.url}</span>
                                            </button>
                                        </td>
                                    </tr>
                                );
                            })}
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</>
  )
}   

export default SiteList