package cn.az.blog.admin.component;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.SignatureAlgorithm;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.authentication.event.InteractiveAuthenticationSuccessEvent;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.AuthenticationException;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.web.authentication.AuthenticationSuccessHandler;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.stereotype.Component;

import javax.annotation.Resource;
import javax.servlet.FilterChain;
import javax.servlet.ServletException;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Date;

/**
 * @author az
 * @since 2021-04-09 07:55
 */
@Component
public class JwtLoginFilter extends UsernamePasswordAuthenticationFilter {

    @Resource
    private AuthenticationManager authenticationManager;

    @Override
    public Authentication attemptAuthentication(HttpServletRequest request, HttpServletResponse response) throws AuthenticationException {
        try {
            User user = new ObjectMapper()
                .readValue(request.getInputStream(), User.class);

            return authenticationManager.authenticate(
                new UsernamePasswordAuthenticationToken(
                    user.getUsername(),
                    user.getPassword(),
                    new ArrayList<>())
            );
        } catch (IOException e) {
            throw new RuntimeException(e);
        }
    }

    /**
     * Default behaviour for successful authentication.
     * <ol>
     * <li>Sets the successful <tt>Authentication</tt> object on the
     * {@link SecurityContextHolder}</li>
     * <li>Informs the configured <tt>RememberMeServices</tt> of the successful login</li>
     * <li>Fires an {@link InteractiveAuthenticationSuccessEvent} via the configured
     * <tt>ApplicationEventPublisher</tt></li>
     * <li>Delegates additional behaviour to the
     * {@link AuthenticationSuccessHandler}.</li>
     * </ol>
     * <p>
     * Subclasses can override this method to continue the {@link FilterChain} after
     * successful authentication.
     *
     * @param authResult the object returned from the <tt>attemptAuthentication</tt>
     *                   method.
     */
    @Override
    protected void successfulAuthentication(HttpServletRequest request, HttpServletResponse response, FilterChain chain, Authentication authResult) throws IOException, ServletException {
        String token = Jwts.builder()
            .setSubject(((org.springframework.security.core.userdetails.User) authResult.getPrincipal()).getUsername())
            .setExpiration(new Date(System.currentTimeMillis() + 60 * 60 * 24 * 1000))
            .signWith(SignatureAlgorithm.HS512, "MyJwtSecret")
            .compact();
        response.addHeader("Authorization", "Bearer " + token);

    }
}
