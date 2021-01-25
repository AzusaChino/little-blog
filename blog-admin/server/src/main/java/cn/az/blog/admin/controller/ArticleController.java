package cn.az.blog.admin.controller;

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
import reactor.core.publisher.Flux;
import reactor.core.publisher.Mono;

import javax.annotation.Resource;

/**
 * <p>
 * 博客文章表 前端控制器
 * </p>
 *
 * @author az
 * @since 2021-01-25
 */
@RestController
@RequestMapping("/article")
public class ArticleController {

    @Resource
    private IArticleService articleService;

    @GetMapping
    public Flux<Article> list() {
        return Flux.fromIterable(articleService.list());
    }

    @PostMapping
    public Mono<?> add(@RequestBody Article article) {
        return Mono.just(articleService.save(article));
    }

    @PutMapping
    public Mono<?> update(@RequestBody Article article) {
        return Mono.just(articleService.updateById(article));
    }

    @DeleteMapping("/{id}")
    public Mono<?> delete(@PathVariable String id) {
        return Mono.just(articleService.removeById(id));
    }

    @GetMapping("/{id}")
    public Mono<Article> detail(@PathVariable String id) {
        return Mono.just(articleService.getById(id));
    }
}
