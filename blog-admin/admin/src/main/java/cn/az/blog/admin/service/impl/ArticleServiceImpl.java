package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.entity.Article;
import cn.az.blog.admin.mapper.ArticleMapper;
import cn.az.blog.admin.service.IArticleService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

/**
 * <p>
 * 博客文章表 服务实现类
 * </p>
 *
 * @author az
 * @since 2021-01-25
 */
@Service
@Transactional(rollbackFor = Exception.class)
public class ArticleServiceImpl extends ServiceImpl<ArticleMapper, Article> implements IArticleService {

}
