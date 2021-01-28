export interface ArticleList {
  Id: string;
  Topic: string;
  Thumbnail: string;
  PublishState: number;
  PublishTime: Date;
  CreateUser: string;
  CreateTime: Date;
  UpdateUser: string;
  UpdateTime: Date;
  IsDelete: number;
}

export interface ArticleDetail extends ArticleList {
  Content: string
}
