const express = require('express');
require('newrelic')
const winston = require('winston')
const newrelicFormatter = require('@newrelic/winston-enricher')(winston)

const logger = winston.createLogger({
  level: "info",
  maxsize: 5242880, // 5MB
  format: winston.format.combine(winston.format.label({ label: "test" }), newrelicFormatter()),
  transports: [new winston.transports.Console()]
})


const app = express();

const host = process.env.NODE_SERVICE_HOST || '0.0.0.0';
const port = process.env.NODE_SERVICE_PORT || 30003;

app.get('/ping', (req, res) => {
  logger.info(req.rawHeaders);
  res.send('pong!')
});

app.listen(port, host, () => {
  logger.info(`Example app listening on port ${port}`)
});