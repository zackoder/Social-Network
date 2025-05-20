"use client"
import { usePathname } from "next/navigation"
import Navbar from "./navbar"

export default function ConditionalNavbar(){
    const pathname = usePathname();
    const isAuthPage = pathname === "/login" || pathname === "/sign-up";
    if (isAuthPage) return null;
    return <Navbar />
}