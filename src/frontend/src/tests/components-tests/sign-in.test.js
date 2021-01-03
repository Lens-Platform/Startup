import React from 'react';
import { shallow } from 'enzyme';
import SignIn from '../../components/sign-in/sign-in.component';

describe('SignIn', () => {
  it('should render a <div />', () => {
    const wrapper = shallow(<SignIn />);
    expect(wrapper.find('div').length).toEqual(2);
  });
});