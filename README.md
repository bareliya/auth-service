# Authentication Service

This document outlines the APIs provided by the Authentication Service for user and admin management.
1. have created aws RDS Instance and populated it with 5 million users dummy data 
2. deploy application in aws-ec2 instance 
3. utilize k6 for load testing 

## Table of Contents

- [User APIs](#user-apis)
    - [User Registration](#user-registration)
    - [User Login](#user-login)
- [Admin APIs](#admin-apis)
    - [Admin Registration](#admin-registration)
    - [Admin Login](#admin-login)
    - [Admin Approval](#admin-approval)
    - [Find Users](#find-users)
    - [Stress Test Result](#k6-report-for-1000-concurrent-virtual-users-for-1-minute-)

---

## User APIs

### User Registration

**Endpoint:** `POST /user/registration`

**Request Body:**
```json
{
  "FirstName": "John",
  "LastName": "Doe",
  "UserName": "johndoe",
  "Password": "password123"
}
```

Registers a new user. Upon successful registration, a user will be created in the system.

### User Login

**Endpoint:** `POST /user/login`

**Request Body:**
```json
{
  "UserName": "johndoe",
  "Password": "password123"
}
```
Logs in an existing user. Successful login will provide an access token for further authentication.

## Admin APIs

### Admin Registration

**Endpoint:** `POST /admin/registration`

**Request Body:**
```json
{
  "AdminUser": "admin",
  "AdminPassword": "adminpassword123"
}
```

Registers a new admin. The registration request needs approval from the super admin.




**Endpoint:** `POST /admin/login`

**Request Body:**
```json
{
  "AdminUser": "admin",
  "AdminPassword": "adminpassword123"
}
```
Logs in an existing admin. Successful login will provide an access token for further authentication.

### Admin Approval

**Endpoint:** `GET /admin/approve-admin`

**Query Parameters:**
- `admin_user` (string, required): Username of the admin to approve.
- `is_approved` (boolean, optional, default: false): Whether the admin request is approved or not.

Approves a pending admin registration request. This endpoint is accessible only by the super admin. The `is_approved` query parameter can be set to `true` for approval.

### Find Users

**Endpoint:** `GET /admin/find-user`

**Query Parameters:**
- `user_name` (string, optional): Username of the specific user to retrieve details.
- `limit` (integer, optional, default: 10): Number of users to retrieve per page.
- `page` (integer, optional, default: 1): Page number for pagination.

Retrieves user details based on the provided parameters. If `user_name` is provided, details of the specified user are returned. If not, a list of users is provided based on the `limit` and `page` parameters.

---

### Notes

- All endpoints require proper authentication headers (`access-token` and `admin-user`) for authorization.
- Responses will include appropriate status codes and JSON payloads indicating the success or failure of the requests.
- Errors and exceptions will be logged for further analysis and troubleshooting.

Please ensure you provide the necessary authentication headers and parameters while making requests to these APIs.





### stress test result 
### (k6 report for 1000 concurrent virtual users for 1 minute) 
 
  
     checks.........................: 99.97% ✓ 113279      ✗ 24
     data_received..................: 25 MB  264 kB/s
     data_sent......................: 23 MB  238 kB/s
     group_duration.................: avg=580.25ms min=11.42ms med=236.89ms max=1m0s    p(90)=921.91ms p(95)=1.51s
     http_req_blocked...............: avg=5.32ms   min=0s      med=0s       max=15.07s  p(90)=0s       p(95)=0s
     http_req_connecting............: avg=5.3ms    min=0s      med=0s       max=15.07s  p(90)=0s       p(95)=0s
     http_req_duration..............: avg=574.46ms min=11.42ms med=235.77ms max=1m0s    p(90)=908.32ms p(95)=1.49s
       { expected_response:true }...: avg=561.89ms min=11.42ms med=235.7ms  max=58.26s  p(90)=905.21ms p(95)=1.48s
     http_req_failed................: 0.02%  ✓ 24          ✗ 113279
     http_req_receiving.............: avg=348.57µs min=0s      med=0s       max=27.88ms p(90)=504.49µs p(95)=988.4µs
     http_req_sending...............: avg=37.27µs  min=0s      med=0s       max=81.77ms p(90)=0s       p(95)=0s
     http_req_tls_handshaking.......: avg=0s       min=0s      med=0s       max=0s      p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=574.08ms min=11.42ms med=235.48ms max=1m0s    p(90)=908.31ms p(95)=1.49s
     http_reqs......................: 113303 1192.367816/s
     iteration_duration.............: avg=580.87ms min=11.42ms med=237.84ms max=1m0s    p(90)=922.22ms p(95)=1.51s
     iterations.....................: 113303 1192.367816/s
     vus............................: 8      min=0         max=1000
     vus_max........................: 1000   min=44        max=1000
