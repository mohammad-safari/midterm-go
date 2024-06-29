package simulations;

import static io.gatling.javaapi.core.CoreDsl.atOnceUsers;
import static io.gatling.javaapi.core.CoreDsl.jsonPath;
import static io.gatling.javaapi.core.CoreDsl.scenario;
import static io.gatling.javaapi.http.HttpDsl.http;
import static io.gatling.javaapi.http.HttpDsl.status;

import io.gatling.javaapi.core.ScenarioBuilder;
import io.gatling.javaapi.core.Simulation;
import io.gatling.javaapi.http.HttpProtocolBuilder;

public class ListBasketsSimulation extends Simulation {

        private static final String baseUrl = "http://localhost:8080"; // Replace with actual URL

        private static final HttpProtocolBuilder httpProtocol = http
                        .baseUrl(baseUrl) // Here is the root for all relative URLs
                        .acceptHeader("application/json"); // Common headers

        private static final ScenarioBuilder scn = scenario("List Baskets")
                        .exec(http("Get Baskets")
                                        .get("/basket/")
                                        .check(status().is(200)) // Assert successful response
                                        .check(jsonPath("$.length").ofInt().gt(0)) // Check for non-empty list
                                                                                   // (optional)
                        );

        {

                setUp(scn.injectOpen(atOnceUsers(10))).protocols(httpProtocol);
        }

}