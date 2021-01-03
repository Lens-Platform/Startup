import React from 'react';
import {SearchBar} from '../../components/search-bar/search-bar.component.jsx';
import {ProductList} from '../../components/product-list/product-list.component';
import FooterPage from '../../pages/footerpage/footerpage.component';

import './homepage.styles.scss';

const HomePage = () => (
    //This is the directory for blackspace
    <div className='homepage'> 
    {/* this section has a list of 4 categories of businesses */}
        {/* <div className='categories'>add popular categories here</div> */}
        <SearchBar/>
        <ProductList/>
        {/* <div className='other options'> space for more product seelctions </div> */}
        <FooterPage/>
    </div>
);

export default HomePage;