import { useState } from 'react'

import Menu from '../components/Menu'

function Home(): JSX.Element {
  const [language, setLanguage] = useState<string>('vn')

  return (
    <div>
      <Menu />
      <div>Home Page</div>
    </div>
  )
}

export default Home