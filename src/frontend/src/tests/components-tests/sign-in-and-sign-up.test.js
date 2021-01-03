import React from 'react';
import { shallow } from 'enzyme';
import SignInAndSignUpPage from '../../pages/sign-in-and-sign-up/sign-in-and-sign-up.component';
import SignUp from '../../components/sign-up/sign-up.component.jsx';
import SignIn from '../../components/sign-in/sign-in.component';


describe('SignInAndSignUpPage', () => {
    let wrapper;
  
    beforeEach(() => wrapper = shallow(<SignInAndSignUpPage />));
  
    it('should render a <div />', () => {
      expect(wrapper.find('div').length).toEqual(1);
    });
  
    it('should render the SignIn Component', () => {
      expect(wrapper.containsMatchingElement(< SignIn/>)).toEqual(true);
    });

    it('should render the SignUp Component', () => {
        expect(wrapper.containsMatchingElement(< SignUp/>)).toEqual(true);
      });
  });