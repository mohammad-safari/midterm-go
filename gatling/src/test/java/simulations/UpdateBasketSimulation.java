package simulations;

import static io.gatling.javaapi.core.CoreDsl.StringBody;
import static io.gatling.javaapi.core.CoreDsl.atOnceUsers;
import static io.gatling.javaapi.core.CoreDsl.scenario;
import static io.gatling.javaapi.http.HttpDsl.http;
import static io.gatling.javaapi.http.HttpDsl.status;

import io.gatling.javaapi.core.ScenarioBuilder;
import io.gatling.javaapi.core.Simulation;
import io.gatling.javaapi.http.HttpProtocolBuilder;

public class UpdateBasketSimulation extends Simulation {

        private static final String baseUrl = "http://localhost:8080";
        private static final String basketId = "your_basket_id"; // Replace with actual ID
        private static final HttpProtocolBuilder httpProtocol = http
                        .baseUrl(baseUrl) // Here is the root for all relative URLs
                        .acceptHeader("application/json"); // Common headers

        private static final String updateData = "{\n" +
                        "  \"name\": \"Updated Basket\"\n" +
                        "}";

        private static final ScenarioBuilder scn = scenario("Update Basket")
                        .exec(http("Update Basket")
                                        .patch(String.format("/basket/%s", basketId))
                                        .body(StringBody(updateData))
                                        .check(status().is(200)) // Assert successful update
                        );

        {
                setUp(scn.injectOpen(atOnceUsers(10))).protocols(httpProtocol);
        }
}
