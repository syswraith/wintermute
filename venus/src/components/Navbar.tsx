import * as React from "react"
import { Link, useNavigate, useLocation } from "react-router-dom"
import { cn } from "@/lib/utils"

import {
  NavigationMenu,
  NavigationMenuItem,
  NavigationMenuLink,
  NavigationMenuList,
  navigationMenuTriggerStyle,
} from "@/components/ui/navigation-menu"
import { Button } from "@/components/ui/button"

export function Navbar() {
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = () => {
    localStorage.removeItem("token");
    navigate("/auth");
  };

  return (
    <div className="flex w-full justify-center p-4 border-b">
      <div className="flex w-full max-w-4xl justify-between items-center">
        <NavigationMenu>
          <NavigationMenuList>
            <NavigationMenuItem>
              <NavigationMenuLink asChild className={cn(navigationMenuTriggerStyle(), location.pathname === "/create" && "bg-white text-black hover:bg-white/90 focus:bg-white focus:text-black data-[active]:bg-white data-[active]:text-black")}>
                <Link to="/create">Create Link</Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
            
            <NavigationMenuItem>
              <NavigationMenuLink asChild className={cn(navigationMenuTriggerStyle(), location.pathname === "/dashboard" && "bg-white text-black hover:bg-white/90 focus:bg-white focus:text-black data-[active]:bg-white data-[active]:text-black")}>
                <Link to="/dashboard">Dashboard</Link>
              </NavigationMenuLink>
            </NavigationMenuItem>
          </NavigationMenuList>
        </NavigationMenu>

        <Button variant="ghost" onClick={handleLogout}>
          Logout
        </Button>
      </div>
    </div>
  )
}
