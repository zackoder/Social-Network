"use client"
import { createContext, useState } from "react";


export const DataContext = createContext()

export const DataProvider = ({children}) => {
    const [selectedContactsIds, setSelectedContactsIds] = useState([]);
    return (
        <DataContext.Provider value={{selectedContactsIds, setSelectedContactsIds}}>
            {children}
        </DataContext.Provider>
    )
} 

