import React, { useState } from "react";
import Form from "react-bootstrap/Form";
import Button from "react-bootstrap/Button";
import "../Styles/Login.css";
import Layout from "../components/Layout";
import UserClient from "../APIClients/UserClient"

export default function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    function validateForm() {
        return email.length > 0 && password.length > 0;
    }

    async function handleSubmit(event) {
        event.preventDefault();
        try{
            const resp = await UserClient.getBearerToken({email,password});
            if (resp.hasOwnProperty('token')){
                localStorage.setItem('token', `Bearer ${resp.token}`)
            }
            if (resp.hasOwnProperty('user_type')){
                localStorage.setItem('userType', resp.user_type);
            } else {
                console.log("Didn't receive a user type on authentication")
            }
            
        }catch(err){
            console.error(err)
        }
        
        // const resp = UserClient.getBearerToken({
        //     email,
        //     password
        // }).then((data) => {
        //     if (data.hasOwnProperty('token')){
        //         localStorage.setItem('token', `Bearer ${data.token}`)
        //     }
        // });
    }



    return (
       <Layout title={'Login Page'} description={'Please Login.'}>
           <div className="Login">
               <Form onSubmit={handleSubmit}>
                   <Form.Group size="lg" controlId="email">
                       <Form.Label>Email</Form.Label>
                       <Form.Control
                           autoFocus
                           type="email"
                           value={email}
                           onChange={(e) => setEmail(e.target.value)}
                       />
                   </Form.Group>
                   <Form.Group size="lg" controlId="password">
                       <Form.Label>Password</Form.Label>
                       <Form.Control
                           type="password"
                           value={password}
                           onChange={(e) => setPassword(e.target.value)}
                       />
                   </Form.Group>
                   <Button block size="lg" type="submit" disabled={!validateForm()}>
                       Login
                   </Button>
               </Form>
           </div>
       </Layout>
    );

    //if logged in, back to store page,
    //error handling page
    //pop up on successful login
}