import {useMutation, useQuery, useQueryClient } from "@tanstack/react-query";
import AddForm from "./AddForm";
import DeleteButton from "./deleteButton";
import { createAPIBridge } from "../lib/apiBridge";
import StatusBadge from "./badge";
import TimeDelta from "./time";
import { DateTime } from "luxon";
import { Site } from "monitor-sdk/models";
import { DeleteSiteRequest } from "monitor-sdk";


function SiteList() {
  const { siteApi, statusApi } = createAPIBridge();

  const { isLoading, error, data } = useQuery({
    queryKey: ["sites"],
    queryFn: () => siteApi.getSites(),
    refetchInterval: 10000,
    retry: false,
  });


  const { data: status } = useQuery({
    queryKey: ["status"],
    queryFn: () => statusApi.status(),
    refetchInterval: 1000,
    retry: false,
  });

  const queryClient = useQueryClient();

  const doDelete = useMutation({
    mutationFn: (site: Site) => {
      const deleteRequest: DeleteSiteRequest = {
        id: site.id,
      };
      return siteApi.deleteSite(deleteRequest);
    },
    onSuccess: () => {
      queryClient.invalidateQueries({queryKey: ["sites"]})
    }
  })

  if (isLoading) {
    return <div>Loading sites...</div>;
  } else if (error) {
    return <div className="text-red-600">{(error as Error).message}</div>;
  }

  return (
    <>
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
        <tbody className="bg-white divide-y divide-gray-200">
          {data?.sites.length === 0 && (
            <tr>
              <td colSpan={2} className={"text-center text-gray-400 py-8"}>
                Nothing to monitor yet. Add a website to see it here.
              </td>
            </tr>
          )}
          {data!.sites.map((site) => {
            const st = status?.sites[site.id];
            const dt = st && DateTime.fromISO(st.checkedAt);
            return (
              <tr key={site.id}>
                <td className="px-3 py-4 text-sm">
                  <div className="flex items-center gap-2">
                    <span className="text-gray-700 font-semibold text-base">{site.name}</span>
                    <StatusBadge status={st} />
                  </div>
                  {dt && (
                    <div className="text-gray-400">
                      Last checked <TimeDelta dt={dt} />
                    </div>
                  )}
                </td>
                <td className="relative whitespace-nowrap py-4 pl-3 pr-4 text-right text-sm font-medium sm:pr-6">
                  <DeleteButton onClick={() => doDelete.mutate(site)} />
                </td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </div>
    </>
  );
}

export default SiteList;
