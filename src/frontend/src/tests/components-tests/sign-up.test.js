import React from 'react';
import { shallow } from 'enzyme';
import SignUp from '../../components/sign-up/sign-up.component';

describe('SignUp', () => {
  it('should render a <div />', () => {
    const wrapper = shallow(<SignUp />);
    expect(wrapper.find('div').length).toEqual(1);
  });
});