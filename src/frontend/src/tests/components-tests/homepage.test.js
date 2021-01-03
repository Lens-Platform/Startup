import React from 'react';
import { shallow } from 'enzyme';
import HomePage from '../../pages/homepage/homepage.component';

describe('HomePage', () => {
  it('should render a <div />', () => {
    const wrapper = shallow(<HomePage />);
    expect(wrapper.find('div').length).toEqual(1);
  });
});