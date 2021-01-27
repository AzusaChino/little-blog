import React from "react";
import "./App.css";
import {Layout, Menu} from "antd";
import {Router, Link, Switch, Route} from 'react-router'
import Article from "./components/article";
import {About} from './views/about'
import {Comment} from './views/comment'

const {Header, Content, Footer} = Layout;

function App() {
    return (
        <div className="App">
            <Layout className="layout">
                <Header>
                    <div className="logo"/>
                    <Router history="">
                        <Menu theme="dark" mode="horizontal">
                            <Menu.Item key="1">
                                <Link to="/">Home</Link>
                            </Menu.Item>
                            <Menu.Item key="2">
                                <Link to="/about">Home</Link>
                            </Menu.Item>
                            <Menu.Item key="3">
                                <Link to="/comment">Home</Link>
                            </Menu.Item>
                        </Menu>
                    </Router>
                </Header>
                <Content>
                    <Switch>
                        <Route path="/">
                            <Article/>
                        </Route>
                        <Route path="/about">
                            <About/>
                        </Route>
                        <Route path="/comment">
                            <Comment/>
                        </Route>
                    </Switch>
                </Content>
                <Footer>built by az</Footer>
            </Layout>
        </div>
    );
}

export default App;
