import { useState, FormEvent, ChangeEvent } from "react";
import { Dialog } from "@headlessui/react";
import { PlusCircleIcon } from "@heroicons/react/24/outline";

function MyForm() {
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
  };

  return (
    <div>
      <button
        className="text-blue-600 hover:text-blue-900"
        onClick={() => setIsOpen(true)}
      >
        <PlusCircleIcon className="h-5 w-5" />
      </button>
      <Dialog open={isOpen} onClose={() => setIsOpen(false)}>
        <Dialog.Title>Dialog Title</Dialog.Title>
        <Dialog.Description>
          <form onSubmit={handleSubmit}>
            <div>
              <label htmlFor="email">Email</label>
              <input
                type="text"
                id="id"
                name="name"
                value={formData.name}
                onChange={handleChange}
              />
            </div>
            <div>
              <label htmlFor="password">Password</label>
              <input
                type="url"
                id="url"
                name="url"
                value={formData.url}
                onChange={handleChange}
              />
            </div>
            <button type="submit">Submit</button>
          </form>
        </Dialog.Description>
      </Dialog>
    </div>
  );
}

export default MyForm;
