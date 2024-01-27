import { SiteApi, AddSiteRequest } from "monitor-sdk";
import { ChangeEvent, useState } from "react";

function AddForm() {
    const [formOpen, setFormOpen] = useState(false)
    const [formValues, setFormValues] = useState({
        url: '',
        siteName: ''
    })

    const handleInput = (event: ChangeEvent<HTMLInputElement>) => {
        const {name, value} = event.target
        setFormValues({
            ...formValues,
            [name]: value
        })
    }

    const handleSiteSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        try {
            const api = new SiteApi();
            const requestData: AddSiteRequest = {
                siteAddRequest: formValues
            }
            await api.addSite(requestData);
            console.log('Site added successfully!');

            setFormValues({
                url: '',
                siteName: ''
            });
            setFormOpen(false);
        } catch (error) {
            console.error('Error adding site:', error);
        }
    };


    if(!formOpen) {
        return (
            <button 
                type="button"
                className="inline-flex items-center justify-center rounded-md border border-transparent bg-indigo-600 px-4 py-2 text-sm font-medium text-white shadow-sm hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 sm:w-auto"
                onClick={() => setFormOpen(true)}
                >
                    Add site
            </button>
        )
    }

    return (
        <form onSubmit={handleSiteSubmit}>
  <div className="flex flex-col md:flex-row md:items-end gap-4">
        <div>
          <input
            type="text"
            value={formValues.url}
            onChange={handleInput}
            placeholder="google.com"
            className="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-indigo-500 focus:ring-indigo-500 sm:text-sm"
          />
        </div>

        <div>
          <button
            type="submit"
            className="inline-flex justify-center rounded-md border border-transparent bg-indigo-600 py-2 px-4 text-sm font-medium text-white shadow-sm enabled:hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:ring-offset-2 disabled:opacity-75"
          >
            Save
          </button>
          <button type="button" onClick={() => setFormOpen(false)}>Cancel</button>
        </div>
      </div>
        </form>
    )
}

export default AddForm