import { useState, useEffect } from 'react'
import { Link } from 'react-router-dom'

export const Header = () => {

  const [scrolled, setScrolled] = useState(false)

  useEffect(() => {
    const handleScroll = () => {
      setScrolled(window.scrollY > 0)
    }

    window.addEventListener('scroll', handleScroll)

    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  return (
    <nav className={`${scrolled ? 'bg-blue/30 backdrop-blur-md border-b border-white/20 shadow-sm' : 'bg-transparent border-b border-white/20 shadow-sm'} px-[60px] py-5 w-full z-50 fixed top-0 transition-all duration-300`}>
      <div className="flex justify-between px-[250px]">
        <div className="flex gap-6">
          <Link to="/explore" className="text-neutral-200 text-base font-semibold">Explore</Link>
          <Link to="/events" className="text-neutral-500 text-base font-semibold">Events</Link>
          <Link to="/calendars" className="text-neutral-500 text-base font-semibold">Calendars</Link>
          <Link to="/movies" className="text-neutral-500 text-base font-semibold">Movies</Link>
        </div>
        <div className="flex gap-4">
          <div className="rounded-md bg-neutral-950 py-1.5 px-3 text-neutral-200 text-base font-semibold">Create Account</div>
        </div>
      </div>
    </nav>
  )
}
