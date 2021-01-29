package cn.az.blog.admin.entity;

import com.baomidou.mybatisplus.annotation.TableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.experimental.Accessors;

import java.io.Serializable;
import java.time.LocalDateTime;

/**
 * <p>用户信息表</p>
 *
 * @author az
 * @version 2021-01-29
 */
@Data
@EqualsAndHashCode(callSuper = false)
@Accessors(chain = true)
@TableName("tb_user")
@ApiModel(value = "User对象", description = "用户信息表")
public class User implements Serializable {
    private String id;
    @ApiModelProperty(value = "用户名称")
    private String username;
    @ApiModelProperty(value = "用户密码")
    private String password;
    @ApiModelProperty(value = "性别：1 男,2 女, 3 未知")
    private Integer gender;
    @ApiModelProperty(value = "生日")
    private LocalDateTime birthday;
    @ApiModelProperty(value = "最近一次登录时间")
    private LocalDateTime lastLoginTime;
    @ApiModelProperty(value = "最近一次登录IP地址")
    private String lastLoginIp;
    @ApiModelProperty(value = "用户昵称或网络名称")
    private String nickname;
    @ApiModelProperty(value = "用户手机号码")
    private String mobile;
    @ApiModelProperty(value = "用户头像图片")
    private String avatar;
    @ApiModelProperty(value = "1 活跃, 2 冻结, 3 废弃, 4 注销")
    private Integer state;
    @ApiModelProperty(value = "创建人")
    private String createUser;
    @ApiModelProperty(value = "创建时间")
    private LocalDateTime createTime;
    @ApiModelProperty(value = "更新人")
    private String updateUser;
    @ApiModelProperty(value = "更新时间")
    private LocalDateTime updateTime;
    @ApiModelProperty(value = "是否已删除(0-未删除,1-已删除)")
    private Integer isDelete;


}
