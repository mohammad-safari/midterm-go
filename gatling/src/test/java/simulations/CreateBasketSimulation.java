package simulations;

import static io.gatling.javaapi.core.CoreDsl.StringBody;
import static io.gatling.javaapi.core.CoreDsl.atOnceUsers;
import static io.gatling.javaapi.core.CoreDsl.jsonPath;
import static io.gatling.javaapi.core.CoreDsl.scenario;
import static io.gatling.javaapi.http.HttpDsl.http;
import static io.gatling.javaapi.http.HttpDsl.status;

import io.gatling.javaapi.core.ScenarioBuilder;
import io.gatling.javaapi.core.Simulation;
import io.gatling.javaapi.http.HttpProtocolBuilder;

public class CreateBasketSimulation extends Simulation {

        private static final String baseUrl = "http://localhost:8080";
        private static final String basketData = "{\n  \"state\":\"PENDING\"\n}"; // Sample basket data

        private static final HttpProtocolBuilder httpProtocol = http
                        .baseUrl(baseUrl) // Here is the root for all relative URLs
                        .acceptHeader("application/json") // Common headers
                        .contentTypeHeader("application/json"); // Ensure content type is set
                        
        private static final ScenarioBuilder scn = scenario("Create Basket")
                        .exec(http("Create Basket")
                                        .post("/basket")
                                        .body(StringBody(basketData))
                                        .check(status().is(201)) // Assert created status
                                        .check(jsonPath("$.id").saveAs("basketId")) // Save basket ID for future use
                        );

        {
                setUp(scn.injectOpen(atOnceUsers(10))).protocols(httpProtocol);
        }
}
