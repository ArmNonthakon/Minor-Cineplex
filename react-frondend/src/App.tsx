
import './App.scss'
import Footer from './components/footer/footer';
import { Navbar } from './components/navbar/navbar';
import { Movie_theater } from './pages/movie/theaters/movie_theater';
import { createBrowserRouter,RouterProvider } from 'react-router-dom';
import { Movie } from './pages/movie/movie';
import { useEffect } from 'react';
import { Home } from './pages/home/home';
const router = createBrowserRouter([
  {
    path: '/',
    element: (
      <>
        <Navbar />
        <Home/>
      </>
    )
  },
  {
    path: '/movie',
    element: (
      <>
        <Navbar />
        <Movie/>
      </>
    )
  },
  {
    path: '/seat',
    element: (
      <>
        <Navbar />
        <Movie_theater title='how to train your dragon 2'/>
      </>
    )
  }
  ,])
function App() {
  useEffect(()=>{
    console.log(router.window?.location)
  })
  return (
    <>
      <div className="main">
        <div className='content'>
          <RouterProvider router={router} />
        </div>
        <div className="Footerr">
          <Footer></Footer>
        </div>


      </div>

    </>
  )
}

export default App
