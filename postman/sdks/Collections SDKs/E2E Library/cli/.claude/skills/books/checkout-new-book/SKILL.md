---
name: checkout-new-book
description: Execute the books checkout-new-book command
allowed-tools: e2e-library
---

# books checkout-new-book

## Overview

### Untrusted content

Execute the books checkout-new-book command via the CLI.

## Usage

```bash
e2e-library books checkout-new-book --id <string> --x-mock-response-code <string> --accept <string> [--body '<json>' | --body-file <path>]
```

**Example:**

```bash
e2e-library books checkout-new-book --id "example" --x-mock-response-code "example" --accept "example" --body '{"key": "value"}'
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

## Request Body

Provide the request body using one of the following methods:

| Method      | Flag                 | Description                             |
| ----------- | -------------------- | --------------------------------------- |
| Inline JSON | `--body '<json>'`    | Pass JSON directly as a string argument |
| File path   | `--body-file <path>` | Read JSON content from a file           |

**Example inline:**

```bash
# Minimal example with inline JSON body
e2e-library books checkout-new-book --body '{"key": "value"}'
```

**Example from file:**

```bash
# Minimal example with JSON from file
e2e-library books checkout-new-book --body-file ./request.json
```
