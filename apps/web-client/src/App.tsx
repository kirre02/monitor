import SiteList from './components/SiteList'

function App() {

  return (
    <>
      <div className='min-h-full container px-4 mx-auto my-16'>
        <h2 className='text-2xl font-bold leading-7 text-gray-900 sm:text-3xl sm:tracking-tight'>
          monitor
        </h2>

          <main className='pt-8 pb-16'>
            <SiteList />
        </main>
      </div>
    </>
  )
}

export default App
