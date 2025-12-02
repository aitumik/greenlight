import { Layout } from "../components/Layout";
import { EventCard } from "../components/EventCard";
import { Intro } from "../components/Intro";
import { CalendarCard } from "../components/CalendarCard";
import { useState } from "react";

const EventList = ({ events, onBuyTicket }) => {
  return (
    <div className="grid grid-cols-1 sm:grid-cols-2 gap-4">
      {events.map((ev) => (
        <div key={ev.title} className="relative">
          <EventCard
            title={ev.title}
            cover={ev.cover}
            start_time={ev.start_time}
            venue={ev.venue}
          />
          <div className="mt-2 flex flex-col gap-1">
            {ev.soldOut ? (
              <span className="text-red-500 text-sm font-semibold">
                Sold Out
              </span>
            ) : (
              <button
                className="bg-green-700 text-white px-4 py-2 rounded text-lg"
                onClick={() => onBuyTicket(ev)}
              >
                Buy Ticket{ev.price ? ` - ${ev.price}` : ""}
              </button>
            )}
            {ev.ticketsAvailable !== undefined && (
              <span className="text-neutral-400 text-sm">
                {ev.ticketsAvailable} tickets left
              </span>
            )}
          </div>
        </div>
      ))}
    </div>
  );
};

const Discover = () => {
  // todo : the first thing is fetch this from the backend or atleast match the state
  // in the db
  const allEvents = [
    {
      title: "Health Sector AI Roundtable 4.0",
      cover:
        "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=0,background=white,quality=75,width=300,height=300/event-covers/jg/a3c30fe6-55ea-4835-984c-127d1049b4ca.png",
      start_time: "Tue,Apr 8, 9:00 AM",
      venue: "Marsabit Plaza",
      price: "$20",
      ticketsAvailable: 12,
      soldOut: false,
    },
    {
      title: "Heels & Green Speed Connect",
      cover:
        "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/no/affec154-182c-4bce-8b49-974396bf503f.png",
      start_time: "Thu,Apr 10, 6:30 PM",
      venue: "iHUB",
      price: "$15",
      ticketsAvailable: 5,
      soldOut: false,
    },
    {
      title: "Launch of Mercury Contamination in Migori Research Study",
      cover:
        "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/6q/f059c56f-a1fe-46e5-b6a5-6c50b8e650a9.png",
      start_time: "Fri,Apr 11, 10:00 AM",
      venue: "Sarova Stanley Nairobi",
      price: "$10",
      ticketsAvailable: 0,
      soldOut: true,
    },
    {
      title: "LW14 - Nairobi, Kenya - Supabase Meetup",
      cover:
        "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/gallery-images/j5/e0d9e03b-63b3-452b-8899-3fc8f848a4bb",
      start_time: "Sat,Apr 12, 1:00 PM",
      venue: "Nairobi Garage",
      price: "$25",
      ticketsAvailable: 8,
      soldOut: false,
    },
    {
      title: "Sui Kenya Community MeetUp",
      cover:
        "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/45/7c45fc3a-5887-457f-8bf3-94be18f25995.png",
      start_time: "Sat,Apr 12, 2:00 PM",
      venue: "Nafasi Connection LLC",
      price: "$30",
      ticketsAvailable: 20,
      soldOut: false,
    },
    {
      title: "The Wake Live Recording",
      cover:
        "https://images.lumacdn.com/cdn-cgi/image/format=auto,fit=cover,dpr=1,background=white,quality=75,width=300,height=300/event-covers/0q/d5ffeaf8-0858-495f-885c-f4894de047f0.jpg",
      start_time: "Sat,Apr 19, 12:00 PM",
      venue: "Kenya Cultural Centre",
      price: "$50",
      ticketsAvailable: 0,
      soldOut: true,
    },
  ];

  const [search, setSearch] = useState("");
  const [sort, setSort] = useState("date");
  const [category, setCategory] = useState("");
  const [showAll, setShowAll] = useState(false);
  const [visibleCount, setVisibleCount] = useState(4);
  const [showTicketModal, setShowTicketModal] = useState(false);
  const [selectedEvent, setSelectedEvent] = useState(null);

  let filteredEvents = allEvents.filter(
    (ev) =>
      ev.title.toLowerCase().includes(search.toLowerCase()) &&
      (category
        ? ev.title.toLowerCase().includes(category.toLowerCase())
        : true),
  );
  if (sort === "date") {
    filteredEvents = filteredEvents.sort(
      (a, b) => new Date(a.start_time) - new Date(b.start_time),
    );
  }

  const eventsToShow = showAll
    ? filteredEvents
    : filteredEvents.slice(0, visibleCount);

  const handleBuyTicket = (event) => {
    setSelectedEvent(event);
    setShowTicketModal(true);
  };

  return (
    <div className="px-[60px]">
      <div className="flex justify-between items-center mb-4">
        <div>
          <p className="text-white text-lg font-semibold">Popular Events</p>
          <p className="text-neutral-500">Nairobi</p>
        </div>
        <div className="flex gap-2 items-center">
          <input
            type="text"
            placeholder="Search events..."
            value={search}
            onChange={(e) => setSearch(e.target.value)}
            className="px-2 py-1 rounded bg-neutral-900 text-white text-sm"
            style={{ minWidth: 180 }}
          />
          <select
            value={sort}
            onChange={(e) => setSort(e.target.value)}
            className="px-2 py-1 rounded bg-neutral-900 text-white text-xs"
          >
            <option value="date">Sort by Date</option>
            <option value="title">Sort by Title</option>
          </select>
          <select
            value={category}
            onChange={(e) => setCategory(e.target.value)}
            className="px-2 py-1 rounded bg-neutral-900 text-white text-xs"
          >
            <option value="">All Categories</option>
            <option value="ai">AI</option>
            <option value="arts">Arts & Culture</option>
            <option value="climate">Climate</option>
            <option value="fitness">Fitness</option>
            <option value="wellness">Wellness</option>
            <option value="crypto">Crypto</option>
          </select>
        </div>
        <div>
          <button
            className="text-xs px-3 py-2 bg-neutral-950 rounded-md"
            onClick={() => setShowAll(true)}
          >
            View All
          </button>
        </div>
      </div>
      <div className="mb-2 text-neutral-400 text-xs">
        Showing {eventsToShow.length} of {filteredEvents.length} events
      </div>
      <EventList events={eventsToShow} onBuyTicket={handleBuyTicket} />
      {!showAll && filteredEvents.length > visibleCount && (
        <div className="flex justify-center mt-4">
          <button
            className="px-4 py-2 bg-neutral-900 rounded text-white"
            onClick={() => setVisibleCount(visibleCount + 4)}
          >
            Load More
          </button>
        </div>
      )}
      {showTicketModal && selectedEvent && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white p-6 rounded shadow-lg min-w-[300px]">
            <h2 className="text-lg font-bold mb-2">Buy Ticket</h2>
            <p className="mb-2">
              Event:{" "}
              <span className="font-semibold">{selectedEvent.title}</span>
            </p>
            <p className="mb-2">Venue: {selectedEvent.venue}</p>
            <p className="mb-2">Date: {selectedEvent.start_time}</p>
            <p className="mb-2">Price: {selectedEvent.price || "Free"}</p>
            <button
              className="bg-green-700 text-white px-4 py-2 rounded"
              onClick={() => {
                setShowTicketModal(false);
                alert("Ticket purchased!");
              }}
            >
              Confirm Purchase
            </button>
            <button
              className="ml-2 px-4 py-2"
              onClick={() => setShowTicketModal(false)}
            >
              Cancel
            </button>
          </div>
        </div>
      )}
      <hr className="text-neutral-800 mt-[24px] " />
    </div>
  );
};

const CategoryCard = ({ icon, name, count }) => {
  return (
    <div style={{ backgroundColor: "#1c1e20" }} className="p-4 rounded-md">
      <p>Icon</p>
      <p className="font-semibold">{name}</p>
      <p className="text-sm text-neutral-500">{count} Events</p>
    </div>
  );
};

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
  ];
  return (
    <div className="px-[60px] pt-[24px] pb-[12px]">
      <p className="text-white font-semibold mb-5">Browse by Category</p>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {categories.map((category) => (
          <CategoryCard
            icon={category.icon}
            name={category.name}
            count={category.count}
            key={category.name}
          />
        ))}
      </div>
      <hr className="text-neutral-800 mt-[24px] " />
    </div>
  );
};

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
      description:
        "Events for designers and all creatives across SF,online and physical",
    },
    {
      icon: "",
      title: "The GenAI Collective",
      venue: "",
      description:
        "The US's largets AI community: 25,0000+ founders,researchers and hackers",
    },
    {
      icon: "",
      title: "Her Workplace",
      venue: "",
      description:
        "The career network for the next generation of women and non binary leaders",
    },
    {
      icon: "",
      title: "South Park Commons",
      venue: "",
      description:
        "South Park Commons helps you get from -1 to 0. To learn more .....",
    },
  ];
  return (
    <div className="px-[60px] py-[16px]">
      <p className="text-white font-semibold mb-5">Featured Calendars</p>
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 gap-4">
        {calendars.map((calendar) => (
          <CalendarCard
            title={calendar.title}
            description={calendar.description}
          />
        ))}
      </div>
      <hr className="text-neutral-800 mt-[24px] " />
    </div>
  );
};

const ExploreLocalEvents = () => {
  return (
    <div className="px-[60px] py-[16px]">
      <h3 className="text-white font-semibold mb-5">Explore Local Events</h3>
      <div className="flex flex-wrap gap-2">
        <div className="text-sm px-3 py-1.5 bg-neutral-950 rounded-md">
          Africa
        </div>
        <div className="text-sm px-3 py-1.5">Asia & Pacific</div>
        <div className="text-sm px-3 py-1.5">Europe</div>
        <div className="text-sm px-3 py-1.5">South America</div>
        <div className="text-sm px-3 py-1.5">North America</div>
      </div>
    </div>
  );
};

export const HomePage = () => {
  return (
    <Layout>
      <div className="w-full min-h-screen" style={{ fontSize: "1.15rem" }}>
        <div className="mx-auto w-full px-4 sm:px-8 md:px-16 lg:px-[250px]">
          <Intro />
          <Discover />
          <BrowseByCategory />
          <FeaturedCalendars />
          <ExploreLocalEvents />
        </div>
      </div>
    </Layout>
  );
};
