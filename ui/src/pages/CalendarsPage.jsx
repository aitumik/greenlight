import { Layout } from '../components/Layout'
import {Intro} from '../components/Intro'

export const CalendarsPage = () => {
    return (
		<Layout>
         <div className="px-[310px] py-[60px]"> 
            <h1 className="text-2xl">Calendars</h1>
            <p className="text-white text-base py-[16px]">Explore popular events near you, browser by category, or check out some of the great community calendars</p>
         </div>
		</Layout>
    )
}
