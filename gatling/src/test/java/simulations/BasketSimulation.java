package simulations;

import com.fasterxml.jackson.databind.ObjectMapper;
import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;

import java.util.HashMap;
import java.util.Iterator;
import java.util.Map;
import java.util.concurrent.ThreadLocalRandom;
import java.util.function.Supplier;
import java.util.stream.Stream;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;

public class BasketSimulation extends Simulation {

    private static final String baseUrl = "http://localhost:8080";

    // Feeder to generate random user data
    private static final Iterator<Map<String, Object>> userFeeder = Stream
            .generate((Supplier<Map<String, Object>>) () -> {
                Map<String, Object> data = new HashMap<>();
                data.put("username", "user" + ThreadLocalRandom.current().nextInt(1, 10000));
                data.put("password", "password");
                return data;
            }).iterator();

    // Feeder to generate random basket data
    private static final Iterator<Map<String, Object>> basketFeeder = Stream
            .generate((Supplier<Map<String, Object>>) () -> {
                Map<String, Object> data = new HashMap<>();
                data.put("state", "PENDING");
                data.put("data", generateRandomJsonData());
                return data;
            }).iterator();

    private static String generateRandomJsonData() {
        Map<String, Object> jsonMap = new HashMap<>();
        jsonMap.put("key1", "value" + ThreadLocalRandom.current().nextInt(1, 100));
        jsonMap.put("key2", "value" + ThreadLocalRandom.current().nextInt(1, 100));
        ObjectMapper mapper = new ObjectMapper();
        try {
            return mapper.writeValueAsString(jsonMap);
        } catch (Exception e) {
            e.printStackTrace();
            return "{}";
        }
    }

    private static final HttpProtocolBuilder httpProtocol = http
            .baseUrl(baseUrl) // Here is the root for all relative URLs
            .acceptHeader("application/json") // Common headers
            .contentTypeHeader("application/json"); // Ensure content type is set

    private static final ScenarioBuilder scn = scenario("User Basket Operations")
            .feed(userFeeder)
            .exec(http("Create User")
                    .post("/user")
                    .body(StringBody(session -> {
                        Map<String, Object> jsonData = new HashMap<>();
                        jsonData.put("username", session.getString("username"));
                        jsonData.put("password", session.getString("password"));
                        ObjectMapper mapper = new ObjectMapper();
                        try {
                            return mapper.writeValueAsString(jsonData);
                        } catch (Exception e) {
                            e.printStackTrace();
                            return "{}";
                        }
                    }))
                    .check(status().is(201))
                    .check(jsonPath("$.id").saveAs("userId")))
            .repeat(5).on(
                    feed(basketFeeder)
                            .exec(http("Create Basket")
                                    .post("/basket")
                                    .body(StringBody(session -> {
                                        Map<String, Object> jsonData = new HashMap<>();
                                        jsonData.put("state", session.getString("state"));
                                        // jsonData.put("data", session.getString("data").getBytes());
                                        jsonData.put("user_id", session.getInt("userId"));
                                        ObjectMapper mapper = new ObjectMapper();
                                        try {
                                            return mapper.writeValueAsString(jsonData);
                                        } catch (Exception e) {
                                            e.printStackTrace();
                                            return "{}";
                                        }
                                    }))
                                    .check(status().is(201))
                                    .check(jsonPath("$.id").saveAs("basketId"))))
            .repeat(2).on(
                    exec(http("Update Basket")
                            .patch("/basket/#{basketId}")
                            .body(StringBody(session -> {
                                Map<String, Object> jsonData = new HashMap<>();
                                jsonData.put("state", "PENDING");
                                jsonData.put("data", generateRandomJsonData().getBytes());
                                jsonData.put("user_id", session.getInt("userId"));
                                ObjectMapper mapper = new ObjectMapper();
                                try {
                                    return mapper.writeValueAsString(jsonData);
                                } catch (Exception e) {
                                    e.printStackTrace();
                                    return "{}";
                                }
                            }))
                            .check(status().is(200))))
                .repeat(1).on( // last change
                    exec(http("Update Basket")
                            .patch("/basket/#{basketId}")
                            .body(StringBody(session -> {
                                Map<String, Object> jsonData = new HashMap<>();
                                jsonData.put("state", "COMPLETED");
                                jsonData.put("data", generateRandomJsonData().getBytes());
                                jsonData.put("user_id", session.getInt("userId"));
                                ObjectMapper mapper = new ObjectMapper();
                                try {
                                    return mapper.writeValueAsString(jsonData);
                                } catch (Exception e) {
                                    e.printStackTrace();
                                    return "{}";
                                }
                            }))
                            .check(status().is(200))))
            .repeat(1).on(
                    exec(http("Delete Basket")
                            .delete("/basket/#{basketId}")
                            .check(status().is(204))));

    {
        setUp(scn.injectOpen(atOnceUsers(10))).protocols(httpProtocol);
    }
}
