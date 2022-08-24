import { useEffect, useState } from 'react'
import { Shortener, UrlTable } from '../components'


const ManagePage = () => {

  useEffect(() => {
  }, [])

  return (
    <div>
        <h1>It's manage page!</h1>
        <Shortener />
        <UrlTable />
    </div>
  )
}

export default ManagePage
