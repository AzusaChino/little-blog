package cn.az.blog.admin.entity;

import com.baomidou.mybatisplus.annotation.TableName;
import io.swagger.annotations.ApiModel;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.experimental.Accessors;

import java.io.Serializable;

/**
 * <p>用户角色关联表</p>
 *
 * @author az
 * @version 2021-01-29
 */
@Data
@EqualsAndHashCode(callSuper = false)
@Accessors(chain = true)
@TableName("tb_user_role_relation")
@ApiModel(value = "UserRoleRelation对象", description = "用户角色关联表")
public class UserRoleRelation implements Serializable {
    private String id;
    private String userId;
    private String roleId;


}
