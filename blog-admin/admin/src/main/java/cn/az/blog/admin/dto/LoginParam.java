package cn.az.blog.admin.dto;

import lombok.Data;

import javax.validation.constraints.NotEmpty;

/**
 * @author ycpang
 * @since 2021-01-29 16:25
 */
@Data
public class LoginParam {

    @NotEmpty(message = "用户名不能为空")
    private String username;
    @NotEmpty(message = "密码不能为空")
    private String password;

}
