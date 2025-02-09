import { BrowserRouter, Route, Routes } from 'react-router'
import './App.css'
import Staging from './routes/Staging'
import Album from './routes/Album'
import Albums from './routes/Albums'
import Trash from './routes/Trash'
import Settings from './routes/Settings'
import Home from './routes/Home'

export const appName = "fani"

export default function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<Home />} />
        <Route path="staging" element={<Staging />} />
        <Route path="albums" element={<Album />} />
        <Route path="albums/:id" element={<Albums />} />
        <Route path="trash" element={<Trash />} />
        <Route path="settings" element={<Settings />} />
      </Routes>
    </BrowserRouter>
  )
}
