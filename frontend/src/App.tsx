import './App.css'
// import * as dotenv from "dotenv"
import Map from './components/Map'
import { ChakraProvider } from '@chakra-ui/react'
import {
  BrowserRouter as Router,
  Routes,
  Link,
  Route,
} from "react-router-dom";
import UploadForm from './components/UploadForm';

function App() {
  // dotenv.config()
  return (
      <ChakraProvider>
        <Router>
          <Routes>
            <Route path="/" element={<Map />} />
            <Route path='/upload' element={<UploadForm/>}/>
          </Routes>
        </Router>
      </ChakraProvider>
  )
}

export default App
