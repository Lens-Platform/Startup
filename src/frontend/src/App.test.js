import React from 'react';
import { shallow } from 'enzyme';
import App from './App';
import {Route, Switch} from 'react-router-dom';
import SignInAndSignUpPage from './pages/sign-in-and-sign-up/sign-in-and-sign-up.component';
import Header from './components/header/header.component';
import HomePage from './pages/homepage/homepage.component';

describe('App', () => {
  let wrapper;

  beforeEach(() => wrapper = shallow(<App />));

  it('should render a <div />', () => {
    expect(wrapper.find('div').length).toEqual(1);
  });

  it('should render the Header Component', () => {
    expect(wrapper.containsMatchingElement(< Header/>)).toEqual(true);
  });


});