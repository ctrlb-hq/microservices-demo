const express = require('express');
const app = express();

const host = process.env.NODE_SERVICE_HOST || '0.0.0.0';
const port = process.env.NODE_SERVICE_PORT || 30003;

app.get('/ping', (req, res) => {
  res.send('pong!')
});

app.listen(port, host, () => {
  console.log(`Example app listening on port ${port}`)
});