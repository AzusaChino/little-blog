package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.entity.Comment;
import cn.az.blog.admin.mapper.CommentMapper;
import cn.az.blog.admin.service.ICommentService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * @author az
 * @version 2021-01-29
 */
@Service
public class CommentServiceImpl extends ServiceImpl<CommentMapper, Comment> implements ICommentService {

}
