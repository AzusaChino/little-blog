package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.entity.Blog;
import cn.az.blog.admin.mapper.BlogMapper;
import cn.az.blog.admin.service.IBlogService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * @author ycpang
 * @since 2021-01-19 10:49
 */
@Service
public class BlogServiceImpl extends ServiceImpl<BlogMapper, Blog> implements IBlogService {
}
