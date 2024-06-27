
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.scss'
import Footer from './components/footer/footer';
import { Home } from './pages/home/home';
import { Navbar } from './components/navbar/navbar';
import { Movie } from './pages/movie/movie';
import { Movie_theater } from './pages/movie/theaters/movie_theater';
import Auth from './pages/auth/auth';


function App() {
  return (
    <>
      <BrowserRouter>
        <div className='main'>
          <Navbar/>
          <div className='content'>
            <Routes>
              <Route path='/' element={<Home />} />
              <Route path='/movie' element={<Movie/>}/>
              <Route path='/movie/:id' element={<Movie_theater/>}/>
              <Route path='/login' element={<Auth state='login'/>}/> 
              <Route path='/signup' element={<Auth state='register'/>}/>
            </Routes> 
          </div>
          <div className='Footerr'>
            <Footer />
          </div>
        </div>



      </BrowserRouter>
    </>
  )
}

export default App
