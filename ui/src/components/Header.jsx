import { useState,useEffect } from 'react'
export const Header = () => {

  const [scrolled,setScrolled] = useState(false)

  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 0)
    }

    window.addEventListener('scroll',handleScroll)

    return () => window.removeEventListener('scroll',handleScroll)
  },[])

  return (
    <nav className={`${scrolled ? 'bg-blue/30 backdrop-blur-md border-b border-white/20 shadow-sm' : 'bg-transparent border-b border-white/20 shadow-sm'} px-[60px] py-5 w-full z-50 fixed top-0 transition-all duration-300`}>
      <div className="flex justify-between">
        <div className="flex gap-6">
          <p className="text-neutral-200 text-sm font-semibold">Events</p>
          <p className="text-neutral-500 text-sm font-semibold">Calendars</p>
          <p className="text-neutral-500 text-sm font-semibold">Discover</p>
          <p className="text-neutral-500 text-sm font-semibold">Movies</p>
        </div>
        <div className="flex gap-4">
          <div className="rounded-md bg-neutral-950 py-1.5 px-3 text-neutral-200 text-xs font-semibold">Create Event</div>
        </div>
      </div>
    </nav>
  )
}
