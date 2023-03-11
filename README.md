# Submission for Dyte - Backend Task


## The Backend for VIT Course Registration with the following functionality - 

1. There are specific slots spread across the week. The timings and IDs of these slots
can be configured by the administrator.
2. There are a number of faculties, the information regarding these faculties can also be
added by the administrator.
3. Courses are allocated at specific slots, the IDs of which were configured before (in
step 1). The administrator can specify the courses that are offered, including the IDs
of the slots at which they’re offered and the IDs of the faculties that can take the
courses.
4. Students can register for courses by specifying the course ID, slot ID(s), and faculty
ID.
5. Only 1 course can be selected at a specific time slot. The selected faculty must be
teaching the course at that slot as well.

## The project implements the following Features 

1. Students and administrators are able to log in using their registration number
and password. The database can be populated with mock data to implement this. Students
don’t have to sign up.

2. On logging in, students are able to select courses according to the constraints
mentioned above. If there’s a clash in slots for the selected course, or any request
that violates the constraints provided above, an error is returned.

3. Created RESTful API routes for CRUD of slots, faculties, and courses. Protected these
features behind authentication/authorization as required, so that only administrators
can access their routes.

4. A student is able to see their registered courses and status via an API route.
Naturally, no other students should be able to see this information.


## API Specifications 

### Admin Routes

1. Create a faculty <br>

Request : <br>

```
{
  "id": "string",
  "name": "string"
}
```

Response:  <br>

```
{
  "success": true,
  "data": {
    "id": "string",
    "name": "string"
  }
}
```
<br>
2. Create a course <br>

Request: <br>
```
{
  "id": "string",
  "name": "string",
  "slot_ids": [
    "string"
  ],
  "faculty_ids": [
    "string"
  ],
  "course_type": "THEORY"
}
```

Response: <br>
```{
  "success": true,
  "data": {
    "id": "string",
    "name": "string",
    "faculties": [
      {
        "id": "string",
        "name": "string"
      }
    ],
    "course_type": "THEORY",
    "allowed_slots": [
      {
        "id": "A1",
        "timings": [
          {
            "day": "MON",
            "start": "2019-08-24T14:15:22Z",
            "end": "2019-08-24T14:15:22Z"
          }
        ]
      }
    ]
  }
}
```
<br>
3. Create a slot <br>

Request: <br>
```
{
  "id": "A1",
  "timings": [
    {
      "day": "MON",
      "start": "2019-08-24T14:15:22Z",
      "end": "2019-08-24T14:15:22Z"
    }
  ]
}
```

Response: <br>
```
{
  "success": true,
  "data": {
    "id": "A1",
    "timings": [
      {
        "day": "MON",
        "start": "2019-08-24T14:15:22Z",
        "end": "2019-08-24T14:15:22Z"
      }
    ]
  }
}
```
<br>
4. Create a Student <br>

Request: <br>
```
{
  "success": true,
  "data": {
    "id": "string",
    "name": "string"
  }
}
```

Response: <br>
```
{
  "success": true,
  "data": {
    "id": "string",
    "name": "string"
  }
}
```

<br><br>

### Student Routes

<br>
1. Get Faculty by ID <br>

Request:

```
{
  "faculty_id" : "string"
}
```

Response:

```
{
  "success": true,
  "data": {
    "id": "string",
    "name": "string"
  }
}
```
