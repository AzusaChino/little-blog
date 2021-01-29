package cn.az.blog.admin.entity;

import com.baomidou.mybatisplus.annotation.TableName;
import io.swagger.annotations.ApiModel;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.experimental.Accessors;

import java.io.Serializable;

/**
 * <p>角色权限关联表</p>
 *
 * @author az
 * @version 2021-01-29
 */
@Data
@EqualsAndHashCode(callSuper = false)
@Accessors(chain = true)
@TableName("tb_role_permission_relation")
@ApiModel(value = "RolePermissionRelation对象", description = "角色权限关联表")
public class RolePermissionRelation implements Serializable {
    private String id;
    private String roleId;
    private String permissionId;


}
