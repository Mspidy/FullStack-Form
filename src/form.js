import React, {Component} from "react";
import axios from 'axios';

class Form extends Component{
    constructor(props){
        super(props)
        this.state={
            firstName:"",
            lastName:"",
            password:"",
            gender:"",
        }
        this.handleSubmit=this.handleSubmit.bind(this)
    }
    firsthandler = (event)=>{
        this.setState({
            firstName: event.target.value
        })
    }
    lasthandler = (event)=>{
        this.setState({
            lastName: event.target.value
        })
    }
    passwordhandler = (event)=>{
        this.setState({
            password: event.target.value
        })
    }
    genderhandler = (event)=>{
        this.setState({
            gender: event.target.value
        })
    }

    handleSubmit= (event)=>{
        console.log("hello")
        alert(`${this.state.firstName} ${this.state.lastName} Registered Successfully !!!!`)
        console.log(this.state);
        this.setState({
            firstName:"",
            lastName:"",
            password:"",
            gender:"",
        })
    event.preventDefault()
    // let temp=this.setState
    console.log("temppppp",this.setState)
    let arr ={
            FirstName:this.state.firstName,
            LastName:this.state.lastName,
            Password:this.state.password,
            Gender:this.state.gender,
    }
    console.log(arr)
    axios.post('http://localhost:8000/person',arr)
        .then(response =>{
            console.log(response)
        })
        .catch(error =>{
            console.log(error)
        })
    }

    render(){
        return(
            <div className="container">
                <form onSubmit={this.handleSubmit}>
                    <h1>User Registration</h1>
                    <label>FirstName:</label> <input type="text" value={this.state.firstName} onChange={this.firsthandler}/>
                    <br/>
                    <br/>
                    <br/>
                    <label>LastName:</label> <input type="text" value={this.state.lastName} onChange={this.lasthandler}/>
                    <br/>
                    <br/>
                    <br/>
                    <label>Password:</label> <input type="password" value={this.state.password} onChange={this.passwordhandler}/>
                    <br/>
                    <br/>
                    <br/>
                    <label>Gender:</label> <select onChange={this.genderhandler} defaultValue="Select Gender">
                        <option defaultValue>Select Gender</option>
                        <option value="male">Male</option>
                        <option value="female">Female</option>
                    </select><br/><br/><br/>
                    <input type="submit" value="Submit"/>
                </form>
            </div>
        )
    }
}

export default Form;