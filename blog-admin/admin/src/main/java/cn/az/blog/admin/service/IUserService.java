package cn.az.blog.admin.service;

import cn.az.blog.admin.entity.Permission;
import cn.az.blog.admin.entity.User;
import com.baomidou.mybatisplus.extension.service.IService;

import java.util.List;

/**
 * @author az
 * @version 2021-01-29
 */
public interface IUserService extends IService<User> {

    /**
     * 根据用户名获取后台管理员
     *
     * @param username the username
     * @return user
     */
    User getUserByUsername(String username);

    /**
     * 注册功能
     *
     * @param user user
     * @return user
     */
    User register(User user);

    /**
     * 登录功能
     *
     * @param username 用户名
     * @param password 密码
     * @return 生成的JWT的token string
     */
    String login(String username, String password);

    /**
     * 获取用户所有权限（包括角色权限和+-权限）
     *
     * @param userId the user id
     * @return the permission list
     */
    List<Permission> getPermissionList(String userId);
}
