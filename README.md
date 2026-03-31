# Stability Team Technical Test Submission

```markdown
# Stability Team Technical Test - Task Manager API

A thread-safe Task Manager API built with Go and Fiber. This repository contains bug fixes and stability improvements submitted for the Fullstack Developer Internship technical assessment.

## Setup & Run
go mod tidy
go run main.go
```
The server will run locally at: `http://localhost:3000`

---

## 🐞 Issues Found & Fixes Implemented

### 1. Handlers (`handlers/task_handler.go`)

* **Unhandled Type Conversions** 
  * **Location:** `GetTask` (Lines 25-31) and `DeleteTask` (Lines 66-72)
  * **Issue:** `strconv.Atoi` errors were ignored using the blank identifier (`_`), leading to silent failures on invalid input.
  * **Fix:** Added explicit error handling to return a `400 Bad Request` if the ID format is invalid.

* **Incorrect HTTP Status Codes**
  * **Location:** `GetTask` (Lines 36-40) and `CreateTask` (Line 60)
  * **Issue:** `GetTask` returned a `200 OK` for missing tasks. `CreateTask` returned an implicit `200 OK` instead of the standard creation code.
  * **Fix:** Enforced `404 Not Found` for missing resources and `201 Created` via `c.Status(fiber.StatusCreated)` upon successful task generation.

* **Missing Input Validation**
  * **Location:** `CreateTask` (Lines 51-55)
  * **Issue:** Accepted completely empty JSON payloads, allowing ghost tasks.
  * **Fix:** Added basic validation `if task.Title == ""` to reject invalid requests with a `400 Bad Request`.

### 2. Data Store (`store/task_store.go`)

* **Loop Variable Pointer Trap**
  * **Location:** `GetTaskByID` (Lines 27-30)
  * **Issue:** Returned a pointer to the shared loop variable `t` (`return &t`), creating a risk of returning incorrect data or accidental memory mutation.
  * **Fix:** Created a safe local copy (`task := t`) and returned its pointer (`return &task`).

* **Missing ID Generation**
  * **Location:** `AddTask` (Lines 38-40)
  * **Issue:** Did not assign unique IDs to new tasks, defaulting all to `0`.
  * **Fix:** Implemented an auto-incrementing `nextID` integer at the package level to handle resource indexing properly.

* **Slice Manipulation Panic**
  * **Location:** `DeleteTask` (Lines 48-52)
  * **Issue:** Modified the slice during iteration without exiting the loop, risking an index out-of-bounds panic.
  * **Fix:** Added a `break` statement to exit the loop immediately after the target task is deleted.

---

## 🛠️ System Improvement: Thread Safety

**Implemented Concurrency Control via `sync.Mutex`**

* **Location:** `store/task_store.go` (Lines 13-14, and injected into all data operations)
* **Issue:** The most critical stability flaw in the original codebase was the lack of thread safety. Because Go HTTP servers handle incoming requests concurrently via goroutines, simultaneous `POST` or `DELETE` requests to the global `Tasks` slice trigger a **concurrent slice write panic**, instantly crashing the application.
* **Solution:** Introduced a `sync.Mutex` (`mu`) to the data store.
* **Implementation:** Wrapped all read and write functions (`GetAllTasks`, `GetTaskByID`, `AddTask`, `DeleteTask`) with `mu.Lock()` and `defer mu.Unlock()` at the start of their execution blocks.
* **Result:** Guarantees mutually exclusive access to the `Tasks` memory slice. This prevents race conditions and ensures the API remains completely stable and crash-free under heavy concurrent load.
