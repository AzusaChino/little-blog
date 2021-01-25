import React from "react";
import "./App.css";
import { Layout, Menu } from "antd";
import {Router} from 'react-router'
import Article from "./components/article";

const { Header, Content, Footer } = Layout;
function App() {
  return (
    <div className="App">
      <Layout className="layout">
        <Header>
          <div className="logo" />
          <Menu theme="dark" mode="horizontal">
            <Menu.Item key="1">Nav 1</Menu.Item>
            <Menu.Item key="2">Nav 2</Menu.Item>
          </Menu>
        </Header>
        <Content>
          <Article />
        </Content>
        <Footer>built by az</Footer>
      </Layout>
    </div>
  );
}

export default App;
