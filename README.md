A very basic JSONApi based REST server that stores and retrieves prompt feedback items.

## Endpoints:

* `GET /api/feedback`: Get all feedback items. Requires the `X-FEEDBACK-APIKEY` header.
* `GET /api/feedback/:id`: Get one feedback item. Requires the `X-FEEDBACK-APIKEY` header.
* `POST /api/feedback`: Create a feedback item. Requires the `AUTHORIZATION: Basic <Octopus Bearer Token>` header.

## Environment Variables

* `FUNCTIONS_CUSTOMHANDLER_PORT`: The port to run on
* `FEEDBACK_SERVICE_API_KEY`: The API key used to get the feedback items
