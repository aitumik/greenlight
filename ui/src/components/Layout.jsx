import { Header } from "./Header"

export const Layout = ({children}) => {
    return (
        <div style={{backgroundColor: "#131517"}} className="text-neutral-300 h-screen">
            <Header />
            <div className="hero-radial">
            </div>
            {children}
        </div>
    )
}