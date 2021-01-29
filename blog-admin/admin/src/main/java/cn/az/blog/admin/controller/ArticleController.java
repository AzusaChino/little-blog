package cn.az.blog.admin.controller;

import cn.az.blog.admin.common.CommonResult;
import cn.az.blog.admin.entity.Article;
import cn.az.blog.admin.service.IArticleService;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.annotation.Resource;
import java.util.List;

/**
 * @author az
 * @version 2021-01-29
 */
@RestController
@RequestMapping("/article")
public class ArticleController {

    @Resource
    private IArticleService articleService;

    @GetMapping
    public CommonResult<List<Article>> list() {
        return CommonResult.success(articleService.list());
    }

    @PostMapping
    public CommonResult<?> add(@RequestBody Article article) {
        return CommonResult.success(articleService.save(article));
    }

    @PutMapping
    public CommonResult<?> update(@RequestBody Article article) {
        return CommonResult.success(articleService.updateById(article));
    }

    @DeleteMapping("/{id}")
    public CommonResult<?> delete(@PathVariable String id) {
        return CommonResult.success(articleService.removeById(id));
    }

    @GetMapping("/{id}")
    public CommonResult<Article> detail(@PathVariable String id) {
        return CommonResult.success(articleService.getById(id));
    }
}
