import { useState, FormEvent, ChangeEvent } from "react";
import { Dialog } from "@headlessui/react";
import { PlusCircleIcon } from "@heroicons/react/24/outline";

function AddForm() {
  const [isOpen, setIsOpen] = useState(false);
  const [formData, setFormData] = useState({
    name: "",
    url: "",
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log(formData);
    setIsOpen(false);
    setFormData({
        name:'',
        url:''
    })
  };

  return (
    <div>
      <button
        className="fixed bottom-8 right-8 z-10 bg-blue-500 text-white p-4 rounded-full shadow-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
        onClick={() => setIsOpen(true)}
      >
        <PlusCircleIcon className="h-5 w-5" />
      </button>
      <Dialog open={isOpen} onClose={() => setIsOpen(false)} className="fixed inset-0 z-10 overflow-y-auto flex items-center justify-center">
        <Dialog.Overlay className="fixed inset-0 bg-black opacity-30" />
        <div className="bg-white p-6 rounded-lg shadow-xl max-w-md w-full relative z-20">
          <Dialog.Title className="text-lg font-semibold mb-4">Add Site</Dialog.Title>
          <Dialog.Description>
            <form onSubmit={handleSubmit}>
              <div className="mb-4">
                <label htmlFor="name" className="block text-sm font-medium text-gray-700">Name</label>
                <input
                  type="text"
                  id="name"
                  name="name"
                  value={formData.name}
                  onChange={handleChange}
                  className="mt-1 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-blue-400 w-full"
                />
              </div>
              <div className="mb-4">
                <label htmlFor="url" className="block text-sm font-medium text-gray-700">URL</label>
                <input
                  type="url"
                  id="url"
                  name="url"
                  value={formData.url}
                  onChange={handleChange}
                  className="mt-1 p-2 border border-gray-300 rounded-md focus:outline-none focus:ring focus:ring-blue-400 w-full"
                />
              </div>
              <button
                type="submit"
                className="inline-block px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 focus:outline-none focus:ring focus:ring-blue-400"
              >
                Submit
              </button>
            </form>
          </Dialog.Description>
        </div>
      </Dialog>
    </div>
  );
}

export default AddForm;
