Go List Manipulation API
Objective : A Go program using the Gin router that maintains a list of integers and updates it based on the sign of input numbers.

How it works

If the list is empty, the number is appended.
If the sign of the input matches the existing list, append the number.
If the sign of the input is different, reduce elements from the front of the list (FIFO) until the input is exhausted.
If the input is 0, request is rejected as invalid.

API Endpoints
1. Add a number
Request:
POST /number
Body (JSON):

{ "number": 5 }

Response (example):

{
  "updated_list": [5]
}


2. Get current list

Request:
GET /number


Response (example):

{
  "current_list": [5, 3, 2]
}
