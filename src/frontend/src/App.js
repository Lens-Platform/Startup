import React from 'react';
import './App.css';
import {Route, Switch} from 'react-router-dom';
import SignInAndSignUpPage from './pages/sign-in-and-sign-up/sign-in-and-sign-up.component';
import Header from './components/header/header.component';
import HomePage from './pages/homepage/homepage.component';

class App extends React.Component {
  render() {
    return (
      <div>
          <Header />
          <Switch>
            <Route exact path='/'  component={HomePage} />
            <Route path='/signin'  component={SignInAndSignUpPage} />
          </Switch>
      
      </div>
    );
  }
}

export default App;
