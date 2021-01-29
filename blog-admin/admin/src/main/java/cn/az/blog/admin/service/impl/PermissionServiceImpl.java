package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.entity.Permission;
import cn.az.blog.admin.mapper.PermissionMapper;
import cn.az.blog.admin.service.IPermissionService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * @author az
 * @version 2021-01-29
 */
@Service
public class PermissionServiceImpl extends ServiceImpl<PermissionMapper, Permission> implements IPermissionService {

}
