const express = require('express');
const app = express();
const config = require('./config/config.js');

app.all('*', (req, res) => {
  res.send('Alive!');
});

app.listen(config.port, () => {
  console.log('Server is listening on port', config.port);
});
