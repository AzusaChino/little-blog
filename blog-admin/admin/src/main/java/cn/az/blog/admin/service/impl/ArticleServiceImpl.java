package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.entity.Article;
import cn.az.blog.admin.mapper.ArticleMapper;
import cn.az.blog.admin.service.IArticleService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * @author az
 * @version 2021-01-29
 */
@Service
public class ArticleServiceImpl extends ServiceImpl<ArticleMapper, Article> implements IArticleService {

}
