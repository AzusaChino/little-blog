package cn.az.blog.admin.controller;

import cn.az.blog.admin.common.CommonResult;
import cn.az.blog.admin.dto.LoginParam;
import cn.az.blog.admin.service.IUserService;
import cn.az.blog.admin.utils.RedisUtils;
import cn.hutool.core.util.StrUtil;
import com.google.common.collect.Maps;
import io.swagger.annotations.Api;
import io.swagger.annotations.ApiOperation;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import redis.clients.jedis.params.SetParams;

import javax.annotation.Resource;
import java.util.Map;
import java.util.Objects;

/**
 * @author az
 * @version 2021-01-29
 */
@Api(value = "用户管理API", tags = "UserController")
@RestController
@RequestMapping("/user")
public class UserController {

    @Resource
    private IUserService userService;

    @Resource
    private RedisUtils redisUtils;

    @Value("${jwt.tokenHeader}")
    private String tokenHeader;

    @ApiOperation("登录")
    @PostMapping("/login")
    public CommonResult<?> login(@RequestBody LoginParam lp) {
        String token = userService.login(lp.getUsername(), lp.getPassword());
        if (Objects.isNull(token) || StrUtil.isBlank(token)) {
            return CommonResult.validateFailed("用户名或密码错误");
        }
        Map<String, String> map = Maps.newHashMap();
        map.put("token", token);
        map.put("tokenHeader", tokenHeader);
        redisUtils.set("USER_" + lp.getUsername(), token,
            SetParams.setParams()
                .nx()
                .px(60000L));
        return CommonResult.success(map);
    }

    @ApiOperation("获取用户权限")
    @GetMapping("/permission/{userId}")
    public CommonResult<?> getPermissionList(@PathVariable String userId) {
        return CommonResult.success(userService.getPermissionList(userId));
    }
}
