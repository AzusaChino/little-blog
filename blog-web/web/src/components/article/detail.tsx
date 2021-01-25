import React from "react";
import { Typography } from "antd";
import { ArticleDetail as ArticleDetailInterface } from "../../model/article";

interface IProps {
  ad: ArticleDetailInterface;
}

const { Title, Paragraph } = Typography;

class ArticleDetail extends React.Component<
  IProps,
  { ad: ArticleDetailInterface },
  {}
> {
  constructor(props: IProps) {
    super(props);
    this.state = {
      ad: this.props.ad || {},
    };
  }

  render() {
    const { Topic, Content } = this.state.ad;
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
