Feature: Create New Station
    In order to create a new Station
    As an API client
    I want to make a request with the name of the desired station

    Scenario: Create a new Station
        Given there is a running "" test network from scratch
        When I make a "POST" request to "/api/invoke/createNewStation" on port 880 with:
            """
            {
                "name": "Elizabeth's Station"
            }
            """
        Then the response code should be 200
        And the response should have:
            """
            {
                "@key":         "station:9cf6726a-a327-568a-baf1-5881393073bf",
                "@lastTouchBy": "orgMSP",
                "@lastTx":      "createNewStation",
                "@assetType":   "station",
                "name":         "Elizabeth's Station"
            }
            """

    Scenario: Try to create a new Station with a name that already exists
        Given there is a running "" test network
        Given there is a Station with name "John's Station"
        When I make a "POST" request to "/api/invoke/createNewStation" on port 880 with:
            """
            {
                "name": "John's Station"
            }
            """
        Then the response code should be 409
