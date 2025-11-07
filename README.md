### 🧩 PROJECT NAME

## Course Enrollment System 🎓

# A backend system to manage students, courses, teachers, enrollments, attendance, and grades.

### 🧱 DATABASE STRUCTURE
## 1️⃣ students

| Column     | Type              | Description       |
| ---------- | ----------------- | ----------------- |
| id         | SERIAL PK         | Unique student ID |
| name       | VARCHAR(255)      | Student name      |
| email      | VARCHAR(255)      | Unique email      |
| user_id    | INT FK → users.id | Unique email      |
| created_at | TIMESTAMP         | Auto timestamp    |

## 2️⃣ teachers

| Column     | Type              | Description       |
| ---------- | ----------------- | ----------------- |
| id         | SERIAL PK         | Unique teacher ID |
| name       | VARCHAR(255)      | Teacher name      |
| email      | VARCHAR(255)      | Unique email      |
| user_id    | INT FK → users.id | Unique email      |
| created_at | TIMESTAMP         | Auto timestamp    |

## 3️⃣ courses

| Column      | Type                 | Description        |
| ----------- | -------------------- | ------------------ |
| id          | SERIAL PK            | Unique course ID   |
| title       | VARCHAR(255)         | Course name        |
| description | TEXT                 | Course description |
| teacher_id  | INT FK → teachers.id | Course instructor  |
| created_at  | TIMESTAMP            | Auto timestamp     |

## 4️⃣ enrollments

| Column      | Type                 | Description          |
| ----------- | -------------------- | -------------------- |
| id          | SERIAL PK            | Enrollment ID        |
| student_id  | INT FK → students.id | Enrolled student     |
| course_id   | INT FK → courses.id  | Course enrolled      |
| enrolled_at | TIMESTAMP            | Enrollment timestamp |

## 5️⃣ attendance

| Column        | Type                    | Description                   |
| ------------- | ----------------------- | ----------------------------- |
| id            | SERIAL PK               | Attendance ID                 |
| enrollment_id | INT FK → enrollments.id | Which student-course combo    |
| date          | DATE                    | Attendance date               |
| status        | VARCHAR(20)             | ("present", "absent", "late") |

## 6️⃣ grades

| Column        | Type                    | Description         |
| ------------- | ----------------------- | ------------------- |
| id            | SERIAL PK               | Grade ID            |
| enrollment_id | INT FK → enrollments.id | Student-course pair |
| grade         | VARCHAR(2)              | Grade letter (A–E)  |
| remarks       | TEXT                    | Optional notes      |

## 7️⃣ users

| Column      | Type                 | Description          |
| ----------- | -------------------- | -------------------- |
| id          | SERIAL PK            | User ID              |
| username    | STRING               | Username for login   |
| password    | STRING               | Hashed Password      |
| created_at  | TIMESTAMP            | Auto timestamp       |


### 🧬 ERD (Entity Relationship Diagram)

```
students ──────< enrollments >────┬─ courses ──┬── teachers
                                  │
                                  ├──< attendance
                                  │
                                  └─── grades
```

## Relationships

🧑‍🏫 One teacher → many courses

📚 One course → many enrollments

👩‍🎓 One student → many enrollments

🗓 One enrollment → many attendance records

🧾 One enrollment → many grade records


### 🧠 BUSINESS FLOW (How the system works)

## 🧍‍♀️ 1. Teacher creates a course

Teacher logs in

Sends POST /courses with title, description

The course is linked to teacher_id

## 🧑‍🎓 2. Student enrolls in a course

Student calls POST /enrollments with course_id

Creates record linking student_id and course_id

## 🕒 3. Attendance tracking

Teacher marks attendance per class date

POST /attendance → save record with status (present/absent/late)

## 🧾 4. Assigning grades

After course ends, teacher assigns grades

POST /grades → set letter grade + remarks

## 📊 5. Viewing data

Students can:

GET /my-courses → list all courses they’re enrolled in

GET /attendance/:course_id → view attendance history

GET /grades/:course_id → see grades

Teachers can:

GET /my-courses → view their courses

GET /courses/:id/students → view enrolled students and their progress