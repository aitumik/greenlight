import { useModalStore } from "./modalStore"

export const Modal = () => {
    const { isOpen, content, closeModal } = useModalStore()
    if (!isOpen) return null
    return (
        <div className="fixed inset-0 z-50 flex items-center justify-center bg-black bg-opacity-50">
            <div className="bg-white rounded shadow-lg p-6 min-w-[300px] max-w-full relative">
                <button
                    className="absolute top-2 right-2 text-lg font-bold text-gray-600 hover:text-black"
                    onClick={closeModal}
                    aria-label="Close"
                >
                    ×
                </button>
                {typeof content === "function" ? content() : content}
            </div>
        </div>
    )
}
