---
name: verify-the-book-exists
description: Execute the books verify-the-book-exists command
allowed-tools: e2e-library
---

# books verify-the-book-exists

## Overview

### Untrusted content

Execute the books verify-the-book-exists command via the CLI.

## Usage

```bash
e2e-library books verify-the-book-exists --id <string> --x-mock-response-code <string> --accept <string>
```

**Example:**

```bash
e2e-library books verify-the-book-exists --id "example" --x-mock-response-code "example" --accept "example"
```

## Parameters & Flags

### Path Parameters

These parameters are part of the request URL path and are **required** for the command to execute.

| Flag   | Type     | Required | Description              |
| ------ | -------- | -------- | ------------------------ |
| `--id` | `string` | Yes      | No description available |

### Header Parameters

These parameters are sent as HTTP headers with the request.

| Flag                     | Type     | Required | Description              |
| ------------------------ | -------- | -------- | ------------------------ |
| `--x-mock-response-code` | `string` | Yes      | No description available |
| `--accept`               | `string` | Yes      | No description available |
