## User CRUD api's with dynamic age calculation

The requirements already defined the tech stack and overall structure, so I focused on implementing the API correctly and incrementally. I used Supabase for PostgreSQL instead of running a local Postgres server. I wrote the SQL first and generated the corresponding Go code using SQLC.

Once the database layer was working, I built the API one endpoint at a time, added the required validation, and tested each step before moving on. This helped me avoid running into multiple issues at the end.

Since age was required to be calculated dynamically, I kept that logic in the service layer and avoided mixing it with HTTP or database code. I also added a small unit test for the age calculation without involving the database or API layer.

Logging, middleware, pagination, and containerization support were added after the core CRUD functionality was complete. The idea was to get the core features stable first and then work on secondary features.

Overall, the approach was to follow the given constraints closely, keep each part of the code focused on a single responsibility, and avoid adding complexity beyond what the assignment required.
