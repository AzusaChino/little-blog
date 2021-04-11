package cn.az.blog.admin.mapper;

import cn.az.blog.admin.entity.Permission;
import com.baomidou.mybatisplus.core.mapper.BaseMapper;
import org.apache.ibatis.annotations.Param;
import org.springframework.stereotype.Repository;

import java.util.List;

/**
 * @author az
 * @version 2021-01-29
 */
@Repository
public interface PermissionMapper extends BaseMapper<Permission> {

    /**
     * query all permissions by userId
     */
    List<Permission> queryByUserId(@Param("userId") String userId);
}
