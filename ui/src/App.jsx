import { useState } from 'react'
import { Header } from './components/Header'
import { EventCard } from './components/EventCard'

const Intro = () => {
  return (
    <div className="py-10 px-[60px]">
      <h1 className="text-3xl text-neutral-300 font-semibold">Discover Events</h1>
      <p className="text-white text-base py-[16px]">Explore popular events near you, browser by category, or check out some of the great community calendars</p>
    </div>
  )
}

const EventList = () => {
  const events = [
    {
      title: "Health Sector AI Roundtable 4.0",
      cover: "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=0,background=white,quality=75,width=300,height=300/event-covers/jg/a3c30fe6-55ea-4835-984c-127d1049b4ca.png",
      start_time: "Tue,Apr 8, 9:00 AM",
      venue: "Marsabit Plaza",
    },
    {
      title: "Heels & Green Speed Connect",
      cover: "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/no/affec154-182c-4bce-8b49-974396bf503f.png",
      start_time: "Thu,Apr 10, 6:30 PM",
      venue: "iHUB",
    },
    {
      title: "Launch of Mercury Contamination in Migori Research Study",
      cover: "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/6q/f059c56f-a1fe-46e5-b6a5-6c50b8e650a9.png",
      start_time: "Fri,Apr 11, 10:00 AM",
      venue: "Sarova Stanley Nairobi",
    },
    {
      title: "LW14 - Nairobi, Kenya - Supabase Meetup",
      cover: "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/gallery-images/j5/e0d9e03b-63b3-452b-8899-3fc8f848a4bb",
      start_time: "Sat,Apr 12, 1:00 PM",
      venue: "Nairobi Garage",
    },
    {
      title: "Sui Kenya Community MeetUp",
      cover: "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/45/7c45fc3a-5887-457f-8bf3-94be18f25995.png",
      start_time: "Sat,Apr 12, 2:00 PM",
      venue: "Nafasi Connection LLC",
    },
    {
      title: "The Wake Live Recording",
      cover: "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/0q/d5ffeaf8-0858-495f-885c-f4894de047f0.jpg",
      start_time: "Sat,Apr 19, 12:00 PM",
      venue: "Kenya Cultural Centre",
    },
  ]

  return (
    <div className="grid grid-cols-2">
      {events.map((ev) => (
        <EventCard title={ev.title} cover={ev.cover} start_time={ev.start_time} venue={ev.venue} key={ev.title} />
      ))}
    </div>
  )
}

const Discover = () => {
  return (
    <div className="px-[60px]">
      <div className="flex justify-between">
        <div>
          <p className="text-white text-lg font-semibold">Popular Events</p>
          <p className="text-neutral-500">Nairobi</p>
        </div>


        <div>
          <div className="text-xs px-3 py-2 bg-neutral-950 rounded-md">View All</div>
        </div>
    {/* <button className="bg-neutral-900 text-xs px-4 py-1.5 rounded-md hover:bg-neutral-450 hover:text-neutral-850">View All</button> */}
      </div>

      <EventList />
      <hr className="text-neutral-800 mt-[24px] "/>

    </div>
  )
}

const CategoryCard = ({icon,name,count}) => {
  return (
    <div style={{backgroundColor: "#1c1e20"}} className="p-4 rounded-md">
          <p>Icon</p>
          <p className="font-semibold">{name}</p>
          <p className="text-sm text-neutral-500">{count} Events</p>
    </div>
  )
}

const BrowseByCategory = () => {
  const categories = [
    {
      icon: "Brain Icon",
      name: "AI",
      count: "1K",
    },
    {
      icon: "Paint",
      name: "Arts & Culture",
      count: "1K",
    },
    {
      icon: "Planet",
      name: "Climate",
      count: "936",
    },
    {
      icon: "Runner",
      name: "Fitness",
      count: "684",
    },
    {
      icon: "Wellness",
      name: "Wellness",
      count: "1K",
    },
    {
      icon: "Bitcoin",
      name: "Crypto",
      count: "1K",
    },
  ]
  return (
    <div className="px-[60px] pt-[24px] pb-[12px]">
      <p className="text-white font-semibold mb-5">Browse by Category</p>
      <div className="grid grid-cols-3 gap-4">
        {categories.map((category) => (
          <CategoryCard icon={category.icon} name={category.name} count={category.count} key={category.name}/>
        ))}
      </div>
      <hr className="text-neutral-800 mt-[24px] "/>
    </div>
  )
}

const CalendarCard = ({icon,title,venue,description}) => {
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

const FeaturedCalendars = () => {
  const calendars = [
    {
      icon: "",
      title: "Build Club",
      venue: "",
      description: "The best place in the world to learn AI",
    },
    {
      icon: "",
      title: "Design Buddies",
      venue: "",
      description: "Events for designers and all creatives across SF,online and physical",
    },
    {
      icon: "",
      title: "The GenAI Collective",
      venue: "",
      description: "The US's largets AI community: 25,0000+ founders,researchers and hackers",
    },
    {
      icon: "",
      title: "Her Workplace",
      venue: "",
      description: "The career network for the next generation of women and non binary leaders",
    },
    {
      icon: "",
      title: "South Park Commons",
      venue: "",
      description: "South Park Commons helps you get from -1 to 0. To learn more .....",
    },
  ]
  return (
    <div className="px-[60px] py-[16px]">
      <p className="text-white font-semibold mb-5">Featured Calendars</p>
      <div className="grid grid-cols-3 gap-4">
        {calendars.map((calendar) => (
          <CalendarCard title={calendar.title} description={calendar.description}/>
        ))}
      </div>
      <hr className="text-neutral-800 mt-[24px] "/>
    </div>
  )
}

const ExploreLocalEvents = () => {
  return (
    <div className="px-[60px] py-[16px]">
      <h3 className="text-white font-semibold mb-5">Explore Local Events</h3>
      <div className="flex">
        <div className="text-sm px-3 py-1.5 bg-neutral-950 rounded-md">Africa</div>
        <div className="text-sm px-3 py-1.5">Asia & Pacific</div>
        <div className="text-sm px-3 py-1.5">Europe</div>
        <div className="text-sm px-3 py-1.5">South America</div>
        <div className="text-sm px-3 py-1.5">North America</div>
      </div>
    </div>
  )
}

const App = () => {
  const [count,  setCount] = useState(0)

  return (
    <div>
      <Header />
      <div style={{backgroundColor: "#131517"}} className="text-neutral-300">
        <div className="hero-radial">
        </div>
        <Intro />
        <Discover />
        <BrowseByCategory />
        <FeaturedCalendars />
        <ExploreLocalEvents />
      </div>
    </div>
  )
}

export default App
