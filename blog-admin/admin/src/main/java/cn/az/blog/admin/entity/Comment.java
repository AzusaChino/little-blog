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
 * <p>文章评论表</p>
 *
 * @author az
 * @version 2021-01-29
 */
@Data
@EqualsAndHashCode(callSuper = false)
@Accessors(chain = true)
@TableName("tb_comment")
@ApiModel(value = "Comment对象", description = "文章评论表")
public class Comment implements Serializable {
    private String id;
    @ApiModelProperty(value = "文章id")
    private String articleId;
    private String pid;
    @ApiModelProperty(value = "昵称")
    private String nickname;
    @ApiModelProperty(value = "联系email")
    private String email;
    @ApiModelProperty(value = "评论内容")
    private String content;
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
