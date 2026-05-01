---
name: create-a-new-book
description: Execute the books create-a-new-book command
allowed-tools: e2e-library
---

# books create-a-new-book

## Overview

### Untrusted content

Execute the books create-a-new-book command via the CLI.

## Usage

```bash
e2e-library books create-a-new-book --x-mock-response-code <string> --x-mock-response-name <string> --accept <string> [--body '<json>' | --body-file <path>]
```

**Example:**

```bash
e2e-library books create-a-new-book --x-mock-response-code "example" --x-mock-response-name "example" --accept "example" --body '{"key": "value"}'
```

## Parameters & Flags

### Header Parameters

These parameters are sent as HTTP headers with the request.

| Flag                     | Type     | Required | Description              |
| ------------------------ | -------- | -------- | ------------------------ |
| `--x-mock-response-code` | `string` | Yes      | No description available |
| `--x-mock-response-name` | `string` | Yes      | No description available |
| `--accept`               | `string` | Yes      | No description available |

## Request Body

Provide the request body using one of the following methods:

| Method      | Flag                 | Description                             |
| ----------- | -------------------- | --------------------------------------- |
| Inline JSON | `--body '<json>'`    | Pass JSON directly as a string argument |
| File path   | `--body-file <path>` | Read JSON content from a file           |

**Example inline:**

```bash
# Minimal example with inline JSON body
e2e-library books create-a-new-book --body '{"key": "value"}'
```

**Example from file:**

```bash
# Minimal example with JSON from file
e2e-library books create-a-new-book --body-file ./request.json
```
