import React from "react";
import {Typography} from "antd";
import {ArticleDetail as ArticleDetailInterface} from "../../model/article";
import {fetchArticleDetail} from "../../api/article";
import {withRouter, RouteComponentProps} from 'react-router-dom'

interface RouterProps {
  id: string
}

interface ArticleDetailProps extends RouteComponentProps<RouterProps> {
  ad: ArticleDetailInterface
}

const {Title, Paragraph} = Typography;

class ArticleDetail extends React.Component<ArticleDetailProps,
  { ad: ArticleDetailInterface },
  {}> {
  constructor(props: ArticleDetailProps) {
    super(props);
    this.state = {
      ad: this.props.ad || {}
    };
  }

  componentDidMount() {
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

export default withRouter(ArticleDetail);
