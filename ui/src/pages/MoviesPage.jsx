import { Layout } from '../components/Layout'
import {Intro} from '../components/Intro'

export const MoviesPage = () => {
  return (
		<Layout>
         <div className="px-[310px] py-[60px]"> 
            <h1 className="text-2xl">Movies</h1>
            <p className="text-white text-base py-[16px]">Explore popular movies near you, browser by category, or check out some of the great community</p>
         </div>
		</Layout>
  )
}
