package cn.az.blog.admin.entity;

import com.baomidou.mybatisplus.annotation.TableName;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import lombok.Data;
import lombok.EqualsAndHashCode;

import java.io.Serializable;
import java.time.LocalDateTime;

/**
 * <p>
 * 博客文章表
 * </p>
 *
 * @author az
 * @since 2021-01-25
 */
@Data
@EqualsAndHashCode(callSuper = false)
@TableName("tb_article")
@ApiModel(value = "Article对象", description = "博客文章表")
public class Article implements Serializable {

    private static final long serialVersionUID = 1L;

    @ApiModelProperty(value = "主键id")
    private String id;

    @ApiModelProperty(value = "文章主题")
    private String topic;

    @ApiModelProperty(value = "缩略图url")
    private String thumbnail;

    @ApiModelProperty(value = "文章内容")
    private String content;

    @ApiModelProperty(value = "发布状态")
    private Integer publishState;

    @ApiModelProperty(value = "发布时间")
    private LocalDateTime publishTime;

    @ApiModelProperty(value = "创建人")
    private String createUser;

    @ApiModelProperty(value = "创建时间")
    private LocalDateTime createTime;

    private String updateUser;

    private LocalDateTime updateTime;

    @ApiModelProperty(value = "是否已删除(0-未删除,1-已删除)")
    private Integer isDelete;


}
