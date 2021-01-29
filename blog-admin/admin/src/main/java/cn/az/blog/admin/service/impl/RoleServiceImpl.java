package cn.az.blog.admin.service.impl;

import cn.az.blog.admin.entity.Role;
import cn.az.blog.admin.mapper.RoleMapper;
import cn.az.blog.admin.service.IRoleService;
import com.baomidou.mybatisplus.extension.service.impl.ServiceImpl;
import org.springframework.stereotype.Service;

/**
 * @author az
 * @version 2021-01-29
 */
@Service
public class RoleServiceImpl extends ServiceImpl<RoleMapper, Role> implements IRoleService {

}
