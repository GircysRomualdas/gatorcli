# Gator - A CLI RSS Feed Aggregator

**Gator** is command-line tool written in Go that allows users to subscribe to RSS feeds, store and browse posts, and manage their subscriptions — all backed by a PostgreSQL database.

---

## Prerequisites

Before running Gator, make sure the following are installed on your system:

- **[Go](https://golang.org/dl/)** (version 1.20+)
- **[PostgreSQL](https://www.postgresql.org/download/)** DB

---

## Installation

1. **Clone the repository:**

   ```bash
   clone https://github.com/GircysRomualdas/gatorcli.git
   cd gator
   go mod tidy
   ```

## Configuration

Gator uses a `.gatorconfig.json` file in the home directory to store:

- The current user
- PostgreSQL connection credentials

### Example `.gatorconfig.json`:

```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "your_username"
}
```

Add .gatorconfig.json to ~ (home).

---

## Usage

Gator supports a variety of CLI commands. After clone, run commands like this:

```bash
go run . <command> [arguments]
```

### User Management

- **Register a user**

  ```bash
  go run . register <username>
  ```

- **Login as a user**

  ```bash
  go run . login <username>
  ```

- **Delete all user's**

  ```bash
  go run . reset
  ```

- **List all users**

  ```bash
  go run . users
  ```

---

### Feed Management

- **Add a new feed**

  ```bash
  go run . addfeed <feed-name> <feed-url>
  ```

- **List all available feeds**

  ```bash
  go run . feeds
  ```

- **Follow a feed**

  ```bash
  go run . follow <feed-url>
  ```

- **Unfollow a feed**

  ```bash
  go run . unfollow <feed-url>
  ```

- **List feeds you’re currently following**

  ```bash
  go run . following
  ```

---

### Aggregation and Browsing

- **Aggregate (fetch and store posts)**

  ```bash
  go run . agg
  ```

- **Browse posts from feeds you follow**

  ```bash
  go run . browse
  ```

  Optionally, limit the number of results:

  ```bash
  go run . browse 5
  ```
