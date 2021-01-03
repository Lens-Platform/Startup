import React from 'react';
import FormInput from '../form-input/form-input.component';
import CustomButton from '../custom-button/custom-button.component';


import './sign-up.styles.scss';

const validEmailRegex = RegExp(/^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i);


class SignUp extends React.Component{
    constructor(){
        super();

        this.state = {
            displayName:'',
            email:'',
            password:'',
            confirmPassword: '',
            errors: {
                displayName: '',
                email: '',
                password: '',
                confirmPassword:''
              }
        };
    }
    

    handleSubmit = async event => {
        event.preventDefault();

        const {displayName,email,password,confirmPassword } = this.state;

        if(password !== confirmPassword){
            alert("password don't match")
            return;
        }

        try {
            //add more here
            this.setState({
                displayName:'',
                email:'',
                password:'',
                confirmPassword:''
            });
        } catch (error) {
            console.error(error);
        }
    };

    handleChange = (event) => {
        event.preventDefault();
        const { name, value } = event.target;
        let errors = this.state.errors;
    
        switch (name) {
          case 'displayName': 
            errors.displayName = 
              value.length < 5
                ? 'Display must be at least 5 characters long!'
                : '';
            break;
          case 'email': 
            errors.email = 
              validEmailRegex.test(value)
                ? ''
                : 'Email is not valid!';
            break;
          case 'password': 
            errors.password = 
              value.length >= 8
                ? 'Password must be at least 8 characters long!'
                : '';
            break;
           case 'confirmPassword':
            errors.confirmPassword =
               value.length >= 8 
                ? 'Password must be at least 8 characters long!'
                : '';
            break;
          default:
            break;
        }
    
        this.setState({errors, [name]: value});
    }

    render(){
        const {displayName,email,password,confirmPassword, errors } = this.state;
       return(
        <div className='sign-up'>
            <h2 className='title'>I do not have an account</h2>
            <span>Sign up with your email and password</span>

            <form className='sign-up-form' onSubmit={this.handleSubmit}>
                <FormInput
                    type='text'
                    name='displayName'
                    value={displayName}
                    onChange = {this.handleChange}
                    label='Display Name'
                    required
                />
                <FormInput
                    type='email'
                    name='email'
                    value={email}
                    onChange = {this.handleChange}
                    label='Email'
                    required
                />
                <FormInput
                    type='password'
                    name='password'
                    value={password}
                    onChange = {this.handleChange}
                    label='Password'
                    required
                />
                <FormInput
                    type='password'
                    name='confirmPassword'
                    value={confirmPassword}
                    onChange = {this.handleChange}
                    label='Confirm Password'
                    required
                />
                <CustomButton type='submit' >SIGN UP </CustomButton>
            </form>
        </div>
       )
    }
}

export default SignUp;