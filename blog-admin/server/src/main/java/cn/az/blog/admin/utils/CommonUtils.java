package cn.az.blog.admin.utils;

import java.util.UUID;

/**
 * common utils
 * @author az
 */
public class CommonUtils {

    private CommonUtils(){
        throw new RunTimeException("illegal access");
    }

    /**
     * get UUID(without "-")
     * @return uuid
     */
    public static String getUuid(){
        return UUID.randomUUID().toString().replaceAll("-", "");
    }
}