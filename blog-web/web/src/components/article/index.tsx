import React from "react";
import {List} from "antd";
import {fetchList} from "../../api/article";
import {Link} from 'react-router-dom'
import {ArticleList as ArticleListDefinition} from "../../model/article";

interface IProps {
  size?: 10;
  data?: ArticleListDefinition[];
}

interface IState {
  data: ArticleListDefinition[];
}

class ArticleList extends React.Component<IProps, IState, {}> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      data: this.props.data || [],
    };
  }

  componentDidMount() {
    fetchList({})
      .then((res) => {
        this.setState({data: res.data});
      })
      .catch((e) => alert("发生了错误" + e));
  }

  render() {
    return (
      <div className="articleContainer">
        <List
          itemLayout="horizontal"
          dataSource={this.state.data}
          renderItem={(item) => (
            <div>
              <h2><Link to={`/article/${item.Id}`}>{item.Topic}</Link></h2>
            </div>
          )}
        />
      </div>
    );
  }
}

export default ArticleList;
