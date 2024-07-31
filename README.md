Proxy Server Documentation
==========================

Overview
--------

This proxy server handles HTTP requests from clients, forwards them to specified third-party services, and returns the responses to clients in JSON format. It also saves the requests and responses locally for future reference.

Features
--------

*   Handles multiple HTTP methods (GET, POST, etc.)
*   Forwards requests to third-party services with custom headers
*   Returns responses in JSON format including status, headers, and content length
*   Saves requests and responses locally

Request Format
--------------

The server expects a JSON request with the following fields:

    {
      "method": "GET",
      "url": "http://google.com",
      "headers": { "Authentication": "Basic bG9naW46cGFzc3dvcmQ=", ... }
    }

Response Format
---------------

The response to the client will be in JSON format with the following fields:

    {
      "id": "requestId",
      "status": <HTTP status code>,
      "headers": { <array of headers from the third-party service response> },
      "length": <length of the response content>
    }

    

HTTP Methods
------------

*   GET: Retrieve data
*   POST: Send data
*   PUT: Update data
*   DELETE: Remove data

HTTP Response Codes
-------------------

*   **200 OK:** Successful request
*   **201 Created:** Resource created successfully
*   **400 Bad Request:** Invalid request
*   **404 Not Found:** Resource not found
*   **500 Internal Server Error:** Server encountered an error
  
How to run it locally
------------------------
**Required Tools: Docker**

**git clone https://github.com/muhityessenin/proxyserver**

**make build**

**make up**

**Those commands will create docker containers and run the program**
