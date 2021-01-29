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
 * <p>参数信息表</p>
 *
 * @author az
 * @version 2021-01-29
 */
@Data
@EqualsAndHashCode(callSuper = false)
@Accessors(chain = true)
@TableName("tb_parameter")
@ApiModel(value = "Parameter对象", description = "参数信息表")
public class Parameter implements Serializable {
    @ApiModelProperty(value = "主键id")
    private String id;
    @ApiModelProperty(value = "父key")
    private String parent;
    @ApiModelProperty(value = "key")
    private String key;
    @ApiModelProperty(value = "value")
    private String value;
    @ApiModelProperty(value = "状态")
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
