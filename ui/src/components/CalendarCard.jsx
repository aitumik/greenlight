export const CalendarCard = ({icon,title,venue,description}) => {
  return (
    <div style={{backgroundColor: "#1c1e20"}} className="p-4 rounded-md">
      <div className="flex justify-between items-center pb-2">
        <p>Icon here</p>
        <div className="text-sm px-3 py-1 bg-neutral-950 rounded-full items-center">Subscribe</div>
      </div>
      <p className="font-semibold">{title}</p>
      <p className="text-sm text-neutral-500">{description}</p>
    </div>
  )
}