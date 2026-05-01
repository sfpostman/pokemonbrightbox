const http = require("http");
const PORT = process.env.PORT || 4501;

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

function readBody(req) {
  return new Promise((resolve) => {
    let data = "";
    req.on("data", (chunk) => (data += chunk));
    req.on("end", () => {
      try {
        resolve(data ? JSON.parse(data) : {});
      } catch {
        resolve({});
      }
    });
  });
}

function json(res, status, body) {
  const payload = JSON.stringify(body);
  res.writeHead(status, {
    "Content-Type": "application/json",
    "x-powered-by": "Postman Local Mock",
  });
  res.end(payload);
}

function uuid() {
  return "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx".replace(/[xy]/g, (c) => {
    const r = (Math.random() * 16) | 0;
    return (c === "x" ? r : (r & 0x3) | 0x8).toString(16);
  });
}

// ---------------------------------------------------------------------------
// Seed data — mirrors the spec examples so the twin starts populated
// ---------------------------------------------------------------------------
const SEED_BOOKS = [
  {
    id: "7f04875b-9201-4c8f-b381-18370f9b2dfb",
    title: "One Hundred Years of Solitude",
    author: "Gabriel García Márquez",
    genre: "fiction",
    yearPublished: 1967,
    checkedOut: false,
    isPermanentCollection: true,
    createdAt: "2024-09-27T15:48:46.962Z",
  },
  {
    id: "a1b2c3d4-0000-4000-8000-000000000001",
    title: "The Great Gatsby",
    author: "F. Scott Fitzgerald",
    genre: "fiction",
    yearPublished: 1925,
    checkedOut: false,
    isPermanentCollection: true,
    createdAt: "2024-09-27T15:48:46.962Z",
  },
  {
    id: "a1b2c3d4-0000-4000-8000-000000000002",
    title: "To Kill a Mockingbird",
    author: "Harper Lee",
    genre: "fiction",
    yearPublished: 1960,
    checkedOut: true,
    isPermanentCollection: false,
    createdAt: "2024-09-27T15:48:46.962Z",
  },
];

// ---------------------------------------------------------------------------
// Server
// ---------------------------------------------------------------------------
const server = http.createServer(async (req, res) => {
  const { method, url } = req;

  // Strip query string for routing
  const path = url.split("?")[0];

  // -------------------------------------------------------------------------
  // @endpoint GET /health
  // -------------------------------------------------------------------------
  if (method === "GET" && path === "/health") {
    return json(res, 200, { status: "ok", service: "library-api-mock" });
  }

  // -------------------------------------------------------------------------
  // @endpoint GET /books
  // -------------------------------------------------------------------------
  if (method === "GET" && path === "/books") {
    // Ensure seed data is present on first run
    let books = await pm.state.get("library:books");
    if (!books) {
      books = SEED_BOOKS;
      await pm.state.set("library:books", books);
    }
    return json(res, 200, books);
  }

  // -------------------------------------------------------------------------
  // @endpoint POST /books
  // -------------------------------------------------------------------------
  if (method === "POST" && path === "/books") {
    const body = await readBody(req);

    // Validate required fields (mirrors real API behaviour)
    if (!body.title || !body.author) {
      return json(res, 400, {
        error: { code: 400, message: "title and author are required." },
      });
    }

    const newBook = {
      id: uuid(),
      title: body.title,
      author: body.author,
      genre: body.genre || "fiction",
      yearPublished: body.yearPublished ? Number(body.yearPublished) : null,
      checkedOut: false,                  // always false on creation
      isPermanentCollection: false,
      createdAt: new Date().toISOString(),
    };

    let books = await pm.state.get("library:books");
    if (!books) books = SEED_BOOKS;
    books = [...books, newBook];
    await pm.state.set("library:books", books);

    return json(res, 201, newBook);
  }

  // -------------------------------------------------------------------------
  // @endpoint GET /books/:id
  // -------------------------------------------------------------------------
  const getMatch = path.match(/^\/books\/([^/]+)$/);
  if (method === "GET" && getMatch) {
    const id = decodeURIComponent(getMatch[1]);
    let books = await pm.state.get("library:books");
    if (!books) {
      books = SEED_BOOKS;
      await pm.state.set("library:books", books);
    }
    const book = books.find((b) => b.id === id);
    if (!book) {
      return json(res, 404, {
        error: { code: 404, message: "The book you are searching for cannot be found." },
      });
    }
    return json(res, 200, book);
  }

  // -------------------------------------------------------------------------
  // @endpoint PATCH /books/:id
  // -------------------------------------------------------------------------
  const patchMatch = path.match(/^\/books\/([^/]+)$/);
  if (method === "PATCH" && patchMatch) {
    const id = decodeURIComponent(patchMatch[1]);
    const body = await readBody(req);

    let books = await pm.state.get("library:books");
    if (!books) {
      books = SEED_BOOKS;
      await pm.state.set("library:books", books);
    }

    const idx = books.findIndex((b) => b.id === id);
    if (idx === -1) {
      return json(res, 404, {
        error: { code: 404, message: "The book you are searching for cannot be found." },
      });
    }

    // Only allow patching mutable fields (checkedOut, genre, yearPublished)
    const updated = {
      ...books[idx],
      ...(body.checkedOut !== undefined && { checkedOut: Boolean(body.checkedOut) }),
      ...(body.genre !== undefined && { genre: body.genre }),
      ...(body.yearPublished !== undefined && { yearPublished: Number(body.yearPublished) }),
    };

    books = [...books.slice(0, idx), updated, ...books.slice(idx + 1)];
    await pm.state.set("library:books", books);

    return json(res, 200, updated);
  }

  // -------------------------------------------------------------------------
  // 404 fallback
  // -------------------------------------------------------------------------
  json(res, 404, { error: { code: 404, message: "Endpoint not defined" } });
});

server.listen(PORT, () => {
  console.log(`library-api-mock running on http://localhost:${PORT}`);
});
