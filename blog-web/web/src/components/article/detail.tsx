import React from "react";
import {Typography} from "antd";
import {ArticleDetail as ArticleDetailInterface} from "../../model/article";
import {fetchArticleDetail} from "../../api/article";
import {RouteComponentProps} from 'react-router-dom'

// router transfer parameter
interface RouterProps {
  id: string
}

// by using RouteComponentProps receive Props from parent
interface ArticleDetailProps extends RouteComponentProps<RouterProps> {
  ad: ArticleDetailInterface
}

// current context state
interface ArticleDetailState {
  ad: ArticleDetailInterface
}

const {Title, Paragraph} = Typography;

class ArticleDetail extends React.Component<ArticleDetailProps, ArticleDetailState> {
  constructor(props: ArticleDetailProps) {
    super(props);
    this.state = {
      ad: this.props.ad || {}
    };
  }

  componentDidMount() {
    this.updateState()
  }

  updateState() {
    const id = this.props.match.params.id;
    fetchArticleDetail(id)
      .then(res => {
        this.setState({
          ad: res.data
        })
      })
  }

  render() {
    const {Topic, Content} = this.state.ad;
    return (
      <div>
        <Typography>
          <Title>{Topic}</Title>
          <Paragraph>{Content}</Paragraph>
        </Typography>
      </div>
    );
  }
}

export default ArticleDetail;
