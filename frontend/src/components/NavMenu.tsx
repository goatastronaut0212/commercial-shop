import React from 'react'
<link href="/dist/output.css" rel="stylesheet" />

function NavMenu() {
  return (
  <>
     <div> 
        <a href=""><img src="/asset/logo.jpg" alt="Logo" /></a>  
        <ul className='list-none'>
            <li>Home</li>
            <li>Products</li>
            <li>About</li>
            <li>New</li>
        </ul>           
        <a href=""><img src="/asset/icon-sreach.jpg" alt="Logo" /></a>  
        <a href=""><img src="/asset/icon-account.png" alt="Logo" /></a>  
        <a href=""><img src="/asset/icon-cart.png" alt="Logo" /></a>  
     </div>
  </>      
  )
}

export default NavMenu