package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.dto.UserDetails;
import cn.az.blog.admin.entity.Permission;
import cn.az.blog.admin.entity.User;
import cn.az.blog.admin.mapper.UserMapper;
import cn.az.blog.admin.service.IUserService;
import cn.az.blog.admin.utils.JwtUtils;
import com.baomidou.mybatisplus.core.toolkit.Wrappers;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import lombok.extern.slf4j.Slf4j;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;
import java.util.Objects;

/**
 * @author az
 * @version 2021-01-29
 */
@Slf4j
@Service
public class UserServiceImpl extends ServiceImpl<UserMapper, User> implements IUserService {

    @Resource
    private JwtUtils jwtUtils;
    @Resource
    private PasswordEncoder passwordEncoder;

    @Override
    public User getUserByUsername(String username) {
        return getOne(Wrappers.<User>lambdaQuery().eq(User::getUsername, username));
    }

    @Override
    public User register(User user) {
        return null;
    }

    @Override
    public String login(String username, String password) {
        String token = null;
        try {
            UserDetails userDetails = loadUserByUsername(username);
            if (!passwordEncoder.matches(password, userDetails.getPassword())) {
                throw new BadCredentialsException("密码错误");
            }
            UsernamePasswordAuthenticationToken authentication = new UsernamePasswordAuthenticationToken(userDetails, null, userDetails.getAuthorities());
            SecurityContextHolder.getContext().setAuthentication(authentication);
            token = jwtUtils.generateToken(userDetails);
        } catch (AuthenticationException e) {
            log.warn("登陆异常: {}", e.getMessage());
        }
        return token;
    }

    @Override
    public List<Permission> getPermissionList(String userId) {
        return null;
    }

    private UserDetails loadUserByUsername(String username) {
        User user = getOne(Wrappers.<User>lambdaQuery().eq(User::getUsername, username));
        if (Objects.nonNull(user)) {
            List<Permission> permissionList = getPermissionList(user.getId());
            return new UserDetails(user, permissionList);
        }
        throw new RuntimeException("");
    }
}
