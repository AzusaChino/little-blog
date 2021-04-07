import React, {CSSProperties} from "react";
import "./App.css";
import {Layout, Menu} from "antd";
import {BrowserRouter as Router, Link, Switch, Route} from 'react-router-dom'
import Article from "./components/article";
import ArticleDetail from "./components/article/detail";
import About from './views/about'
import Comment from './views/comment'

const {Header, Content, Footer} = Layout;

const contentStyle: CSSProperties = {
  margin: '24px 16px 0',
  overflow: 'initial'
}

function App() {
  return (
    <div className="App">
      <Router>
        <Layout className="layout">
          <Header>
            <div className="logo"/>
            <Menu theme="dark" mode="horizontal">
              <Menu.Item key="1">
                <Link to="/">Home</Link>
              </Menu.Item>
              <Menu.Item key="2">
                <Link to="/about">About</Link>
              </Menu.Item>
              <Menu.Item key="3">
                <Link to="/comment">Comment</Link>
              </Menu.Item>
            </Menu>
          </Header>
          <Content style={contentStyle}>
            <Switch>
              {/* 使用exact进行精确匹配 */}
              <Route exact path="/" component={Article}/>
              <Route path="/about" component={About}/>
              <Route path="/comment" component={Comment}/>
              <Route path="/article/:id" component={ArticleDetail}/>
              <Route render={() => <h1>Page Not Found</h1>}/>
            </Switch>
          </Content>
          <Footer>built by az</Footer>
        </Layout>
      </Router>
    </div>
  );
}

export default App;
