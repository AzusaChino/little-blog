package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.entity.Permission;
import cn.az.blog.admin.entity.User;
import cn.az.blog.admin.mapper.UserMapper;
import cn.az.blog.admin.service.IUserService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @author az
 * @version 2021-01-29
 */
@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements IUserService {

    @Override
    public User getUserByUsername(String username) {
        return null;
    }

    @Override
    public User register(User user) {
        return null;
    }

    @Override
    public String login(String username, String password) {
        return null;
    }

    @Override
    public List<Permission> getPermissionList(String userId) {
        return null;
    }
}
