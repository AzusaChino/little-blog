package cn.az.blog.admin.utils;

import com.fasterxml.jackson.core.type.TypeReference;
import com.fasterxml.jackson.databind.DeserializationFeature;
import com.fasterxml.jackson.databind.ObjectMapper;

/**
 * @author ycpang
 * @since 2021-01-29 16:31
 */
public class JsonUtils {

    private JsonUtils() {
        throw new IllegalArgumentException();
    }

    /**
     * ObjectMapper
     */
    private static final ObjectMapper OM = new ObjectMapper();

    /**
     * 序列化
     */
    public static <T> String toJson(T t) {
        try {
            return OM.writeValueAsString(t);
        } catch (Exception e) {
            throw new RuntimeException("序列化异常", e);
        }
    }

    /**
     * 简单类型-反序列化
     */
    public static <T> T toObj(String json, Class<T> clazz) {
        OM.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        try {
            return OM.readValue(json, clazz);
        } catch (Exception e) {
            throw new RuntimeException("反序列化异常", e);
        }
    }

    /**
     * 复杂类型-反序列化
     */
    public static <T> T toObj(String json, TypeReference<T> tr) {
        OM.configure(DeserializationFeature.FAIL_ON_UNKNOWN_PROPERTIES, false);
        try {
            return OM.readValue(json, tr);
        } catch (Exception e) {
            throw new RuntimeException("反序列化异常", e);
        }
    }
}
