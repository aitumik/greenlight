import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { ExplorePage } from "./pages/ExplorePage";
import { CalendarsPage } from "./pages/CalendarsPage";
import { MoviesPage } from "./pages/MoviesPage";
import { EventsPage } from "./pages/EventsPage";


const App = () => {

  return (
    <BrowserRouter>
        <Routes>
          <Route path="/" element={<ExplorePage />} />
          <Route path="/explore" element={<ExplorePage />} />
          <Route path="/calendars" element={<CalendarsPage />} />
          <Route path="/movies" element={<MoviesPage />} />
          <Route path="/events" element={<EventsPage />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
    </BrowserRouter>
  )
}

export default App
