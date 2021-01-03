import React from 'react';
import './search-bar.styles.scss';

export  const SearchBar = () => (
    // <input  
    //     className='search'
    //     type='search' 
    //     placeholder='Type here ...'
    //     autocomplete="on" 
    // />

    <form className='search-form'>
        <div className='business-search'>
            <div className='search-element'>
                <label className='search-label'> Type here</label>
                <input 
                    className='search-input'
                    type='text'
                    autocomplete='on'
                    placeholder='Company or Product'
                />
            </div>
            <div className='search-element'>
                <label className='search-label'> Location</label>
                <input 
                    className='search-input'
                    type='text'
                    autocomplete='on'
                    placeholder='City, State'
                />
            </div>
            <a type="submit" class="search-button">Search</a>
        </div>
    </form>
)

export default SearchBar;