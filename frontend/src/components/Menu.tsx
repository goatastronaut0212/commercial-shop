function Menu(): JSX.Element {
  let language: string = 'vn'
  let page: JSX.Element = <></>;

  if (language === 'vn') {
    page = (
      <nav className='flex'>
        <a href='/'><img src='./public/asset/logo.jpg' alt='Logo' /></a>
        <ul className='flex list-none'>
          <li className='p-10 border-2'><a href='/' className='font-barlow text-5xl'>Trang chủ</a></li>
          <li className='p-10 border-2'><a href='/product' className='font-barlow text-5xl'>Sản phẩm</a></li>
          <li className='p-10 border-2'><a href='/about' className='font-barlow text-5xl'>Về chúng tôi</a></li>
          <li className='p-10 border-2'><a href='/news' className='font-barlow text-5xl'>Tin mới</a></li>
        </ul>
        <a href=''><img src='./public/asset/icon-sreach.jpg' alt='Logo' /></a>
        <a href=''><img src='./public/asset/icon-account.png' alt='Logo' /></a>
        <a href=''><img src='./public/asset/icon-cart.png' alt='Logo' /></a>
        <select name='language'>
          <option>VN 🇻🇳</option>
          <option>EN 🇺🇸</option>
        </select>
      </nav>
    )
  } else if (language === 'en') {
    page = (
      <nav className='flex'>
        <a href='/'><img src='./public/asseach.jpg' alt='Logo' /></a>
        <a href=""><img src='./public/asset/icon-acct/logo.jpg' alt='Logo' /></a>
        <ul className='flex list-none'>
          <li className='p-10 border-2'><a href='/' className='font-barlow text-5xl'>Home</a></li>
          <li className='p-10 border-2'><a href='/product' className='font-barlow text-5xl'>Product</a></li>
          <li className='p-10 border-2'><a href='/about' className='font-barlow text-5xl'>About</a></li>
          <li className='p-10 border-2'><a href='/news' className='font-barlow text-5xl'>News</a></li>
        </ul>
        <a href=""><img src='./public/asset/icon-sreount.png' alt='Logo' /></a>
        <a href=""><img src='./public/asset/icon-cart.png' alt='Logo' /></a>
        <select name='language'>
          <option>VN 🇻🇳</option>
          <option>EN 🇺🇸</option>
        </select>
      </nav>
    )
  }

  return page
}

export default Menu
