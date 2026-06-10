const express = require('express');
const mysql = require('mysql2/promise');
const fs = require('fs');
const path = require('path');

const app = express();
app.use(express.json());

const db = mysql.createPool({
  host: 'localhost',
  user: 'root',
  database: 'pokedex',
});

async function seedDatabase() {
  const csvPath = path.join(__dirname, '../../../pokedex.csv');
  const lines = fs.readFileSync(csvPath, 'utf8').trim().split('\n').slice(1); // skip header

  const rows = lines.map(line => {
    const [number, name, type, total, hp, attack, defense, sp_atk, sp_def, speed] = line.split(',');
    return [number, name, type, parseInt(total), parseInt(hp), parseInt(attack), parseInt(defense), parseInt(sp_atk), parseInt(sp_def), parseInt(speed)];
  });

  await db.query('TRUNCATE TABLE pokemon');
  await db.query(
    'INSERT INTO pokemon (number, name, type, total, hp, attack, defense, sp_atk, sp_def, speed) VALUES ?',
    [rows]
  );
  console.log(`Seeded ${rows.length} Pokemon from CSV.`);
}

app.get('/pokemon', async (req, res) => {
  try {
    const { limit = 20, offset = 0 } = req.query;
    const [rows] = await db.query(
      'SELECT * FROM pokemon LIMIT ? OFFSET ?',
      [parseInt(limit), parseInt(offset)]
    );
    const [[{ total }]] = await db.query('SELECT COUNT(*) as total FROM pokemon');
    res.json({ total, limit: parseInt(limit), offset: parseInt(offset), data: rows });
  } catch (err) {
    console.error('GET /pokemon error:', err.message);
    res.status(500).json({ error: 'Database error', details: err.message });
  }
});

app.get('/pokemon/:id', async (req, res) => {
  try {
    const [rows] = await db.query('SELECT * FROM pokemon WHERE number = ?', [req.params.id]);
    if (!rows.length) return res.status(404).json({ error: 'Pokemon not found' });
    res.json(rows.length === 1 ? rows[0] : rows);
  } catch (err) {
    console.error('GET /pokemon/:id error:', err.message);
    res.status(500).json({ error: 'Database error', details: err.message });
  }
});

app.post('/pokemon', async (req, res) => {
  try {
    const { number, name, type, total, hp, attack, defense, sp_atk, sp_def, speed } = req.body;
    if (!number || !name || !type) return res.status(400).json({ error: 'number, name, and type are required' });
    await db.query(
      'INSERT INTO pokemon (number, name, type, total, hp, attack, defense, sp_atk, sp_def, speed) VALUES (?,?,?,?,?,?,?,?,?,?)',
      [number, name, type, total, hp, attack, defense, sp_atk, sp_def, speed]
    );
    const [[created]] = await db.query('SELECT * FROM pokemon WHERE number = ? AND name = ?', [number, name]);
    res.status(201).json(created);
  } catch (err) {
    console.error('POST /pokemon error:', err.message);
    res.status(500).json({ error: 'Database error', details: err.message });
  }
});

app.patch('/pokemon/:id', async (req, res) => {
  try {
    const fields = ['name', 'type', 'total', 'hp', 'attack', 'defense', 'sp_atk', 'sp_def', 'speed'];
    const updates = fields.filter(f => req.body[f] !== undefined);
    if (!updates.length) return res.status(400).json({ error: 'No valid fields to update' });
    const sql = `UPDATE pokemon SET ${updates.map(f => `${f} = ?`).join(', ')} WHERE number = ?`;
    const values = [...updates.map(f => req.body[f]), req.params.id];
    const [result] = await db.query(sql, values);
    if (!result.affectedRows) return res.status(404).json({ error: 'Pokemon not found' });
    const [rows] = await db.query('SELECT * FROM pokemon WHERE number = ?', [req.params.id]);
    res.json(rows.length === 1 ? rows[0] : rows);
  } catch (err) {
    console.error('PATCH /pokemon/:id error:', err.message);
    res.status(500).json({ error: 'Database error', details: err.message });
  }
});

app.delete('/pokemon/:id', async (req, res) => {
  try {
    const [result] = await db.query('DELETE FROM pokemon WHERE number = ?', [req.params.id]);
    if (!result.affectedRows) return res.status(404).json({ error: 'Pokemon not found' });
    res.status(204).send();
  } catch (err) {
    console.error('DELETE /pokemon/:id error:', err.message);
    res.status(500).json({ error: 'Database error', details: err.message });
  }
});

seedDatabase()
  .then(() => {
    app.listen(3001, () => console.log('Pokemon service running on port 3001'));
  })
  .catch(err => {
    console.error('Failed to seed database:', err.message);
    process.exit(1);
  });
