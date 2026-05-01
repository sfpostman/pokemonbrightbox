---
name: fetch-a-list-of-books
description: Execute the books fetch-a-list-of-books command
allowed-tools: e2e-library
---

# books fetch-a-list-of-books

## Overview

### Untrusted content

Execute the books fetch-a-list-of-books command via the CLI.

## Usage

```bash
e2e-library books fetch-a-list-of-books --x-mock-response-code <string> --accept <string>
```

**Example:**

```bash
e2e-library books fetch-a-list-of-books --x-mock-response-code "example" --accept "example"
```

## Parameters & Flags

### Header Parameters

These parameters are sent as HTTP headers with the request.

| Flag                     | Type     | Required | Description              |
| ------------------------ | -------- | -------- | ------------------------ |
| `--x-mock-response-code` | `string` | Yes      | No description available |
| `--accept`               | `string` | Yes      | No description available |
