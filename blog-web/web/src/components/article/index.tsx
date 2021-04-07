import React, {CSSProperties} from "react";
import {Carousel, List, message} from "antd";
import {fetchList} from "../../api/article";
import {Link} from 'react-router-dom'
import {ArticleList as ArticleListDefinition} from "../../model/article";

interface IState {
  pageNum: number,
  pageSize: number,
  data: Array<ArticleListDefinition>;
}

const size = 10;

/***************style section****************/

const carouselStyle: CSSProperties = {
  height: '160px',
  color: '#fff',
  lineHeight: '160px',
  textAlign: 'center',
  background: '#364d79'
}

const articleContainer: CSSProperties = {
  margin: "auto",
  width: "600px",
  background: 'fff'
}

const defaultState: IState = {
  pageNum: 1,
  pageSize: 10,
  data: []
}

class ArticleList extends React.Component<{}, IState> {

  state = defaultState;

  componentDidMount() {
    this.doFetch(false)
  }

  /**
   * fetch data function
   * @param append
   */
  doFetch(append: boolean) {
    let close = function () {
    }
    message.loading("加载中...", onclose = close)
    const {pageNum} = this.state;
    fetchList({
      pageNum: pageNum,
      pageSize: size
    })
      .then((res) => {
        const {data} = res;
        this.setState({
          data: append ? this.state.data.concat(data.Data) : data.Data
        })
      })
      .catch(e => message.error("发生了错误" + e))
      .finally(() => close())
  }

  scrollToTop() {
    // reset pageNum
    this.setState({
      pageNum: 1
    })
  }

  // scroll down fetch more
  scrollToBottom() {
    this.setState({
      pageNum: this.state.pageNum + 1
    })
    this.doFetch(true)
  }

  render() {
    const {data} = this.state;
    return (
      <div className="mainContainer">
        <div className='carouselContainer'>
          <Carousel>
            <div>
              <h3 style={carouselStyle}>1</h3>
            </div>
            <div>
              <h3 style={carouselStyle}>2</h3>
            </div>
          </Carousel>
        </div>
        <div style={articleContainer}>
          <List
            itemLayout="horizontal"
            dataSource={data}
            renderItem={(item) => (
              <List.Item>
                <List.Item.Meta
                  title={<Link to={`/article/${item.Id}`}>{item.Topic}</Link>}
                />
              </List.Item>
            )}
          />
        </div>
      </div>
    );
  }
}

export default ArticleList;
