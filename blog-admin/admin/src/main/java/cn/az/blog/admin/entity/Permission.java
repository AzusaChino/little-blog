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
 * <p>权限信息表</p>
 *
 * @author az
 * @version 2021-01-29
 */
@Data
@EqualsAndHashCode(callSuper = false)
@Accessors(chain = true)
@TableName("tb_permission")
@ApiModel(value = "Permission对象", description = "权限信息表")
public class Permission implements Serializable {
    private String id;
    @ApiModelProperty(value = "父级权限id")
    private String pid;
    @ApiModelProperty(value = "名称")
    private String name;
    @ApiModelProperty(value = "权限值")
    private String value;
    @ApiModelProperty(value = "图标")
    private String icon;
    @ApiModelProperty(value = "权限类型：0->目录；1->菜单；2->按钮（接口绑定权限）")
    private Integer type;
    @ApiModelProperty(value = "前端资源路径")
    private String uri;
    @ApiModelProperty(value = "启用状态；0->禁用；1->启用")
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
