import { TrashIcon } from "@heroicons/react/24/outline";



interface DeleteButtonProps {
  onClick: () => void; // Define the type of onClick prop
}

function DeleteButton({ onClick }: DeleteButtonProps) {
  return (
    <button
      onClick={onClick}
      className="text-red-400 hover:text-red-800"
    >
      <TrashIcon className="h-5 w-5" />
    </button>
  );
}

export default DeleteButton;